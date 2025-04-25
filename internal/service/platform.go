package service

import (
	"context"
	"demo/api/platform/v1"
	"demo/internal/entity"
	"demo/internal/repository"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	PlatformService interface {
		Get(ctx context.Context, req *v1.GetReq) (*v1.GetRes, error)
		Create(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error)
		Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error)
		Delete(ctx context.Context, id int64) error
	}

	platformService struct {
		platformRepo repository.PlatformRepository
	}
)

func NewPlatformService(platformRepo repository.PlatformRepository) PlatformService {
	return &platformService{
		platformRepo: platformRepo,
	}
}

func (s *platformService) Get(ctx context.Context, req *v1.GetReq) (*v1.GetRes, error) {
	if req.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid ID")
	}

	platform, err := s.platformRepo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to get platform")
	}
	if platform == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "platform not found")
	}

	return &v1.GetRes{
		Id:   int64(platform.ID),
		Name: platform.Name,
	}, nil
}

func (s *platformService) Create(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error) {
	if req.Name == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "name must not be empty")
	}

	newPlatform := &entity.Platform{
		Name:   req.Name,
		Config: req.Config,
	}

	created, err := s.platformRepo.Create(ctx, newPlatform)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to create platform")
	}

	return &v1.CreateRes{
		Id:     int64(created.ID),
		Name:   created.Name,
		Config: created.Config,
	}, nil
}

func (s *platformService) Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error) {
	if req.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid ID")
	}
	if req.Name == "" && req.Config == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "name and config must not be empty")
	}

	existingPlatform, err := s.platformRepo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to get platform")
	}

	if req.Name != "" {
		existingPlatform.Name = req.Name
	}
	if req.Config != "" {
		existingPlatform.Config = req.Config
	}

	updatePlatform, err := s.platformRepo.Update(ctx, existingPlatform)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to update platform")
	}

	return &v1.UpdateRes{
		Id:       int64(updatePlatform.ID),
		Name:     updatePlatform.Name,
		Config:   updatePlatform.Config,
		UpdateAt: gtime.NewFromTime(updatePlatform.UpdatedAt),
	}, nil
}

func (s *platformService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "invalid ID")
	}

	err := s.platformRepo.Delete(ctx, id)
	if err != nil {
		return gerror.WrapCode(gcode.CodeInternalError, err, "failed to delete platform")
	}

	return nil
}
