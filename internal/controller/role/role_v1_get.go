package role

import (
	"context"

	v1 "github.com/quannv/strongbody-api/api/role/v1"
	"github.com/quannv/strongbody-api/global"
)

func (c *ControllerV1) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	res, err = c.roleService.GetByID(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)

		return res, err
	}
	global.Logger.Info(ctx, "Get role success")

	return res, nil
}
