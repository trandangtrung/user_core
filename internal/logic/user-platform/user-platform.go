package userPlatformL

import (
	"context"
	v1 "demo/api/userPlatform/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/do"
	"demo/utility/token"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

type UserPlatform struct {
}

func (s *UserPlatform) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	err = dao.UserPlatform.Ctx(ctx).Where("id", req.Id).Scan(&res)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *UserPlatform) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)
	createdAt := gtime.NewFromTime(time.Now())

	data, err := dao.UserPlatform.Ctx(ctx).Data(do.UserPlatform{
		UserId:     req.UserId,
		PlatformId: req.PlatformId,
		CreatedBy:  payload.Id,
		CreatedAt:  createdAt,
	}).Insert()

	if err != nil {
		return nil, err
	}

	id, _ := data.LastInsertId()

	res = &v1.CreateRes{
		Id:         id,
		UserId:     req.UserId,
		PlatformId: req.PlatformId,
		CreateBy:   int64(payload.Id),
		CreateAt:   createdAt,
	}

	return res, nil
}

func (s *UserPlatform) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)

	updateAt := gtime.NewFromTime(time.Now())

	_, err = dao.UserPlatform.Ctx(ctx).Where("id", req.Id).Data(do.UserPlatform{
		UserId:     req.UserId,
		PlatformId: req.PlatformId,
		UpdatedBy:  payload.Id,
		UpdatedAt:  updateAt,
	}).Update()

	if err != nil {
		return nil, err
	}

	res = &v1.UpdateRes{
		UserId:     req.UserId,
		PlatformId: req.PlatformId,
		UpdateBy:   int64(payload.Id),
		UpdateAt:   updateAt,
	}

	return res, nil
}

func (s *UserPlatform) Delete(ctx context.Context, id int64) error {
	db := dao.UserPlatform.Ctx(ctx).Where("id", id)

	_, err := db.Delete()

	if err != nil {
		return err
	}

	return nil
}
