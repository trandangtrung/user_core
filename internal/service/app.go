package service

import (
	"context"
	v1 "demo/api/app/v1"
	"demo/internal/entity"
	"demo/internal/repository"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	AppService interface {
		Get(ctx context.Context, req *v1.GetReq) (*v1.GetRes, error)
		Create(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error)
		Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error)
		Delete(ctx context.Context, id int64) error
	}

	appService struct {
		appRepo repository.AppRepository
	}
)

func NewAppService(appRepo repository.AppRepository) AppService {
	return &appService{
		appRepo: appRepo,
	}
}

func (s *appService) Get(ctx context.Context, req *v1.GetReq) (*v1.GetRes, error) {
	if req.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid ID")
	}

	app, err := s.appRepo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to get platform")
	}
	if app == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "platform not found")
	}

	return &v1.GetRes{
		Id:   int64(app.ID),
		Name: app.Name,
	}, nil
}

func (s *appService) Create(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error) {
	if req.Name == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "name must not be empty")
	}

	newApp := &entity.App{
		Name:   req.Name,
		Config: req.Config,
	}

	created, err := s.appRepo.Create(ctx, newApp)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to create platform")
	}

	return &v1.CreateRes{
		Id:     int64(created.ID),
		Name:   created.Name,
		Config: created.Config,
	}, nil
}

func (s *appService) Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error) {
	if req.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid ID")
	}
	if req.Name == "" && req.Config == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "name and config must not be empty")
	}

	existingApp, err := s.appRepo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to get app")
	}

	if req.Name != "" {
		existingApp.Name = req.Name
	}
	if req.Config != "" {
		existingApp.Config = req.Config
	}

	updatePlatform, err := s.appRepo.Update(ctx, existingApp)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to update app")
	}

	return &v1.UpdateRes{
		Id:       int64(updatePlatform.ID),
		Name:     updatePlatform.Name,
		Config:   updatePlatform.Config,
		UpdateAt: gtime.NewFromTime(updatePlatform.UpdatedAt),
	}, nil
}

func (s *appService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "invalid ID")
	}

	err := s.appRepo.Delete(ctx, id)
	if err != nil {
		return gerror.WrapCode(gcode.CodeInternalError, err, "failed to delete platform")
	}

	return nil
}
