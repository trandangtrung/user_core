package auth

import (
	"context"

	v1 "strongbody-api/api/auth/v1"
	"strongbody-api/global"
)

func (c *ControllerV1) Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error) {
	res, err = c.authService.Signup(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)
		return nil, err
	}

	global.Logger.Info(ctx, "Sign up success")

	return res, err

}
