package repository

import (
	"context"
	"demo/internal/entity"

	"gorm.io/gorm"
)

type (
	UserAppRepository interface {
		Create(ctx context.Context, userApp *entity.UserApp) (*entity.UserApp, error)
		GetById(ctx context.Context, id int) (*entity.UserApp, error)
		GetByUserId(ctx context.Context, userId int) ([]*entity.UserApp, error)
		GetUser(ctx context.Context, id int) (*entity.User, error)
		Update(ctx context.Context, userApp *entity.UserApp) (*entity.UserApp, error)
		Delete(ctx context.Context, userApp *entity.UserApp) error
	}

	userAppRepository struct {
		db *gorm.DB
	}
)

func NewUserAppRepository(db *gorm.DB) UserAppRepository {
	return &userAppRepository{
		db: db,
	}
}

func (r *userAppRepository) Create(ctx context.Context, userApp *entity.UserApp) (*entity.UserApp, error) {
	err := r.db.WithContext(ctx).Create(userApp).Error
	if err != nil {
		return nil, err
	}
	return userApp, nil
}

func (r *userAppRepository) GetById(ctx context.Context, id int) (*entity.UserApp, error) {
	var userApp entity.UserApp
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&userApp).Error
	if err != nil {
		return nil, err
	}
	return &userApp, nil
}

func (r *userAppRepository) GetByUserId(ctx context.Context, userId int) ([]*entity.UserApp, error) {
	var userApp []*entity.UserApp
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&userApp).Error
	if err != nil {
		return nil, err
	}
	return userApp, nil
}

func (r *userAppRepository) GetUser(ctx context.Context, id int) (*entity.User, error) {
	var userApp entity.UserApp
	err := r.db.WithContext(ctx).Preload("User").Where("id = ?", id).First(&userApp).Error
	if err != nil {
		return nil, err
	}
	return userApp.User, nil
}

func (r *userAppRepository) Update(ctx context.Context, userApp *entity.UserApp) (*entity.UserApp, error) {
	err := r.db.WithContext(ctx).Save(userApp).Error
	if err != nil {
		return nil, err
	}
	return userApp, nil
}

func (r *userAppRepository) Delete(ctx context.Context, userApp *entity.UserApp) error {
	return r.db.WithContext(ctx).Delete(userApp).Error
}
