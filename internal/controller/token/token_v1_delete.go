package token

import (
	"context"

	v1 "demo/api/token/v1"
	"demo/global"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = c.tokenService.Delete(ctx, req.Id)

	if err != nil {
		global.Logger.Error(ctx, err)

		return &v1.DeleteRes{
			Status: "error",
		}, err
	}

	return &v1.DeleteRes{
		Status: "success",
	}, nil
}
