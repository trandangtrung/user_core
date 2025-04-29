package service

import (
	"context"
	"errors"
	"fmt"
	v1 "strongbody-api/api/user/v1"
	"strongbody-api/internal/entity"
	"strongbody-api/internal/repository"
	utils "strongbody-api/utility"
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
	//Start transaction
	tx := u.userRepo.BeginTx(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//Check if email already exists
	existingUser, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if existingUser != nil {
		tx.Rollback()
		return nil, errors.New("email already exists")
	}
	//Hash password
	passwordHashed, err := utils.HashPassword(req.Password)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//Create user
	newUser := &entity.User{
		Email:          req.Email,
		PasswordHashed: passwordHashed,
	}
	createdUser, err := u.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Gán role và apps trong cùng transaction
	if err := u.userRepo.AssignRole(ctx, tx, newUser.ID, req.Role); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := u.userRepo.AssignApps(ctx, tx, newUser.ID, req.Apps); err != nil {
		tx.Rollback()
		return nil, err
	}

	return &v1.CreateRes{
		Status:  "ok",
		ID:      int64(createdUser.ID),
		Message: "User created successfully",
	}, nil
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
