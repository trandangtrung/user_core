package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type platform struct {
}

type GetReq struct {
	g.Meta `path:"/platform/{id}" method:"get" tags:"platform" summary:"Get platform"`
	Id     int64 `json:"id" v:"required"`
}

type GetRes struct {
	Id       int64       `json:"id"`
	Name     string      `json:"name" v:"required"`
	Config   string      `json:"config" v:"required"`
	CreateAt *gtime.Time `json:"created_at"`
	CreateBy int64       `json:"created_by"`
	UpdateAt *gtime.Time `json:"updated_at"`
	UpdateBy int64       `json:"updated_by"`
}

type CreateReq struct {
	g.Meta `path:"/platform" method:"post" tags:"platform" summary:"Create platform"`
	Name   string `json:"name" v:"required"`
	Config string `json:"config" v:"required"`
}

type CreateRes struct {
	Id       int64       `json:"id" v:"required"`
	Name     string      `json:"name" v:"required"`
	Config   string      `json:"config" v:"required"`
	CreateAt *gtime.Time `json:"created_at"`
	CreateBy int64       `json:"created_by"`
}

type UpdateReq struct {
	g.Meta `path:"/platform/{id}" method:"patch" tags:"platform" summary:"Update platform"`
	Id     int64  `json:"id" v:"required"`
	Name   string `json:"name" v:"required"`
	Config string `json:"config" v:"required"`
}

type UpdateRes struct {
	Id       int64       `json:"id"`
	Name     string      `json:"name" v:"required"`
	Config   string      `json:"config" v:"required"`
	UpdateAt *gtime.Time `json:"updated_at"`
	UpdateBy int64       `json:"updated_by"`
}

type DeleteReq struct {
	g.Meta `path:"/platform/{id}" method:"delete" tags:"platform" summary:"Delete platform"`
	Id     int64 `json:"id" v:"required"`
}

type DeleteRes struct {
	Status string `json:"status"`
}
