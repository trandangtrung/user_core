package user

import (
	"context"

	v1 "demo/api/user/v1"
	"demo/global"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	global.Logger.Debug(ctx, req)

	res, err = c.userService.Delete(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)

		return nil, err
	}
	global.Logger.Info(ctx, "Create user role success")
	return res, nil
}
