package repository

import (
	"context"
	"strongbody-api/internal/entity"

	"gorm.io/gorm"
)

type (
	UserRoleRepository interface {
		Create(ctx context.Context, userRole *entity.UserRoles) (*entity.UserRoles, error)
		GetById(ctx context.Context, id int) (*entity.UserRoles, error)
		Update(ctx context.Context, userRole *entity.UserRoles) (*entity.UserRoles, error)
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

func (r *userRoleRepository) Create(ctx context.Context, userRole *entity.UserRoles) (*entity.UserRoles, error) {
	err := r.db.WithContext(ctx).Create(userRole).Error
	if err != nil {
		return nil, err
	}
	return userRole, nil
}

func (r *userRoleRepository) GetById(ctx context.Context, id int) (*entity.UserRoles, error) {
	var userRole entity.UserRoles
	err := r.db.WithContext(ctx).First(&userRole, id).Error
	if err != nil {
		return nil, err
	}
	return &userRole, nil
}

func (r *userRoleRepository) Update(ctx context.Context, userRole *entity.UserRoles) (*entity.UserRoles, error) {
	err := r.db.WithContext(ctx).Save(userRole).Error
	if err != nil {
		return nil, err
	}
	return userRole, nil
}

func (r *userRoleRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.UserRoles{}, id).Error
}
