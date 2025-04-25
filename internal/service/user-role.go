package service

import (
	"context"
	v1 "demo/api/userRole/v1"
	"demo/internal/repository"
)

type (
	UserRoleService interface {
		Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
		Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
		Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
		Delete(ctx context.Context, id int64) error
	}
	userRoleService struct {
		userRoleRepo repository.UserRoleRepository
	}
)

func NewUserRoleService(userRoleRepo repository.UserRoleRepository) UserRoleService {
	return &userRoleService{
		userRoleRepo: userRoleRepo,
	}
}

func (u *userRoleService) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	return nil, nil
}

func (u *userRoleService) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return nil, nil
}

func (u *userRoleService) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, nil
}

func (u *userRoleService) Delete(ctx context.Context, id int64) error {
	return nil
}
