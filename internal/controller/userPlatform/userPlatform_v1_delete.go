package userPlatform

import (
	"context"

	v1 "demo/api/userPlatform/v1"
	"demo/global"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = c.userPlatform.Delete(ctx, req.Id)

	if err != nil {
		global.Logger.Error(ctx, err)

		return &v1.DeleteRes{
			Status: "error",
		}, err
	}
	global.Logger.Info(ctx, "Delete user role  success")

	return &v1.DeleteRes{
		Status: "Delete success",
	}, nil
}
