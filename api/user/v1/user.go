package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"user" summary:"Get user"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     int64  `json:"id" v:"required"`
}
type GetRes struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type UpdateReq struct {
	g.Meta   `path:"/user/{id}" method:"put" tags:"user" summary:"Update user"`
	Scope    string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id       int64  `v:"required"`
	Email    string `v:"required"`
	Password string `v:"required"`
}
type UpdateRes struct {
}

type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"user" summary:"Delete user"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     int64  `v:"required"`
}
type DeleteRes struct {
}
