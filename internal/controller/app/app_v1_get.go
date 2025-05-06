package app

import (
	"context"

	v1 "github.com/quannv/strongbody-api/api/app/v1"
	"github.com/quannv/strongbody-api/global"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	res, err = c.appService.Get(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)

		return res, err
	}
	global.Logger.Info(ctx, "Get platform success")

	return res, nil
}
