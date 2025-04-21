package token

import (
	"context"

	v1 "demo/api/token/v1"
	"demo/global"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	global.Logger.Debug(ctx, req)

	err = c.token.Create(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)

		return &v1.CreateRes{
			Status: "error",
		}, err
	}

	return &v1.CreateRes{
		Status: "success",
	}, nil
}
