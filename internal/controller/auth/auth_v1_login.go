package auth

import (
	"context"

	v1 "strongbody-api/api/auth/v1"
	"strongbody-api/global"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	user, err := c.authService.Login(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)
		return nil, err
	}

	global.Logger.Info(ctx, "Login success")

	return user, err
}
