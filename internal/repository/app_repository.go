package repository

import (
	"context"
	"errors"

	"github.com/quannv/strongbody-api/internal/entity"

	"gorm.io/gorm"
)

type (
	AppRepository interface {
		CreateApp(ctx context.Context, app *entity.App) (*entity.App, error)
		GetAppByID(ctx context.Context, id uint) (*entity.App, error)
		UpdateApp(ctx context.Context, app *entity.App) (*entity.App, error)
		DeleteApp(ctx context.Context, id uint) error
	}
	appRepository struct {
		db *gorm.DB
	}
)

func NewAppRepository(db *gorm.DB) AppRepository {
	return &appRepository{db: db}
}

func (a *appRepository) CreateApp(ctx context.Context, app *entity.App) (*entity.App, error) {
	if err := a.db.WithContext(ctx).Create(app).Error; err != nil {
		return nil, err
	}
	return app, nil
}

func (a *appRepository) GetAppByID(ctx context.Context, id uint) (*entity.App, error) {
	var app entity.App
	if err := a.db.WithContext(ctx).First(&app, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &app, nil
}

func (a *appRepository) UpdateApp(ctx context.Context, app *entity.App) (*entity.App, error) {
	if err := a.db.WithContext(ctx).Save(app).Error; err != nil {
		return nil, err
	}
	return app, nil
}

func (a *appRepository) DeleteApp(ctx context.Context, id uint) error {
	if err := a.db.WithContext(ctx).Delete(&entity.App{}, id).Error; err != nil {
		return err
	}
	return nil
}
