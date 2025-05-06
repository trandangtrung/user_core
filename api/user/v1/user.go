package v1

import (
	userDto "github.com/quannv/strongbody-api/internal/dto/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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
	Status  string `json:"status"`
	ID      int64  `json:"id"`
	Message string `json:"message"`
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
	Id        uint        `json:"id" v:"required"`
	Email     string      `json:"email" v:"required|email"`
	Role      uint        `json:"role" v:"required"`
	Apps      []uint      `json:"apps" v:"required"`
	UpdatedAt *gtime.Time `json:"updated_at"`
	UpdatedBy uint        `json:"updated_by"`
}

type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"user" summary:"Delete user"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     uint   `json:"id" v:"required"`
}
type DeleteRes struct {
	Status string `json:"status" v:"required"`
}

type ListUsersReq struct {
	g.Meta `path:"/user" method:"get" tags:"user" summary:"List users with filters"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`

	Page     int    `in:"query" default:"1" description:"Page number" json:"page"`
	Limit    int    `in:"query" default:"10" description:"Number of records per page" json:"limit"`
	OrderBy  string `in:"query" default:"id" description:"Order by field" json:"order_by"`
	OrderDir string `in:"query" default:"ASC" description:"Order direction: ASC or DESC" json:"order_dir"`
	Keyword  string `in:"query" description:"Search keyword (applies to user_name, email)" json:"keyword"`

	UserName string `in:"query" description:"Filter by user name" json:"user_name"`
	Email    string `in:"query" description:"Filter by email" json:"email"`
	Mobile   string `in:"query" description:"Filter by mobile number" json:"mobile"`
	Country  string `in:"query" description:"Filter by country" json:"country"`
	City     string `in:"query" description:"Filter by city" json:"city"`

	FromBirthDate string `in:"query" description:"Start birth date (format: YYYY-MM-DD)" json:"from_birth_date"`
	ToBirthDate   string `in:"query" description:"End birth date (format: YYYY-MM-DD)" json:"to_birth_date"`
}
type ListUsersRes struct {
	Total       int               `json:"total" description:"Total users found"`
	TotalPage   int               `json:"total_page" description:"Total pages"`
	CurrentPage int               `json:"current_page" description:"Current page"`
	Limit       int               `json:"limit" description:"Number of records per page"`
	Data        []userDto.UserRes `json:"data" description:"List of users"`
}
