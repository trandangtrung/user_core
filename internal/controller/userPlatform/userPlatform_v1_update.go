package userPlatform

import (
	"context"

	v1 "demo/api/userPlatform/v1"
	"demo/global"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	res, err = c.userPlatformService.Update(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)

		return res, err
	}
	global.Logger.Info(ctx, "Update user role success")

	return res, nil
}
