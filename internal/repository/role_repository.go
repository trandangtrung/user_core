package repository

import (
	"context"
	"errors"

	"github.com/quannv/strongbody-api/internal/entity"

	"gorm.io/gorm"
)

type (
	RoleRepository interface {
		CreateRole(ctx context.Context, role *entity.Role) (*entity.Role, error)
		GetRoleByID(ctx context.Context, id uint) (*entity.Role, error)
		GetRolesByUserIDAndAppName(ctx context.Context, userID uint, appName string) ([]*entity.Role, error)
		UpdateRole(ctx context.Context, role *entity.Role) (*entity.Role, error)
		DeleteRole(ctx context.Context, id uint) error
	}
	roleRepository struct {
		db *gorm.DB
	}
)

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) CreateRole(ctx context.Context, role *entity.Role) (*entity.Role, error) {
	if err := r.db.WithContext(ctx).Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) GetRoleByID(ctx context.Context, id uint) (*entity.Role, error) {
	var role entity.Role
	if err := r.db.WithContext(ctx).First(&role, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetRolesByUserIDAndAppName(ctx context.Context, userID uint, appName string) ([]*entity.Role, error) {
	var roles []*entity.Role

	err := r.db.WithContext(ctx).
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Joins("JOIN apps ON apps.id = roles.app_id").
		Where("user_roles.user_id = ? AND apps.name = ?", userID, appName).
		Find(&roles).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) UpdateRole(ctx context.Context, role *entity.Role) (*entity.Role, error) {
	if err := r.db.WithContext(ctx).Save(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) DeleteRole(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}
