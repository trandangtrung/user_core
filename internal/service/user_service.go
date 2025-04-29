package service

import (
	"context"
	v1 "demo/api/user/v1"
	"demo/internal/entity"
	"demo/internal/repository"
	utils "demo/utility"
	"errors"
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
	// check if the email is already exist
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("email is already exist")
	}
	if user != nil {
		return nil, errors.New("email is already exist")
	}

	// create user
	passwordHashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user = &entity.User{
		UserName:       req.UserName,
		Email:          req.Email,
		PasswordHashed: passwordHashed,
		Mobile:         req.Mobile,
	}
	user, err = u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// create user role
	role, err := u.userRepo.GetRoleByID(ctx, req.Role)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errors.New("role is not exist")
	}
	err = u.userRepo.CreateUserRole(ctx, user, role)
	if err != nil {
		return nil, err
	}

	// create user app
	var apps []*entity.App
	for _, appID := range req.Apps {
		app, err := u.userRepo.GetAppByID(ctx, appID)
		if err != nil {
			return nil, err
		}
		if app == nil {
			return nil, errors.New("app is not exist")
		}
		apps = append(apps, app)
	}
	for _, app := range apps {
		err = u.userRepo.CreateUserApp(ctx, user, app)
		if err != nil {
			return nil, err
		}
	}

	return &v1.CreateRes{
		Status: "ok",
	}, nil
}

func (u *userService) GetByID(ctx context.Context, req *v1.GetReq) (res *v1.GetRes, err error) {
	user, err := u.userRepo.GetUserByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &v1.GetRes{
		Id:    user.ID,
		Email: user.Email,
	}, nil
}

func (u *userService) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	user, err := u.userRepo.GetUserByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	passwordHashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user.Email = req.Email
	user.PasswordHashed = passwordHashed

	_, err = u.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateRes{}, nil
}

func (u *userService) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = u.userRepo.DeleteUser(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &v1.DeleteRes{
		Status: "ok",
	}, nil
}
