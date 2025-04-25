package service

import (
	"context"
	v1 "demo/api/userPlatform/v1"
	"demo/internal/entity"
	"demo/internal/repository"
	"errors"

	"github.com/gogf/gf/v2/os/gtime"
)

type (
	UserPlatformService interface {
		Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
		Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
		Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
		Delete(ctx context.Context, id int64) error
	}
	userPlatformService struct {
		userPlatformRepo repository.UserPlatformRepository
	}
)

func NewUserPlatformService(userPlatformRepo repository.UserPlatformRepository) UserPlatformService {
	return &userPlatformService{
		userPlatformRepo: userPlatformRepo,
	}
}

func (l *userPlatformService) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	record, err := l.userPlatformRepo.GetById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, errors.New("user-platform get not found")
	}

	return &v1.GetRes{
		Id:         int64(record.ID),
		UserId:     int64(record.UserID),
		PlatformId: int64(record.PlatformID),
		CreateAt:   gtime.New(record.CreatedAt),
		CreateBy:   int64(record.CreatedBy),
	}, nil
}

func (l *userPlatformService) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	userPlatform := &entity.UserPlatform{
		UserID:     uint(req.UserId),
		PlatformID: uint(req.PlatformId),
	}
	created, err := l.userPlatformRepo.Create(ctx, userPlatform)
	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{
		Id:         int64(created.ID),
		UserId:     int64(created.UserID),
		PlatformId: int64(created.PlatformID),
		CreateAt:   gtime.New(created.CreatedAt),
		CreateBy:   int64(created.CreatedBy),
	}, nil
}

func (l *userPlatformService) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {

	existing, err := l.userPlatformRepo.GetById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("user-platform record not found")
	}

	existing.UserID = uint(req.UserId)
	existing.PlatformID = uint(req.PlatformId)

	updated, err := l.userPlatformRepo.Update(ctx, existing)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateRes{
		Id:         int64(updated.ID),
		UserId:     int64(updated.UserID),
		PlatformId: int64(updated.PlatformID),
		UpdateAt:   gtime.New(updated.UpdatedAt),
		UpdateBy:   int64(updated.UpdatedBy),
	}, nil
}

func (l *userPlatformService) Delete(ctx context.Context, id int64) error {
	userPlatform := &entity.UserPlatform{UserID: uint(id)}
	return l.userPlatformRepo.Delete(ctx, userPlatform)
}
