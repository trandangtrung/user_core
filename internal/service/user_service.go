package service

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "github.com/quannv/strongbody-api/api/user/v1"
	"github.com/quannv/strongbody-api/internal/repository"
	utils "github.com/quannv/strongbody-api/utility"
	rescode "github.com/quannv/strongbody-api/utility/resCode"
)

type (
	UserService interface {
		CreateByAdmin(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error)
		GetByID(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
		ListUsers(ctx context.Context, req *v1.ListUsersReq) (res *v1.ListUsersRes, err error)
		Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
		Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	}
	userService struct {
		userRepo repository.UserRepository
		roleRepo repository.RoleRepository
		appRepo  repository.AppRepository
	}
)

func NewUserService(userRepo repository.UserRepository, roleRepo repository.RoleRepository, appRepo repository.AppRepository) UserService {
	return &userService{
		userRepo: userRepo,
		roleRepo: roleRepo,
		appRepo:  appRepo,
	}
}

func (u *userService) CreateByAdmin(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error) {
	_, err := u.userRepo.CreateUserByAdmin(ctx, req)

	if err != nil {
		return nil, gerror.WrapCode(rescode.UserCreateFailed, err, "failed to create user")
	}
	return &v1.CreateRes{
		Status: "ok",
	}, nil
}

func (u *userService) GetByID(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	user, err := u.userRepo.GetUserByID(ctx, uint(req.Id))
	if err != nil {
		return nil, gerror.NewCode(rescode.UserGetFailed, "failed to get user")
	}

	return &v1.GetRes{
		Id:    user.ID,
		Email: user.Email,
	}, nil
}

func (u *userService) ListUsers(ctx context.Context, req *v1.ListUsersReq) (res *v1.ListUsersRes, err error) {
	return &v1.ListUsersRes{}, nil
}

func (u *userService) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	user, err := u.userRepo.GetUserByID(ctx, uint(req.Id))
	if err != nil {
		return nil, gerror.NewCode(rescode.UserNotFound, "user is not exist")
	}

	passwordHashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, gerror.NewCode(rescode.HashPasswordFailed, "hash password failed")
	}

	user.Email = req.Email
	user.PasswordHashed = passwordHashed

	_, err = u.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, gerror.WrapCode(rescode.UserUpdateFailed, err, "failed to update user")
	}
	return &v1.UpdateRes{
		Id:        user.ID,
		Email:     user.Email,
		Role:      uint(user.Roles[0].ID),
		Apps:      []uint{uint(user.Apps[0].ID)},
		UpdatedBy: *user.UpdatedBy,
		UpdatedAt: gtime.NewFromTime(user.UpdatedAt),
	}, nil
}

func (u *userService) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = u.userRepo.DeleteUser(ctx, uint(req.Id))
	if err != nil {
		return nil, gerror.WrapCode(rescode.UserDeleteFailed, err, "failed to delete user")
	}

	return &v1.DeleteRes{
		Status: "ok",
	}, nil
}
