package service

import (
	"context"

	v1 "github.com/quannv/strongbody-api/api/app/v1"
	"github.com/quannv/strongbody-api/internal/entity"
	"github.com/quannv/strongbody-api/internal/repository"
	rescode "github.com/quannv/strongbody-api/utility/resCode"

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
		appRepo repository.AppRepository
	}
)

func NewAppService(appRepo repository.AppRepository) AppService {
	return &appService{
		appRepo: appRepo,
	}
}

func (s *appService) Get(ctx context.Context, req *v1.GetReq) (*v1.GetRes, error) {

	app, err := s.appRepo.GetAppByID(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(rescode.AppGetFailed, err, "failed to get app")
	}
	if app == nil {
		return nil, gerror.NewCode(rescode.AppNotFound, "platform not found")
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

	created, err := s.appRepo.CreateApp(ctx, newApp)
	if err != nil {
		return nil, gerror.WrapCode(rescode.AppCreateFailed, err, "failed to create app")
	}

	return &v1.CreateRes{
		Id:     created.ID,
		Name:   created.Name,
		Config: created.Config,
	}, nil
}

func (s *appService) Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error) {
	// check if the app exists
	existingApp, err := s.appRepo.GetAppByID(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(rescode.AppGetFailed, err, "failed to get app")
	}

	updatePlatform, err := s.appRepo.UpdateApp(ctx, existingApp)
	if err != nil {
		return nil, gerror.WrapCode(rescode.AppUpdateFailed, err, "failed to update app")
	}

	return &v1.UpdateRes{
		Id:       updatePlatform.ID,
		Name:     updatePlatform.Name,
		Config:   updatePlatform.Config,
		UpdateAt: gtime.NewFromTime(updatePlatform.UpdatedAt),
	}, nil
}

func (s *appService) Delete(ctx context.Context, id uint) error {
	err := s.appRepo.DeleteApp(ctx, id)
	if err != nil {
		return gerror.WrapCode(rescode.AppDeleteFailed, err, "failed to delete app")
	}

	return nil
}
