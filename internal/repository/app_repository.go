package repository

import (
	"context"
	"demo/internal/entity"

	"gorm.io/gorm"
)

type (
	AppRepository interface {
		Create(ctx context.Context, app *entity.App) (*entity.App, error)
		GetByID(ctx context.Context, id int64) (*entity.App, error)
		Update(ctx context.Context, app *entity.App) (*entity.App, error)
		Delete(ctx context.Context, id int64) error
	}

	appRepository struct {
		db *gorm.DB
	}
)

func NewAppRepository(db *gorm.DB) AppRepository {
	return &appRepository{db: db}
}

func (r *appRepository) Create(ctx context.Context, app *entity.App) (*entity.App, error) {
	if err := r.db.WithContext(ctx).Create(app).Error; err != nil {
		return nil, err
	}
	return app, nil
}

func (r *appRepository) GetByID(ctx context.Context, id int64) (*entity.App, error) {
	var app entity.App
	if err := r.db.WithContext(ctx).First(&app, id).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *appRepository) Update(ctx context.Context, app *entity.App) (*entity.App, error) {
	if err := r.db.WithContext(ctx).Model(&entity.App{}).
		Where("id = ?", app.ID).
		Updates(app).Error; err != nil {
		return nil, err
	}
	return app, nil
}

func (r *appRepository) Delete(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).Delete(&entity.App{}, id).Error; err != nil {
		return err
	}
	return nil
}
