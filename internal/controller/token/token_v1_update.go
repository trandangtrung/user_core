package token

import (
	"context"

	v1 "demo/api/token/v1"
	"demo/global"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	res, err = c.token.Update(ctx, req.Id, req)

	if err != nil {
		global.Logger.Error(ctx, err)

		return res, err
	}

	return res, nil
}
