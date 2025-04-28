package repository

import (
	"context"
	"demo/internal/entity"

	"gorm.io/gorm"
)

type (
	UserAppRepository interface {
		Create(ctx context.Context, userApp *entity.UserApps) (*entity.UserApps, error)
		GetById(ctx context.Context, id int) (*entity.UserApps, error)
		Update(ctx context.Context, userApp *entity.UserApps) (*entity.UserApps, error)
		Delete(ctx context.Context, userApp *entity.UserApps) error
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

func (r *userAppRepository) Create(ctx context.Context, userApp *entity.UserApps) (*entity.UserApps, error) {
	err := r.db.WithContext(ctx).Create(userApp).Error
	if err != nil {
		return nil, err
	}
	return userApp, nil
}

func (r *userAppRepository) GetById(ctx context.Context, id int) (*entity.UserApps, error) {
	var userApp entity.UserApps
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&userApp).Error
	if err != nil {
		return nil, err
	}
	return &userApp, nil
}

func (r *userAppRepository) Update(ctx context.Context, userApp *entity.UserApps) (*entity.UserApps, error) {
	err := r.db.WithContext(ctx).Save(userApp).Error
	if err != nil {
		return nil, err
	}
	return userApp, nil
}

func (r *userAppRepository) Delete(ctx context.Context, userApp *entity.UserApps) error {
	return r.db.WithContext(ctx).Delete(userApp).Error
}
