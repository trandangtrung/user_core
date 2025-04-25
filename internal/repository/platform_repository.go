package repository

import (
	"context"
	"demo/internal/entity"

	"gorm.io/gorm"
)

type (
	PlatformRepository interface {
		Create(ctx context.Context, platform *entity.Platform) (*entity.Platform, error)
		GetByID(ctx context.Context, id uint) (*entity.Platform, error)
		Update(ctx context.Context, platform *entity.Platform) (*entity.Platform, error)
		Delete(ctx context.Context, id uint) error
	}

	platformRepository struct {
		db *gorm.DB
	}
)

func NewPlatformRepository(db *gorm.DB) PlatformRepository {
	return &platformRepository{db: db}
}

func (r *platformRepository) Create(ctx context.Context, platform *entity.Platform) (*entity.Platform, error) {
	if err := r.db.WithContext(ctx).Create(platform).Error; err != nil {
		return nil, err
	}
	return platform, nil
}

func (r *platformRepository) GetByID(ctx context.Context, id uint) (*entity.Platform, error) {
	var platform entity.Platform
	if err := r.db.WithContext(ctx).First(&platform, id).Error; err != nil {
		return nil, err
	}
	return &platform, nil
}

func (r *platformRepository) Update(ctx context.Context, platform *entity.Platform) (*entity.Platform, error) {
	if err := r.db.WithContext(ctx).Save(platform).Error; err != nil {
		return nil, err
	}
	return platform, nil
}

func (r *platformRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Platform{}, id).Error; err != nil {
		return err
	}
	return nil
}
