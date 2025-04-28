package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetReq struct {
	g.Meta `path:"/user-app/{id}" method:"get" tags:"user_app" summary:"Get user app"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     int64  `json:"id" v:"required"`
}

type GetRes struct {
	Id       int64       `json:"id"`
	UserId   int64       `json:"userId" v:"required"`
	AppId    int64       `json:"appId" v:"required"`
	CreateAt *gtime.Time `json:"created_at"`
	CreateBy int64       `json:"created_by"`
	UpdateAt *gtime.Time `json:"updated_at"`
	UpdateBy int64       `json:"updated_by"`
}

type CreateReq struct {
	g.Meta `path:"/user-app" method:"post" tags:"user_app" summary:"Create user app"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	UserId int64  `json:"userId" v:"required"`
	AppId  int64  `json:"appId" v:"required"`
}

type CreateRes struct {
	Id       int64       `json:"id" v:"required"`
	UserId   int64       `json:"userId" v:"required"`
	AppId    int64       `json:"appId" v:"required"`
	CreateAt *gtime.Time `json:"created_at"`
	CreateBy int64       `json:"created_by"`
}

type UpdateReq struct {
	g.Meta `path:"/user-app/{id}" method:"patch" tags:"user_app" summary:"Update user app"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     int64  `json:"id" v:"required"`
	UserId int64  `json:"userId" v:"required"`
	AppId  int64  `json:"appId" v:"required"`
}

type UpdateRes struct {
	Id       int64       `json:"id"`
	UserId   int64       `json:"userId" v:"required"`
	AppId    int64       `json:"appId" v:"required"`
	UpdateAt *gtime.Time `json:"updated_at"`
	UpdateBy int64       `json:"updated_by"`
}

type DeleteReq struct {
	g.Meta `path:"/user-app/{id}" method:"delete" tags:"user_app" summary:"Delete user app"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     int64  `json:"id" v:"required"`
}

type DeleteRes struct {
	Status string `json:"status"`
}
