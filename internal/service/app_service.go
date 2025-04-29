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
		Delete(ctx context.Context, id uint) error
	}
	appService struct {
		userRepo repository.UserRepository
	}
)

func NewAppService(userRepo repository.UserRepository) AppService {
	return &appService{
		userRepo: userRepo,
	}
}

func (s *appService) Get(ctx context.Context, req *v1.GetReq) (*v1.GetRes, error) {

	app, err := s.userRepo.GetAppByID(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to get platform")
	}
	if app == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "platform not found")
	}

	return &v1.GetRes{
		Id:   app.ID,
		Name: app.Name,
	}, nil
}

func (s *appService) Create(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error) {

	newApp := &entity.App{
		Name:   req.Name,
		Config: req.Config,
	}

	created, err := s.userRepo.CreateApp(ctx, newApp)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to create platform")
	}

	return &v1.CreateRes{
		Id:     created.ID,
		Name:   created.Name,
		Config: created.Config,
	}, nil
}

func (s *appService) Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error) {
	// check if the app exists
	existingApp, err := s.userRepo.GetAppByID(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to get app")
	}

	updatePlatform, err := s.userRepo.UpdateApp(ctx, existingApp)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to update app")
	}

	return &v1.UpdateRes{
		Id:       updatePlatform.ID,
		Name:     updatePlatform.Name,
		Config:   updatePlatform.Config,
		UpdateAt: gtime.NewFromTime(updatePlatform.UpdatedAt),
	}, nil
}

func (s *appService) Delete(ctx context.Context, id uint) error {
	err := s.userRepo.DeleteApp(ctx, id)
	if err != nil {
		return gerror.WrapCode(gcode.CodeInternalError, err, "failed to delete platform")
	}

	return nil
}
