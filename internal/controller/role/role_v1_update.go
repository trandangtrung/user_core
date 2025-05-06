package role

import (
	"context"

	v1 "github.com/quannv/strongbody-api/api/role/v1"
	"github.com/quannv/strongbody-api/global"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	res, err = c.roleService.Update(ctx, req)

	if err != nil {
		global.Logger.Error(ctx, err)

		return res, err
	}
	global.Logger.Info(ctx, "Update role success")

	return res, nil
}
