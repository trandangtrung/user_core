package service

import (
	"context"
	v1 "demo/api/userRole/v1"
	"demo/internal/entity"
	"demo/internal/repository"

	"github.com/gogf/gf/v2/os/gtime"
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
	userRole, err := u.userRoleRepo.GetById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	// Trả về kết quả
	return &v1.GetRes{
		Id:       int64(userRole.ID),
		UserId:   int64(userRole.UserID),
		RoleId:   int64(userRole.RoleID),
		CreateAt: gtime.New(userRole.CreatedAt),
		CreateBy: int64(userRole.CreatedBy),
		UpdateAt: gtime.New(userRole.UpdatedAt),
	}, nil
}

func (u *userRoleService) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	userRole := &entity.UserRole{
		UserID: uint(req.UserId),
		RoleID: uint(req.RoleId),
	}

	// Gọi repository để tạo mới UserRole
	created, err := u.userRoleRepo.Create(ctx, userRole)
	if err != nil {
		return nil, err
	}

	// Trả về response
	return &v1.CreateRes{
		Id:       int64(created.ID),
		UserId:   int64(created.UserID),
		RoleId:   int64(created.RoleID),
		CreateAt: gtime.New(created.CreatedAt),
		CreateBy: int64(created.CreatedBy),
	}, nil
}

func (u *userRoleService) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	userRole, err := u.userRoleRepo.GetById(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	userRole.UserID = uint(req.UserId)
	userRole.RoleID = uint(req.RoleId)

	// Lưu lại các thay đổi vào DB
	updatedUserRole, err := u.userRoleRepo.Update(ctx, userRole)
	if err != nil {
		return nil, err
	}

	// Trả về response
	return &v1.UpdateRes{
		Id:       int64(updatedUserRole.ID),
		UserId:   int64(updatedUserRole.UserID),
		RoleId:   int64(updatedUserRole.RoleID),
		UpdateAt: gtime.New(updatedUserRole.UpdatedAt),
		UpdateBy: int64(updatedUserRole.UpdatedBy),
	}, nil
}

func (u *userRoleService) Delete(ctx context.Context, id int64) error {
	err := u.userRoleRepo.Delete(ctx, int(id))
	if err != nil {
		return err
	}
	return nil
}
