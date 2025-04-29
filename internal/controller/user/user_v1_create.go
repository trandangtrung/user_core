package user

import (
	"context"

	v1 "demo/api/user/v1"
	"demo/global"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	user, err := c.userService.CreateByAdmin(ctx, req)
	if err != nil {
		global.Logger.Error(ctx, err)
		return nil, err
	}

	global.Logger.Info(ctx, "Create user success")
	return user, err
}
