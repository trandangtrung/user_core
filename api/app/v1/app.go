package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetReq struct {
	g.Meta `path:"/app/{id}" method:"get" tags:"app" summary:"Get app"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     uint   `json:"id" v:"required"`
}

type GetRes struct {
	Id       uint        `json:"id"`
	Name     string      `json:"name"`
	Config   string      `json:"config"`
	CreateAt *gtime.Time `json:"created_at"`
	CreateBy uint        `json:"created_by"`
	UpdateAt *gtime.Time `json:"updated_at"`
	UpdateBy uint        `json:"updated_by"`
}

type CreateReq struct {
	g.Meta `path:"/app" method:"post" tags:"app" summary:"Create app"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Name   string `json:"name" v:"required|length:1,20"`
	Config string `json:"config" v:"required"`
}

type CreateRes struct {
	Id       uint        `json:"id" v:"required"`
	Name     string      `json:"name" v:"required"`
	Config   string      `json:"config" v:"required"`
	CreateAt *gtime.Time `json:"created_at"`
	CreateBy uint        `json:"created_by"`
}

type UpdateReq struct {
	g.Meta `path:"/app/{id}" method:"patch" tags:"app" summary:"Update app"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     uint   `json:"id" v:"required"`
	Name   string `json:"name" v:"required|length:1,20"`
	Config string `json:"config" v:"required"`
}

type UpdateRes struct {
	Id       uint        `json:"id"`
	Name     string      `json:"name"`
	Config   string      `json:"config"`
	UpdateAt *gtime.Time `json:"updated_at"`
	UpdateBy uint        `json:"updated_by"`
}

type DeleteReq struct {
	g.Meta `path:"/app/{id}" method:"delete" tags:"app" summary:"Delete app"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     uint   `json:"id" v:"required"`
}

type DeleteRes struct {
	Status string `json:"status"`
}
