package auth

import (
	"context"

	v1 "github.com/quannv/strongbody-api/api/auth/v1"
	"github.com/quannv/strongbody-api/global"
)

func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	res, err = c.authService.RefreshToken(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)
		return nil, err
	}

	global.Logger.Info(ctx, "get access token success")

	return res, err
}
