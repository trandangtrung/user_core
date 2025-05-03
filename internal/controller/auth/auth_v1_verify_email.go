package auth

import (
	"context"

	v1 "github.com/quannv/strongbody-api/api/auth/v1"
	"github.com/quannv/strongbody-api/global"
)

func (c *ControllerV1) VerifyEmail(ctx context.Context, req *v1.VerifyEmailReq) (res *v1.VerifyEmailRes, err error) {
	res, err = c.authService.VerifyEmail(ctx, req)
	if err != nil {
		global.Logger.Error(ctx, err)
		return nil, err
	}
	global.Logger.Info(ctx, "Verify email success")
	return res, nil
}
