package userRoleL

import (
	"context"
	v1 "demo/api/userRole/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/do"
	"demo/utility/token"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

type UserRole struct {
}

func (s *UserRole) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	err = dao.UserRole.Ctx(ctx).Where("id", req.Id).Scan(&res)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *UserRole) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)
	createdAt := gtime.NewFromTime(time.Now())

	data, err := dao.UserRole.Ctx(ctx).Data(do.UserRole{
		UserId:    req.UserId,
		RoleId:    req.RoleId,
		CreatedBy: payload.Id,
		CreatedAt: createdAt,
	}).Insert()

	if err != nil {
		return nil, err
	}

	id, _ := data.LastInsertId()

	res = &v1.CreateRes{
		Id:       id,
		UserId:   req.UserId,
		RoleId:   req.RoleId,
		CreateBy: int64(payload.Id),
		CreateAt: createdAt,
	}

	return res, nil
}

func (s *UserRole) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)

	updateAt := gtime.NewFromTime(time.Now())

	_, err = dao.UserRole.Ctx(ctx).Where("id", req.Id).Data(do.UserRole{
		UserId:    req.UserId,
		RoleId:    req.RoleId,
		UpdatedBy: payload.Id,
		UpdatedAt: updateAt,
	}).Update()

	if err != nil {
		return nil, err
	}

	res = &v1.UpdateRes{
		UserId:   req.UserId,
		RoleId:   req.RoleId,
		UpdateBy: int64(payload.Id),
		UpdateAt: updateAt,
	}

	return res, nil
}

func (s *UserRole) Delete(ctx context.Context, id int64) error {
	db := dao.UserRole.Ctx(ctx).Where("id", id)

	_, err := db.Delete()

	if err != nil {
		return err
	}

	return nil
}
