package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Role struct{}

type GetReq struct {
	g.Meta `path:"/role/{id}" method:"get" tags:"role" summary:"Get role"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     uint   `json:"id" v:"required"`
}

type GetRes struct {
	Id          uint        `json:"id"`
	AppId       uint        `json:"app_id"`
	Name        string      `json:"name"`
	Key         string      `json:"key"`
	Description string      `json:"description"`
	CreatedAt   *gtime.Time `json:"created_at"`
	CreatedBy   *uint       `json:"created_by"`
	UpdatedAt   *gtime.Time `json:"updated_at"`
	UpdatedBy   uint        `json:"updated_by"`
}

type CreateReq struct {
	g.Meta      `path:"/role" method:"post" tags:"role" summary:"Create role"`
	Scope       string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Key         string `json:"key" v:"required"`
	AppId       uint   `json:"app_id" v:"required"`
	Name        string `json:"name" v:"required|length:1,20"`
	Description string `json:"description" v:"required|length:1,300"`
}

type CreateRes struct {
	Id          uint        `json:"id"`
	AppId       uint        `json:"app_id"`
	Key         string      `json:"key"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CreatedAt   *gtime.Time `json:"created_at"`
	CreatedBy   *uint       `json:"created_by"`
}

type UpdateReq struct {
	g.Meta      `path:"/role/{id}" method:"patch" tags:"role" summary:"Update role"`
	Scope       string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id          uint   `json:"id" v:"required"`
	Name        string `json:"name" v:"required|length:1,20"`
	Key         string `json:"key" v:"required"`
	AppId       uint   `json:"app_id" v:"required"`
	Description string `json:"description" v:"required|length:1,300"`
}

type UpdateRes struct {
	Id          uint        `json:"id"`
	AppId       uint        `json:"app_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	UpdatedAt   *gtime.Time `json:"updated_at"`
	UpdatedBy   uint        `json:"updated_by"`
}

type DeleteReq struct {
	g.Meta `path:"/role/{id}" method:"delete" tags:"role" summary:"Delete role"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     uint   `json:"id" v:"required"`
}

type DeleteRes struct {
	Status string `json:"status"`
}
