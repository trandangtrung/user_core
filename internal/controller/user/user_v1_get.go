package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "strongbody-api/api/user/v1"
	"strongbody-api/global"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	global.Logger.Debug(ctx, req)

	res, err = c.userService.GetByID(ctx, req)
	if err != nil {
		global.Logger.Error(ctx, err)
		return nil, gerror.NewCode(gcode.CodeNotFound, "user not found")
	}

	global.Logger.Info(ctx, "Get user success")
	return res, nil
}
