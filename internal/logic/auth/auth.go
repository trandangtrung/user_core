package authL

import (
	"context"
	v1 "demo/api/auth/v1"
	v1Token "demo/api/token/v1"
	"demo/global"
	"demo/internal/consts"
	"demo/internal/dao"
	tokenL "demo/internal/logic/token"
	"demo/internal/model/do"
	utils "demo/utility"
	"demo/utility/token"
	"fmt"

	"github.com/gogf/gf/frame/g"
)

type user struct {
	Id           int64
	Email        string
	PasswordHash string
	CreateAt     string
}

type AuthLogic struct {
	tokenL tokenL.Token
}

func (a *AuthLogic) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {

	user, err := a.CheckEmail(ctx, req.Email)

	if err != nil {
		return nil, fmt.Errorf("email is not exist")
	}

	err = utils.CheckPassword(user.PasswordHash, req.Password)

	if err != nil {
		return nil, fmt.Errorf("password is wrong")
	}

	// sau get role để đổi vô
	accessToken, _, err := global.Token.CreateToken(int(user.Id), "", global.VariableEnv.TimeAccess)

	if err != nil {
		return nil, err
	}

	refreshToken, _, err := global.Token.CreateToken(int(user.Id), "", global.VariableEnv.TimeRefresh)

	if err != nil {
		return nil, err
	}

	scope := ctx.Value(consts.AuthorizationScope).(string)

	role, err := a.GetRoleByUserIDAndPlatform(ctx, scope, user.Id)

	if err != nil {
		return nil, err
	}

	reqToken := &v1Token.CreateReq{
		User_id:      user.Id,
		RefreshToken: refreshToken,
		Scope:        scope,
		
	}

	err = a.tokenL.Create(ctx, reqToken)

	if err != nil {
		return nil, err
	}

	return &v1.LoginRes{
		User: v1.User{
			Email: user.Email,
			Role:  role.Name,
		},
		Token: v1.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (a *AuthLogic) Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error) {

	_, err = a.CheckEmail(ctx, req.Email)

	if err == nil {
		return nil, err
	}

	hashPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	_, err = dao.Users.Ctx(ctx).Data(do.Users{
		Email:        req.Email,
		PasswordHash: hashPassword,
	}).Insert()

	if err != nil {
		return nil, err
	}

	return &v1.SignupRes{
		Status: "Sign up success",
	}, nil
}

func (a *AuthLogic) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)

	accessToken, _, err := global.Token.CreateToken(int(payload.Id), payload.Permissions, global.VariableEnv.TimeAccess)

	if err != nil {
		return nil, err
	}

	return &v1.RefreshTokenRes{
		AccessToken: accessToken,
	}, nil
}

func (a *AuthLogic) CheckEmail(ctx context.Context, email string) (user, error) {
	var user user
	err := dao.Users.Ctx(ctx).Where("email", email).Scan(&user)

	if err != nil {
		return user, err

	}

	global.Logger.Debug(ctx, user)

	return user, nil
}

type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a *AuthLogic) GetRoleByUserIDAndPlatform(ctx context.Context, platformName string, userId int64) (*Role, error) {

	var role *Role
	err := g.DB().Model("Role r").
		LeftJoin("UserRole ur", "ur.role_id = r.id").
		LeftJoin("Users u", "u.id = ur.user_id").
		LeftJoin("UserPlatform sup", "sup.user_id = u.id").
		LeftJoin("Platform sp", "sp.id = sup.platform_id").
		Where(g.Map{
			"sp.name": platformName,
			"u.id":    userId,
		}).
		Fields("r.*").
		Struct(&role)

	if err != nil {
		return nil, err
	}

	return role, nil
}
