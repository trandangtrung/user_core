package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"user" summary:"Get user"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     uint   `json:"id" v:"required"`
}
type GetRes struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

type CreateReq struct {
	g.Meta   `path:"/user" method:"post" tags:"user" summary:"Create user"`
	Scope    string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	UserName string `json:"userName" v:"required"`
	Email    string `json:"email" v:"required|email"`
	Password string `json:"password" v:"required"`
	Mobile   string `json:"mobile" v:"required"`
	Role     uint   `json:"role" v:"required"`
	Apps     []uint `json:"apps" v:"required"`
}
type CreateRes struct {
	Status string `json:"status"`
}

type UpdateReq struct {
	g.Meta `path:"/user/{id}" method:"put" tags:"user" summary:"Update user"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     uint   `json:"id" v:"required"`
	Role   uint   `json:"role" v:"required"`
	Apps   []uint `json:"apps" v:"required"`
	Email  string `json:"email" v:"required|email"`

	// tùy theo bài toán có được phép đổi mật khẩu không
	Password string `json:"password" v:"required"`
}
type UpdateRes struct {
	Id    uint   `json:"id" v:"required"`
	Email string `json:"email" v:"required|email"`
	Role  uint   `json:"role" v:"required"`
	Apps  []uint `json:"apps" v:"required"`
}

type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"user" summary:"Delete user"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     uint   `json:"id" v:"required"`
}
type DeleteRes struct {
	Status string `json:"status" v:"required"`
}
