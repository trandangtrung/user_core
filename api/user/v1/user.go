package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"user" summary:"Get user"`
	Id     int64 `json:"id" v:"required"`
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
	Role     int    `json:"role" v:"required|min:1"`
	Apps     []int  `json:"apps" v:"required"`
}

type CreateRes struct {
	Status string `json:"status"`
}

type UpdateReq struct {
	g.Meta `path:"/user/{id}" method:"put" tags:"user" summary:"Update user"`
	Id     int64  `json:"id" v:"required|min:1"`
	Role   int    `json:"role" v:"required|min:1"`
	Apps   []int  `json:"apps" v:"required"`
	Email  string `json:"email" v:"required|email"`

	// tùy theo bài toán có được phép đổi mật khẩu không
	Password string `json:"password" v:"required|password"`
}
type UpdateRes struct {
	Id    int64  `json:"id" v:"required|min:1"`
	Email string `json:"email" v:"required|email"`
	Role  int    `json:"role" v:"required|min:1"`
	Apps  []int  `json:"apps" v:"required"`
}

type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"user" summary:"Delete user"`
	Id     int64 `json:"id" v:"required|min:1"`
}

type DeleteRes struct {
	Status string `json:"status" v:"required"`
}
