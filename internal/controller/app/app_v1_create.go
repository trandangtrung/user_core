package app

import (
	"context"

	v1 "strongbody-api/api/app/v1"
	"strongbody-api/global"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	global.Logger.Debug(ctx, req)

	res, err = c.appService.Create(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)

		return nil, err
	}
	global.Logger.Info(ctx, "Create app success")
	return res, nil

}
