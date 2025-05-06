package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/quannv/strongbody-api/api/user/v1"
	"github.com/quannv/strongbody-api/global"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	global.Logger.Debug(ctx, req)

	res, err = c.userService.Update(ctx, req)
	if err != nil {
		global.Logger.Error(ctx, err)
		return nil, gerror.NewCode(gcode.CodeNotFound, "user not found")
	}

	global.Logger.Info(ctx, "Update user success")
	return res, nil
}
