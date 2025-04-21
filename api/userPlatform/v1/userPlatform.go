package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type userPlatform struct {
}

type GetReq struct {
	g.Meta `path:"/user-platform/{id}" method:"get" tags:"userPlatform" summary:"Get userPlatform"`
	Id     int64 `json:"id" v:"required"`
}

type GetRes struct {
	Id         int64       `json:"id"`
	UserId     int64       `json:"userId" v:"required"`
	PlatformId int64       `json:"platformId" v:"required"`
	CreateAt   *gtime.Time `json:"created_at"`
	CreateBy   int64       `json:"created_by"`
	UpdateAt   *gtime.Time `json:"updated_at"`
	UpdateBy   int64       `json:"updated_by"`
}

type CreateReq struct {
	g.Meta     `path:"/user-platform" method:"post" tags:"userPlatform" summary:"Create userPlatform"`
	UserId     int64 `json:"userId" v:"required"`
	PlatformId int64 `json:"platformId" v:"required"`
}

type CreateRes struct {
	Id         int64       `json:"id" v:"required"`
	UserId     int64       `json:"userId" v:"required"`
	PlatformId int64       `json:"platformId" v:"required"`
	CreateAt   *gtime.Time `json:"created_at"`
	CreateBy   int64       `json:"created_by"`
}

type UpdateReq struct {
	g.Meta     `path:"/user-platform/{id}" method:"patch" tags:"userPlatform" summary:"Update userPlatform"`
	Id         int64 `json:"id" v:"required"`
	UserId     int64 `json:"userId" v:"required"`
	PlatformId int64 `json:"platformId" v:"required"`
}

type UpdateRes struct {
	Id         int64       `json:"id"`
	UserId     int64       `json:"userId" v:"required"`
	PlatformId int64       `json:"platformId" v:"required"`
	UpdateAt   *gtime.Time `json:"updated_at"`
	UpdateBy   int64       `json:"updated_by"`
}

type DeleteReq struct {
	g.Meta `path:"/user-platform/{id}" method:"delete" tags:"userPlatform" summary:"Delete userPlatform"`
	Id     int64 `json:"id" v:"required"`
}

type DeleteRes struct {
	Status string `json:"status"`
}
