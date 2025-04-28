package app

import (
	"context"

	v1 "demo/api/app/v1"
	"demo/global"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	res, err = c.appService.Update(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)

		return res, err
	}
	global.Logger.Info(ctx, "Update platform success")

	return res, nil
}
