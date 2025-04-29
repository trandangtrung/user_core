package repository

import (
	"context"
	"strongbody-api/internal/entity"

	"gorm.io/gorm"
)

type (
	RoleRepository interface {
		Create(ctx context.Context, role *entity.Role) (*entity.Role, error)
		GetByID(ctx context.Context, id uint) (*entity.Role, error)
		GetRolesByUserIDAndPlatformName(ctx context.Context, platformName string, userID uint) ([]*entity.Role, error)
		Update(ctx context.Context, role *entity.Role) (*entity.Role, error)
		Delete(ctx context.Context, id uint) error
	}
	roleRepository struct {
		db *gorm.DB
	}
)

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(ctx context.Context, role *entity.Role) (*entity.Role, error) {
	if err := r.db.WithContext(ctx).Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) GetByID(ctx context.Context, id uint) (*entity.Role, error) {
	var role entity.Role
	if err := r.db.WithContext(ctx).First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetRolesByUserIDAndPlatformName(ctx context.Context, platformName string, userID uint) ([]*entity.Role, error) {
	var roles []*entity.Role

	err := r.db.WithContext(ctx).
		Model(&entity.Role{}).
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Joins("JOIN platforms ON platforms.id = roles.platform_id").
		Where("user_roles.user_id = ? AND platforms.name = ?", userID, platformName).
		Preload("Platform").
		Find(&roles).Error

	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) Update(ctx context.Context, role *entity.Role) (*entity.Role, error) {
	if err := r.db.WithContext(ctx).Save(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *roleRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}
