package service

import (
	"context"
	v1 "demo/api/role/v1"
	"demo/internal/repository"
)

type (
	RoleService interface {
		Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
		Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
		Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
		Delete(ctx context.Context, id int64) error
	}
	roleService struct {
		roleRepo repository.RoleRepository
	}
)

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}

func (l *roleService) Get(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	return nil, nil
}

func (l *roleService) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return nil, nil
}

func (l *roleService) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return nil, nil
}

func (l *roleService) Delete(ctx context.Context, id int64) error {
	return nil
}
