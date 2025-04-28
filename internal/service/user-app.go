package service

import (
	"context"
	v1 "demo/api/userApp/v1"
	"demo/internal/entity"
	"demo/internal/repository"
	"errors"

	"github.com/gogf/gf/v2/os/gtime"
)

type (
	UserAppService interface {
		Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
		Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
		Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
		Delete(ctx context.Context, id int64) error
	}
	userAppService struct {
		userAppRepo repository.UserAppRepository
	}
)

func NewUserAppService(userAppRepo repository.UserAppRepository) UserAppService {
	return &userAppService{
		userAppRepo: userAppRepo,
	}
}

func (l *userAppService) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	record, err := l.userAppRepo.GetById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, errors.New("user-platform get not found")
	}

	return &v1.GetRes{
		Id:       int64(record.ID),
		UserId:   int64(record.UserID),
		AppId:    int64(record.AppID),
		CreateAt: gtime.New(record.CreatedAt),
		CreateBy: int64(record.CreatedBy),
	}, nil
}

func (l *userAppService) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	userPlatform := &entity.UserApp{
		UserID: uint(req.UserId),
		AppID:  uint(req.AppId),
	}
	created, err := l.userAppRepo.Create(ctx, userPlatform)
	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{
		Id:       int64(created.ID),
		UserId:   int64(created.UserID),
		AppId:    int64(created.AppID),
		CreateAt: gtime.New(created.CreatedAt),
		CreateBy: int64(created.CreatedBy),
	}, nil
}

func (l *userAppService) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {

	existing, err := l.userAppRepo.GetById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("user-platform record not found")
	}

	existing.UserID = uint(req.UserId)
	existing.AppID = uint(req.AppId)

	updated, err := l.userAppRepo.Update(ctx, existing)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateRes{
		Id:       int64(updated.ID),
		UserId:   int64(updated.UserID),
		AppId:    int64(updated.AppID),
		UpdateAt: gtime.New(updated.UpdatedAt),
		UpdateBy: int64(updated.UpdatedBy),
	}, nil
}

func (l *userAppService) Delete(ctx context.Context, id int64) error {
	userPlatform := &entity.UserApp{UserID: uint(id)}
	return l.userAppRepo.Delete(ctx, userPlatform)
}
