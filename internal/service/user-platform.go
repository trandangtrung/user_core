package service

import (
	"context"
	v1 "demo/api/userPlatform/v1"
	"demo/internal/repository"
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
	return nil, nil
}

func (l *userPlatformService) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return nil, nil
}

func (l *userPlatformService) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, nil
}

func (l *userPlatformService) Delete(ctx context.Context, id int64) error {
	return nil
}
