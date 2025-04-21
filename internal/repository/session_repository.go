package repository

import (
	"context"
	"demo/internal/entity"

	"gorm.io/gorm"
)

type (
	SessionRepository interface {
		Create(ctx context.Context, session *entity.Session) error
		GetByID(ctx context.Context, id uint) (*entity.Session, error)
		GetUser(ctx context.Context, id uint) (*entity.User, error)
		Update(ctx context.Context, session *entity.Session) error
		Delete(ctx context.Context, id uint) error
	}

	sessionRepository struct {
		db *gorm.DB
	}
)

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) Create(ctx context.Context, session *entity.Session) error {
	if err := r.db.WithContext(ctx).Create(session).Error; err != nil {
		return err
	}
	return nil
}

func (r *sessionRepository) GetByID(ctx context.Context, id uint) (*entity.Session, error) {
	var session entity.Session
	if err := r.db.WithContext(ctx).First(&session, id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *sessionRepository) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	var session entity.Session
	if err := r.db.WithContext(ctx).Preload("User").First(&session, id).Error; err != nil {
		return nil, err
	}
	return session.User, nil
}

func (r *sessionRepository) Update(ctx context.Context, session *entity.Session) error {
	if err := r.db.WithContext(ctx).Save(session).Error; err != nil {
		return err
	}
	return nil
}

func (r *sessionRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Session{}, id).Error; err != nil {
		return err
	}
	return nil
}
