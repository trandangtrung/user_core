package repository

import (
	"context"
	"demo/internal/entity"

	"gorm.io/gorm"
)

type (
	TokenRepository interface {
		Create(ctx context.Context, token *entity.Token) error
		GetByID(ctx context.Context, id uint) (*entity.Token, error)
		GetByUserID(ctx context.Context, userID uint) (*entity.Token, error)
		GetByToken(ctx context.Context, token string) (*entity.Token, error)
		GetUser(ctx context.Context, id uint) (*entity.User, error)
		Update(ctx context.Context, token *entity.Token) error
		Delete(ctx context.Context, id uint) error
	}

	tokenRepository struct {
		db *gorm.DB
	}
)

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) Create(ctx context.Context, token *entity.Token) error {
	if err := r.db.WithContext(ctx).Create(token).Error; err != nil {
		return err
	}
	return nil
}

func (r *tokenRepository) GetByID(ctx context.Context, id uint) (*entity.Token, error) {
	var token entity.Token
	if err := r.db.WithContext(ctx).First(&token, id).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *tokenRepository) GetByUserID(ctx context.Context, userID uint) (*entity.Token, error) {
	var token entity.Token
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *tokenRepository) GetByToken(ctx context.Context, token string) (*entity.Token, error) {
	var t entity.Token
	if err := r.db.WithContext(ctx).Where("refresh_token = ?", token).First(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *tokenRepository) GetUser(ctx context.Context, id uint) (*entity.User, error) {
	var token entity.Token
	if err := r.db.WithContext(ctx).Preload("User").First(&token, id).Error; err != nil {
		return nil, err
	}
	return token.User, nil
}

func (r *tokenRepository) Update(ctx context.Context, token *entity.Token) error {
	if err := r.db.WithContext(ctx).Save(token).Error; err != nil {
		return err
	}
	return nil
}

func (r *tokenRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Token{}, id).Error; err != nil {
		return err
	}
	return nil
}
