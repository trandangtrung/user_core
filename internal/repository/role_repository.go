package repository

import (
	"context"
	"demo/internal/entity"

	"gorm.io/gorm"
)

type (
	RoleRepository interface {
		Create(ctx context.Context, role *entity.Role) error
		GetByID(ctx context.Context, id uint) (*entity.Role, error)
		GetPlatform(ctx context.Context, id uint) (*entity.Platform, error)
		Update(ctx context.Context, role *entity.Role) error
		Delete(ctx context.Context, id uint) error
	}

	roleRepository struct {
		db *gorm.DB
	}
)

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(ctx context.Context, role *entity.Role) error {
	if err := r.db.WithContext(ctx).Create(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) GetByID(ctx context.Context, id uint) (*entity.Role, error) {
	var role entity.Role
	if err := r.db.WithContext(ctx).First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetPlatform(ctx context.Context, id uint) (*entity.Platform, error) {
	var role entity.Role
	if err := r.db.WithContext(ctx).Preload("Platform").First(&role, id).Error; err != nil {
		return nil, err
	}
	return role.Platform, nil
}

func (r *roleRepository) Update(ctx context.Context, role *entity.Role) error {
	if err := r.db.WithContext(ctx).Save(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}
