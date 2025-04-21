package role

import (
	"context"
	v1 "demo/api/role/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/do"
	"demo/utility/token"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

type Role struct {
}

func (s *Role) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	err = dao.Role.Ctx(ctx).Where("id", req.Id).Scan(&res)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *Role) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)
	createdAt := gtime.NewFromTime(time.Now())

	data, err := dao.Role.Ctx(ctx).Data(do.Role{
		UserPlatformId: req.UserPlatformId,
		Name:           req.Name,
		Description:    req.Description,
		CreatedBy:      payload.Id,
		CreatedAt:      createdAt,
	}).Insert()

	if err != nil {
		return nil, err
	}

	id, _ := data.LastInsertId()

	res = &v1.CreateRes{
		Id:             id,
		UserPlatformId: req.UserPlatformId,
		Name:           req.Name,
		Description:    req.Description,
		CreateBy:       int64(payload.Id),
		CreateAt:       createdAt,
	}

	return res, nil
}

func (s *Role) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)

	updateAt := gtime.NewFromTime(time.Now())

	_, err = dao.Role.Ctx(ctx).Where("id", req.Id).Data(do.Role{
		UserPlatformId: req.UserPlatformId,
		Name:           req.Name,
		Description:    req.Description,
		UpdatedBy:      payload.Id,
		UpdatedAt:      updateAt,
	}).Update()

	if err != nil {
		return nil, err
	}

	res = &v1.UpdateRes{
		UserPlatformId: req.UserPlatformId,
		Name:           req.Name,
		Description:    req.Description,
		UpdateBy:       int64(payload.Id),
		UpdateAt:       updateAt,
	}

	return res, nil
}

func (s *Role) Delete(ctx context.Context, id int64) error {
	db := dao.Role.Ctx(ctx).Where("id", id)

	_, err := db.Delete()

	if err != nil {
		return err
	}

	return nil
}
