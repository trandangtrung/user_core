package service

import (
	"context"
	v1 "demo/api/auth/v1"
	"demo/global"
	"demo/internal/config"
	"demo/internal/consts"
	"demo/internal/entity"
	"demo/internal/repository"
	utils "demo/utility"
	"demo/utility/token"
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	AuthService interface {
		Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
		Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error)
		RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error)
	}
	authService struct {
		userRepo  repository.UserRepository
		roleRepo  repository.RoleRepository
		tokenRepo repository.TokenRepository
	}
)

func NewAuthService(userRepo repository.UserRepository, roleRepo repository.RoleRepository, tokenRepo repository.TokenRepository) AuthService {
	return &authService{
		userRepo:  userRepo,
		roleRepo:  roleRepo,
		tokenRepo: tokenRepo,
	}
}

func (a *authService) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {

	user, err := a.checkEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("email is not exist")
	}

	err = utils.CheckPassword(user.PasswordHashed, req.Password)
	if err != nil {
		return nil, fmt.Errorf("password is wrong")
	}

	// sau get role để đổi vô
	accessToken, _, err := global.Token.CreateToken(int(user.ID), "", config.GetConfig().JwtCfg.TimeAccess)
	if err != nil {
		return nil, err
	}

	refreshToken, _, err := global.Token.CreateToken(int(user.ID), "", config.GetConfig().JwtCfg.TimeRefresh)
	if err != nil {
		return nil, err
	}

	scope := ctx.Value(consts.AuthorizationScope).(string)

	roles, err := a.roleRepo.GetRolesByUserIDAndPlatformName(ctx, scope, uint(user.ID))
	if err != nil {
		return nil, err
	}
	if roles == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "role is not exist")
	}
	if len(roles) == 0 {
		return nil, gerror.NewCode(gcode.CodeNotFound, "role is not exist")
	}

	token, err := a.tokenRepo.Create(ctx, &entity.Token{
		UserID:       uint(user.ID),
		RefreshToken: refreshToken,
		Scope:        scope,
	})
	if err != nil {
		return nil, err
	}

	return &v1.LoginRes{
		User: v1.User{
			Email: user.Email,
			Role:  roles[0].Name,
		},
		Token: v1.Token{
			AccessToken:  accessToken,
			RefreshToken: token.RefreshToken,
		},
	}, nil
}

func (a *authService) Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error) {

	_, err = a.checkEmail(ctx, req.Email)
	if err == nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	_, err = a.userRepo.Create(ctx, &entity.User{
		Email:          req.Email,
		PasswordHashed: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return &v1.SignupRes{
		Status: "Sign up success",
	}, nil
}

func (a *authService) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)

	accessToken, _, err := global.Token.CreateToken(int(payload.Id), payload.Permissions, config.GetConfig().JwtCfg.TimeAccess)
	if err != nil {
		return nil, err
	}

	return &v1.RefreshTokenRes{
		AccessToken: accessToken,
	}, nil
}

func (a *authService) checkEmail(ctx context.Context, email string) (*entity.User, error) {

	user, err := a.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "email is not exist")
	}
	if user == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "email is not exist")
	}

	global.Logger.Debug(ctx, user)
	return user, nil
}
