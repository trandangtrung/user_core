package service

import (
	"context"

	v1 "github.com/quannv/strongbody-api/api/auth/v1"
	"github.com/quannv/strongbody-api/global"
	"github.com/quannv/strongbody-api/internal/config"
	"github.com/quannv/strongbody-api/internal/consts"
	"github.com/quannv/strongbody-api/internal/entity"
	"github.com/quannv/strongbody-api/internal/repository"
	utils "github.com/quannv/strongbody-api/utility"
	rescode "github.com/quannv/strongbody-api/utility/resCode"
	"github.com/quannv/strongbody-api/utility/token"

	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	AuthService interface {
		Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
		Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error)
		RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error)
		LoginByToken(ctx context.Context, req *v1.LoginByTokenReq) (res *v1.LoginByTokenRes, err error)
	}
	authService struct {
		userRepo repository.UserRepository
		roleRepo repository.RoleRepository
	}
)

func NewAuthService(userRepo repository.UserRepository, roleRepo repository.RoleRepository) AuthService {
	return &authService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (a *authService) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	// check if the email is already exist
	user, err := a.checkEmail(ctx, req.Email)
	if err != nil {
		return nil, gerror.NewCode(rescode.EmailNotFound, "email not found")
	}

	err = utils.CheckPassword(user.PasswordHashed, req.Password)
	if err != nil {
		return nil, gerror.NewCode(rescode.EmailNotFound, "email not found")
	}

	accessToken, _, err := global.Token.CreateToken(int(user.ID), "", config.GetConfig().JwtCfg.TimeAccess)
	if err != nil {
		return nil, gerror.NewCode(rescode.AccessTokenCreationFailed, err.Error())
	}

	refreshToken, _, err := global.Token.CreateToken(int(user.ID), "", config.GetConfig().JwtCfg.TimeRefresh)
	if err != nil {
		return nil, gerror.NewCode(rescode.RefreshTokenCreationFailed, err.Error())
	}

	scope := ctx.Value(consts.AuthorizationScope).(string)

	roles, err := a.roleRepo.GetRolesByUserIDAndAppName(ctx, user.ID, scope)
	if err != nil {
		return nil, err
	}
	if roles == nil {
		return nil, gerror.NewCode(rescode.RoleNotFound, "role is not exist")
	}
	if len(roles) == 0 {
		return nil, gerror.NewCode(rescode.RoleNotFound, "role is not exist")
	}

	token, err := a.userRepo.CreateToken(ctx, &entity.Token{
		UserID:       user.ID,
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

func (a *authService) LoginByToken(ctx context.Context, req *v1.LoginByTokenReq) (res *v1.LoginByTokenRes, err error) {

	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)

	user, err := a.userRepo.GetUserByID(ctx, uint(payload.Id))

	if err != nil {
		return nil, gerror.NewCode(rescode.UserNotFound, "user is not exist")
	}

	accessToken, _, err := global.Token.CreateToken(payload.Id, "", config.GetConfig().JwtCfg.TimeAccess)
	if err != nil {
		return nil, gerror.NewCode(rescode.AccessTokenCreationFailed, err.Error())
	}

	refreshToken, _, err := global.Token.CreateToken(payload.Id, "", config.GetConfig().JwtCfg.TimeRefresh)
	if err != nil {
		return nil, gerror.NewCode(rescode.AccessTokenCreationFailed, err.Error())
	}

	scope := ctx.Value(consts.AuthorizationScope).(string)

	roles, err := a.roleRepo.GetRolesByUserIDAndAppName(ctx, user.ID, scope)
	if err != nil {
		return nil, err
	}
	if roles == nil {
		return nil, gerror.NewCode(rescode.UserNotFound, "user is not exist")
	}
	if len(roles) == 0 {
		return nil, gerror.NewCode(rescode.UserNotFound, "user is not exist")
	}

	token, err := a.userRepo.CreateToken(ctx, &entity.Token{
		UserID:       uint(payload.Id),
		RefreshToken: refreshToken,
		Scope:        scope,
	})
	if err != nil {
		return nil, err
	}

	return &v1.LoginByTokenRes{
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
	//	// check if the email is already exist
	_, err = a.checkEmail(ctx, req.Email)
	if err == nil {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, gerror.NewCode(rescode.HashPasswordFailed, "hash password failed")
	}

	_, err = a.userRepo.CreateUser(ctx, &entity.User{
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
	// check if the refresh token is valid
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)

	accessToken, _, err := global.Token.CreateToken(int(payload.Id), payload.Permissions, config.GetConfig().JwtCfg.TimeAccess)
	if err != nil {
		return nil, gerror.NewCode(rescode.AccessTokenCreationFailed, err.Error())
	}

	return &v1.RefreshTokenRes{
		AccessToken: accessToken,
	}, nil
}

func (a *authService) checkEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := a.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, gerror.NewCode(rescode.EmailNotFound, "email is not exist")
	}
	if user == nil {
		return nil, gerror.NewCode(rescode.EmailNotFound, "email is not exist")
	}

	global.Logger.Debug(ctx, user)
	return user, nil
}
