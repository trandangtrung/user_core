package repository

import (
	"context"
	"demo/internal/entity"

	"gorm.io/gorm"
)

type (
	UserPlatformRepository interface {
		Create(ctx context.Context, userPlatform *entity.UserPlatform) (*entity.UserPlatform, error)
		GetById(ctx context.Context, id int) (*entity.UserPlatform, error)
		GetByUserId(ctx context.Context, userId int) ([]*entity.UserPlatform, error)
		GetUser(ctx context.Context, id int) (*entity.User, error)
		Update(ctx context.Context, userPlatform *entity.UserPlatform) (*entity.UserPlatform, error)
		Delete(ctx context.Context, userPlatform *entity.UserPlatform) error
	}

	userPlatformRepository struct {
		db *gorm.DB
	}
)

func NewUserPlatformRepository(db *gorm.DB) UserPlatformRepository {
	return &userPlatformRepository{
		db: db,
	}
}

func (r *userPlatformRepository) Create(ctx context.Context, userPlatform *entity.UserPlatform) (*entity.UserPlatform, error) {
	err := r.db.WithContext(ctx).Create(userPlatform).Error
	if err != nil {
		return nil, err
	}
	return userPlatform, nil
}

func (r *userPlatformRepository) GetById(ctx context.Context, id int) (*entity.UserPlatform, error) {
	var userPlatform entity.UserPlatform
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&userPlatform).Error
	if err != nil {
		return nil, err
	}
	return &userPlatform, nil
}

func (r *userPlatformRepository) GetByUserId(ctx context.Context, userId int) ([]*entity.UserPlatform, error) {
	var userPlatforms []*entity.UserPlatform
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&userPlatforms).Error
	if err != nil {
		return nil, err
	}
	return userPlatforms, nil
}

func (r *userPlatformRepository) GetUser(ctx context.Context, id int) (*entity.User, error) {
	var userPlatform entity.UserPlatform
	err := r.db.WithContext(ctx).Preload("User").Where("id = ?", id).First(&userPlatform).Error
	if err != nil {
		return nil, err
	}
	return userPlatform.User, nil
}

func (r *userPlatformRepository) Update(ctx context.Context, userPlatform *entity.UserPlatform) (*entity.UserPlatform, error) {
	err := r.db.WithContext(ctx).Save(userPlatform).Error
	if err != nil {
		return nil, err
	}
	return userPlatform, nil
}

func (r *userPlatformRepository) Delete(ctx context.Context, userPlatform *entity.UserPlatform) error {
	return r.db.WithContext(ctx).Delete(userPlatform).Error
}
