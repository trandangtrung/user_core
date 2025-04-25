package service

import (
	"context"
	v1 "demo/api/platform/v1"
	"demo/internal/repository"
)

type (
	PlatformService interface {
		Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
		Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
		Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
		Delete(ctx context.Context, id int64) error
	}
	platformService struct {
		platformRepo repository.PlatformRepository
	}
)

func NewPlatformService(platformRepo repository.PlatformRepository) PlatformService {
	return &platformService{
		platformRepo: platformRepo,
	}
}

func (l *platformService) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	return nil, nil
}

func (l *platformService) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return nil, nil
}

func (l *platformService) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, nil
}

func (l *platformService) Delete(ctx context.Context, id int64) error {
	return nil
}
