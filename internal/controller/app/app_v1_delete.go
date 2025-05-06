package app

import (
	"context"

	v1 "github.com/quannv/strongbody-api/api/app/v1"
	"github.com/quannv/strongbody-api/global"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = c.appService.Delete(ctx, req.Id)

	if err != nil {
		global.Logger.Error(ctx, err)

		return &v1.DeleteRes{
			Status: "error",
		}, err
	}
	global.Logger.Info(ctx, "Delete platform success")

	return &v1.DeleteRes{
		Status: "Delete success",
	}, nil
}
