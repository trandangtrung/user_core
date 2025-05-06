package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/quannv/strongbody-api/api/auth/v1"
)

func (c *ControllerV1) ResendVerifyEmail(ctx context.Context, req *v1.ResendVerifyEmailReq) (res *v1.ResendVerifyEmailRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
