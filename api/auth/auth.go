// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"github.com/quannv/strongbody-api/api/auth/v1"
)

type IAuthV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error)
	RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error)
	LoginByToken(ctx context.Context, req *v1.LoginByTokenReq) (res *v1.LoginByTokenRes, err error)
}
