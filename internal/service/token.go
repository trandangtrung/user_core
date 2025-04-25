package service

import (
	"context"
	v1 "demo/api/token/v1"
	"demo/internal/entity"
	"demo/internal/repository"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	TokenService interface {
		Create(ctx context.Context, req *v1.CreateReq) error
		Update(ctx context.Context, id uint, req *v1.UpdateReq) (*v1.UpdateRes, error)
		Delete(ctx context.Context, id uint) error
		Get(ctx context.Context, id uint) (*v1.GetRes, error)
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
func (s *tokenService) Get(ctx context.Context, id uint) (*v1.GetRes, error) {
	if id < 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "Invalid Parameter ID")
	}
	token, err := s.tokenRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return &v1.GetRes{
		Id:           int64(token.ID),
		User_id:      int64(token.UserID),
		RefreshToken: token.RefreshToken,
		Scope:        token.Scope,
	}, nil
}

func (s *tokenService) Create(ctx context.Context, req *v1.CreateReq) error {
	if req == nil || req.RefreshToken == "" || req.Scope == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "refreshToken and scope must not be empty")
	}

	token := &entity.Token{
		UserID:       uint(req.User_id),
		RefreshToken: req.RefreshToken,
		Scope:        req.Scope,
	}

	if _, err := s.tokenRepo.Create(ctx, token); err != nil {
		return gerror.WrapCode(gcode.CodeInternalError, err, "failed to create token")
	}

	return nil
}

func (s *tokenService) Update(ctx context.Context, id uint, req *v1.UpdateReq) (*v1.UpdateRes, error) {
	if id <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid ID")
	}
	if req == nil || req.RefreshToken == "" || req.Scope == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "refreshToken and scope must not be empty")
	}

	token, err := s.tokenRepo.GetByID(ctx, id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to get token")
	}
	if token == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "token not found")
	}

	token.RefreshToken = req.RefreshToken
	token.Scope = req.Scope

	updated, err := s.tokenRepo.Update(ctx, token)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to update token")
	}

	return &v1.UpdateRes{
		Id:           int64(updated.ID),
		User_id:      int64(updated.UserID),
		RefreshToken: updated.RefreshToken,
		Scope:        updated.Scope,
	}, nil
}

func (s *tokenService) Delete(ctx context.Context, id uint) error {
	if id <= 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "invalid ID")
	}

	if err := s.tokenRepo.Delete(ctx, id); err != nil {
		return gerror.WrapCode(gcode.CodeInternalError, err, "failed to delete token")
	}

	return nil
}
