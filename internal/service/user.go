package service

import (
	"context"
	v1 "demo/api/user/v1"
	"demo/internal/repository"
	utils "demo/utility"
)

type (
	UserService interface {
		CreateByAdmin(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error)
		GetByID(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error)
		Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
		Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	}
	userService struct {
		userRepo repository.UserRepository
	}
)

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) CreateByAdmin(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error) {
	// check email

	// create user

	// assign user - role

	// assign user - app

	return nil, nil
}

func (u *userService) GetByID(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	user, err := u.userRepo.GetByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &v1.GetRes{
		Id:    int64(user.ID),
		Email: user.Email,
	}, nil
}

func (u *userService) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	user, err := u.userRepo.GetByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	passwordHashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user.Email = req.Email
	user.PasswordHashed = passwordHashed

	_, err = u.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateRes{}, nil
}

func (u *userService) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = u.userRepo.Delete(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRes{}, nil
}
