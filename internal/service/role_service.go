package service

import (
	"context"

	v1 "github.com/quannv/strongbody-api/api/role/v1"
	"github.com/quannv/strongbody-api/internal/consts"
	"github.com/quannv/strongbody-api/internal/entity"
	"github.com/quannv/strongbody-api/internal/repository"
	rescode "github.com/quannv/strongbody-api/utility/resCode"
	"github.com/quannv/strongbody-api/utility/token"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	RoleService interface {
		GetByID(ctx context.Context, req *v1.GetReq) (*v1.GetRes, error)
		Create(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error)
		Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error)
		Delete(ctx context.Context, id uint) error
	}
	roleService struct {
		roleRepo repository.RoleRepository
	}
)

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}

func (s *roleService) GetByID(ctx context.Context, req *v1.GetReq) (*v1.GetRes, error) {

	role, err := s.roleRepo.GetRoleByID(ctx, uint(req.Id))
	if err != nil {
		return nil, gerror.WrapCode(rescode.RoleGetFailed, err, "failed to get role")
	}
	if role == nil {
		return nil, gerror.WrapCode(rescode.RoleNotFound, err, "role not found")

	}

	return &v1.GetRes{
		Id:          role.ID,
		AppId:       role.AppID,
		Name:        role.Name,
		Description: role.Description,
		CreatedBy:   role.CreatedBy,
		UpdatedBy:   *role.UpdatedBy,
	}, nil
}

func (s *roleService) Create(ctx context.Context, req *v1.CreateReq) (*v1.CreateRes, error) {
	if req.AppId <= 0 || req.Name == "" || req.Description == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "invalid input")
	}

	newRole := &entity.Role{
		AppID:       req.AppId,
		Name:        req.Name,
		Description: req.Description,
	}

	created, err := s.roleRepo.CreateRole(ctx, newRole)
	if err != nil {
		return nil, gerror.WrapCode(rescode.RoleCreateFailed, err, "failed to create role")
	}

	return &v1.CreateRes{
		Id:          created.ID,
		AppId:       created.AppID,
		Name:        created.Name,
		Description: created.Description,
		CreatedBy:   created.CreatedBy,
	}, nil
}

func (s *roleService) Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error) {
	payload, _ := ctx.Value(consts.AuthorizationKey).(*token.Payload)

	var id int = payload.Id
	convertedID := uint(id)
	pointerID := &convertedID

	existingRole, err := s.roleRepo.GetRoleByID(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(rescode.RoleGetFailed, err, "failed to get role")
	}
	if existingRole == nil {
		return nil, gerror.WrapCode(rescode.RoleNotFound, err, "role not found")
	}

	if req.Name != "" {
		existingRole.Name = req.Name
	}
	if req.Description != "" {
		existingRole.Description = req.Description
	}

	existingRole.UpdatedBy = pointerID

	updated, err := s.roleRepo.UpdateRole(ctx, existingRole)
	if err != nil {
		return nil, gerror.WrapCode(rescode.RoleUpdateFailed, err, "failed to update role")
	}

	return &v1.UpdateRes{
		Id:          updated.ID,
		AppId:       updated.AppID,
		Name:        updated.Name,
		Description: updated.Description,
		UpdatedBy:   *updated.UpdatedBy,
		UpdatedAt:   gtime.NewFromTime(updated.UpdatedAt),
	}, nil
}

func (s *roleService) Delete(ctx context.Context, id uint) error {
	if id <= 0 {
		return gerror.NewCode(gcode.CodeInvalidParameter, "invalid ID")
	}

	err := s.roleRepo.DeleteRole(ctx, id)
	if err != nil {
		return gerror.WrapCode(rescode.RoleDeleteFailed, err, "failed to delete role")
	}

	return nil
}
