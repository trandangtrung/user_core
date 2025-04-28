package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"user" summary:"Get user"`
	Id     int64  `json:"id" v:"required"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
}
type GetRes struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type CreateReq struct {
	g.Meta   `path:"/user" method:"post" tags:"user" summary:"Create user"`
	UserName string `json:"userName" v:"required"`
	Email    string `json:"email" v:"required|email"`
	Password string `json:"password" v:"required|password"`
	Role     int    `json:"role" v:"required"`
	Apps     []int  `json:"apps" v:"required"`
}

type CreateRes struct {
	Status string
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
