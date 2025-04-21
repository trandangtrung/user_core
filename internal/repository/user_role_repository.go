package repository

import (
	"context"
	"demo/internal/entity"

	"gorm.io/gorm"
)

type (
	UserRoleRepository interface {
		Create(ctx context.Context, userRole *entity.UserRole) error
		GetById(ctx context.Context, id int) (*entity.UserRole, error)
		Update(ctx context.Context, userRole *entity.UserRole) error
		Delete(ctx context.Context, id int) error
	}

	userRoleRepository struct {
		db *gorm.DB
	}
)

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{
		db: db,
	}
}

func (r *userRoleRepository) Create(ctx context.Context, userRole *entity.UserRole) error {
	return r.db.WithContext(ctx).Create(userRole).Error
}

func (r *userRoleRepository) GetById(ctx context.Context, id int) (*entity.UserRole, error) {
	var userRole entity.UserRole
	err := r.db.WithContext(ctx).First(&userRole, id).Error
	if err != nil {
		return nil, err
	}
	return &userRole, nil
}

func (r *userRoleRepository) Update(ctx context.Context, userRole *entity.UserRole) error {
	return r.db.WithContext(ctx).Save(userRole).Error
}

func (r *userRoleRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.UserRole{}, id).Error
}
