package tokenL

import (
	"context"
	v1 "demo/api/token/v1"
	"demo/internal/dao"
	"demo/internal/model/do"
)

type Token struct {
}

func (s *Token) Create(ctx context.Context, in *v1.CreateReq) error {
	_, err := dao.Tokens.Ctx(ctx).Data(do.Tokens{
		UserId:       in.User_id,
		RefreshToken: in.RefreshToken,
		Scope:        in.Scope,
	}).Insert()

	if err != nil {
		return err
	}

	return nil
}

func (s *Token) Update(ctx context.Context, id int64, in *v1.UpdateReq) (*v1.UpdateRes, error) {
	_, err := dao.Tokens.Ctx(ctx).Where("id", id).Data(do.Tokens{
		UserId:       in.User_id,
		RefreshToken: in.RefreshToken,
		Scope:        in.Scope,
	}).Update()

	if err != nil {
		return nil, err
	}

	return &v1.UpdateRes{}, nil
}

func (s *Token) Delete(ctx context.Context, id int64) error {
	db := dao.Tokens.Ctx(ctx).Where("id", id)

	_, err := db.Delete()

	if err != nil {
		return err
	}

	return nil
}
