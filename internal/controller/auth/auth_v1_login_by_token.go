package auth

import (
	"context"

	v1 "github.com/quannv/strongbody-api/api/auth/v1"
	"github.com/quannv/strongbody-api/global"
)

func (c *ControllerV1) LoginByToken(ctx context.Context, req *v1.LoginByTokenReq) (res *v1.LoginByTokenRes, err error) {
	user, err := c.authService.LoginByToken(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)
		return nil, err
	}

	global.Logger.Info(ctx, "Login success")

	return user, err
}
