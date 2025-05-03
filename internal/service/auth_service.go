package service

import (
	"context"
	"time"

	v1 "github.com/quannv/strongbody-api/api/auth/v1"
	"github.com/quannv/strongbody-api/global"
	"github.com/quannv/strongbody-api/internal/config"
	"github.com/quannv/strongbody-api/internal/consts"
	"github.com/quannv/strongbody-api/internal/entity"
	"github.com/quannv/strongbody-api/internal/repository"
	"github.com/quannv/strongbody-api/internal/storage/postgres"
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
		VerifyEmail(ctx context.Context, req *v1.VerifyEmailReq) (res *v1.VerifyEmailRes, err error)
		ResendVerifyEmail(ctx context.Context, req *v1.ResendVerifyEmailReq) (res *v1.ResendVerifyEmailRes, err error)
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
	user, err := a.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, gerror.NewCode(rescode.EmailNotFound, "email is not exist")
	}
	if user == nil {
		return nil, gerror.NewCode(rescode.EmailNotFound, "email is not exist")
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
	// get database connection
	db := postgres.GetDatabaseConnection().Connection

	// check if the email is already exist
	user, err := a.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, gerror.NewCode(rescode.EmailNotFound, "email is not exist")
	}
	if user != nil {
		return nil, gerror.NewCode(rescode.EmailAlreadyExists, "email already exist")
	}

	// generate otp
	otp := utils.GenerateOTP()

	// hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, gerror.NewCode(rescode.InternalError, err.Error())
	}

	// start transaction
	tx := db.Begin()
	if tx.Error != nil {
		return nil, gerror.NewCode(rescode.InternalError, tx.Error.Error())
	}

	// check if pending user is exist
	pendingUser, err := a.userRepo.GetPendingUserByEmail(ctx, tx, req.Email)
	if err != nil {
		tx.Rollback()
		return nil, gerror.NewCode(rescode.InternalError, err.Error())
	}
	if pendingUser != nil {
		// update pending user
		pendingUser.PasswordHashed = hashedPassword
		_, err = a.userRepo.UpdatePendingUser(ctx, tx, pendingUser)
		if err != nil {
			tx.Rollback()
			return nil, gerror.NewCode(rescode.InternalError, err.Error())
		}
	} else {
		// create new pending user
		pendingUser = &entity.PendingUser{
			Email:          req.Email,
			PasswordHashed: hashedPassword,
		}
		_, err = a.userRepo.CreatePendingUser(ctx, tx, pendingUser)
		if err != nil {
			tx.Rollback()
			return nil, gerror.NewCode(rescode.InternalError, err.Error())
		}
	}

	// create email otp
	_, err = a.userRepo.CreateEmailOTP(ctx, tx, &entity.EmailOTP{
		Email:    req.Email,
		OTP:      otp,
		ExpireAt: time.Now().Add(10 * time.Minute),
	})
	if err != nil {
		tx.Rollback()
		return nil, gerror.NewCode(rescode.InternalError, err.Error())
	}

	// send email
	var toEmail []string
	toEmail = append(toEmail, req.Email)
	err = global.Gmail.SendEmail("Strongbody - Verify Email", "Your OTP is: "+otp, toEmail, nil, nil, nil)
	if err != nil {
		tx.Rollback()
		return nil, gerror.NewCode(rescode.EmailOTPSendFailed, err.Error())
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, gerror.NewCode(rescode.InternalError, "transaction failed")
	}

	return &v1.SignupRes{
		Status: "Send email successfully",
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

func (a *authService) VerifyEmail(ctx context.Context, req *v1.VerifyEmailReq) (res *v1.VerifyEmailRes, err error) {
	// get database connection
	db := postgres.GetDatabaseConnection().Connection

	// start transaction
	tx := db.Begin()
	if tx.Error != nil {
		return nil, gerror.NewCode(rescode.InternalError, tx.Error.Error())
	}

	// check if the otp is valid
	isValid, err := a.userRepo.IsOtpValid(ctx, tx, req.Email, req.Otp)
	if err != nil {
		return nil, gerror.NewCode(rescode.EmailOTPInvalid, "otp is invalid")
	}
	if !isValid {
		return nil, gerror.NewCode(rescode.EmailOTPInvalid, "otp is invalid")
	}

	// get pending user
	pendingUser, err := a.userRepo.GetPendingUserByEmail(ctx, tx, req.Email)
	if err != nil {
		tx.Rollback()
		return nil, gerror.NewCode(rescode.InternalError, err.Error())
	}
	if pendingUser == nil {
		tx.Rollback()
		return nil, gerror.NewCode(rescode.EmailNotFound, "email is not exist")
	}

	// create user
	_, err = a.userRepo.CreateUser(ctx, tx, &entity.User{
		Email:          pendingUser.Email,
		PasswordHashed: pendingUser.PasswordHashed,
	})
	if err != nil {
		tx.Rollback()
		return nil, gerror.NewCode(rescode.InternalError, err.Error())
	}

	// delete pending user
	err = a.userRepo.DeletePendingUser(ctx, tx, pendingUser.ID)
	if err != nil {
		tx.Rollback()
		return nil, gerror.NewCode(rescode.InternalError, err.Error())
	}

	// update email otp
	emailOtp, err := a.userRepo.GetEmailOTPByEmail(ctx, tx, req.Email)
	if err != nil {
		tx.Rollback()
		return nil, gerror.NewCode(rescode.InternalError, err.Error())
	}
	emailOtp.Used = true
	_, err = a.userRepo.UpdateEmailOTP(ctx, tx, emailOtp)
	if err != nil {
		tx.Rollback()
		return nil, gerror.NewCode(rescode.InternalError, err.Error())
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, gerror.NewCode(rescode.InternalError, "transaction failed")
	}

	return &v1.VerifyEmailRes{
		Status: "Verify email successfully",
	}, nil
}

func (a *authService) ResendVerifyEmail(ctx context.Context, req *v1.ResendVerifyEmailReq) (res *v1.ResendVerifyEmailRes, err error) {
	return nil, nil
}
