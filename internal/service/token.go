package service

import (
	"context"
	v1 "demo/api/token/v1"
	"demo/internal/repository"
)

type (
	TokenService interface {
		Create(ctx context.Context, in *v1.CreateReq) error
		Update(ctx context.Context, id int64, in *v1.UpdateReq) (*v1.UpdateRes, error)
		Delete(ctx context.Context, id int64) error
	}
	tokenService struct {
		tokenRepo repository.TokenRepository
	}
)

func NewTokenService(tokenRepo repository.TokenRepository) TokenService {
	return &tokenService{
		tokenRepo: tokenRepo,
	}
}

func (l *tokenService) Create(ctx context.Context, in *v1.CreateReq) error {
	return nil
}

func (l *tokenService) Update(ctx context.Context, id int64, in *v1.UpdateReq) (*v1.UpdateRes, error) {
	return nil, nil
}

func (l *tokenService) Delete(ctx context.Context, id int64) error {
	return nil
}
