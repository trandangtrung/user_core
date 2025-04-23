package userPlatform

import (
	"context"

	v1 "demo/api/userPlatform/v1"
	"demo/global"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	res, err = c.userPlatformService.Get(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)

		return res, err
	}
	global.Logger.Info(ctx, "Get user role  success")

	return res, nil
}
