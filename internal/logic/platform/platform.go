package platform

import (
	"context"
	v1 "demo/api/platform/v1"
	"demo/internal/consts"
	"demo/internal/dao"
	"demo/internal/model/do"
	"demo/utility/token"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

type Platform struct {
}

func (s *Platform) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	err = dao.Platform.Ctx(ctx).Where("id", req.Id).Scan(&res)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *Platform) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)
	createdAt := gtime.NewFromTime(time.Now())

	data, err := dao.Platform.Ctx(ctx).Data(do.Platform{
		Name:     req.Name,
		Config:   req.Config,
		CreatedBy: payload.Id,
		CreatedAt: createdAt,
	}).Insert()

	if err != nil {
		return nil, err
	}

	id, _ := data.LastInsertId()

	res = &v1.CreateRes{
		Id:       id,
		Name:     req.Name,
		Config:   req.Config,
		CreateBy: int64(payload.Id),
		CreateAt: createdAt,
	}

	return res, nil
}

func (s *Platform) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)

	updateAt := gtime.NewFromTime(time.Now())

	_, err = dao.Platform.Ctx(ctx).Where("id", req.Id).Data(do.Platform{
		Name:     req.Name,
		Config:   req.Config,
		UpdatedBy: payload.Id,
		UpdatedAt: updateAt,
	}).Update()

	if err != nil {
		return nil, err
	}

	res = &v1.UpdateRes{
		Name:     req.Name,
		Config:   req.Config,
		UpdateBy: int64(payload.Id),
		UpdateAt: updateAt,
	}

	return res, nil
}

func (s *Platform) Delete(ctx context.Context, id int64) error {
	db := dao.Platform.Ctx(ctx).Where("id", id)

	_, err := db.Delete()

	if err != nil {
		return err
	}

	return nil
}
