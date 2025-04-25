package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Role struct {
}

type GetReq struct {
	g.Meta `path:"/role/{id}" method:"get" tags:"role" summary:"Get role"`
	Id     int64 `json:"id" v:"required"`
}

type GetRes struct {
	Id             int64       `json:"id"`
	UserPlatformId int64       `json:"userPlatformId" v:"required"`
	Name           string      `json:"name" v:"required"`
	Description    string      `json:"description" v:"required"`
	CreateAt       *gtime.Time `json:"created_at"`
	CreateBy       int64       `json:"created_by"`
	UpdateAt       *gtime.Time `json:"updated_at"`
	UpdateBy       int64       `json:"updated_by"`
}

type CreateReq struct {
	g.Meta         `path:"/role" method:"post" tags:"role" summary:"Create role"`
	UserPlatformId int64  `json:"userPlatformId" v:"required"`
	Name           string `json:"name" v:"required"`
	Description    string `json:"description" v:"required"`
	CreateBy       int64  `json:"create_by"`
}

type CreateRes struct {
	Id             int64       `json:"id" v:"required"`
	UserPlatformId int64       `json:"userPlatformId" v:"required"`
	Name           string      `json:"name" v:"required"`
	Description    string      `json:"description" v:"required"`
	CreateAt       *gtime.Time `json:"created_at"`
	CreateBy       int64       `json:"created_by"`
}

type UpdateReq struct {
	g.Meta         `path:"/role/{id}" method:"patch" tags:"role" summary:"Update role"`
	Id             int64  `json:"id" v:"required"`
	UserPlatformId int64  `json:"userPlatformId" v:"required"`
	Name           string `json:"name" v:"required"`
	Description    string `json:"description" v:"required"`
}

type UpdateRes struct {
	Id             int64       `json:"id"`
	UserPlatformId int64       `json:"userPlatformId" v:"required"`
	Name           string      `json:"name" v:"required"`
	Description    string      `json:"description" v:"required"`
	UpdateAt       *gtime.Time `json:"updated_at"`
	UpdateBy       int64       `json:"updated_by"`
}

type DeleteReq struct {
	g.Meta `path:"/role/{id}" method:"delete" tags:"role" summary:"Delete role"`
	Id     int64 `json:"id" v:"required"`
}

type DeleteRes struct {
	Status string `json:"status"`
}
