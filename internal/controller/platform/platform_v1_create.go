package platform

import (
	"context"

	v1 "demo/api/platform/v1"
	"demo/global"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	global.Logger.Debug(ctx, req)

	res, err = c.platformService.Create(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)

		return nil, err
	}
	global.Logger.Info(ctx, "Create platform success")
	return res, nil

}
