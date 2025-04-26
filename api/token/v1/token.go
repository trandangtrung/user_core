package v1

import "github.com/gogf/gf/v2/frame/g"

type GetReq struct {
	g.Meta `path:"/session/{id}" method:"get" tags:"session" summary:"Create session"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     int64  `json:"id" v:"required"`
}

type GetRes struct {
	Id           int64  `json:"id"`
	User_id      int64  `json:"user_id"`
	RefreshToken string `json:"refreshToken"`
	Scope        string `json:"scope"`
}

type CreateReq struct {
	g.Meta       `path:"/session" method:"post" tags:"session" summary:"Create session"`
	User_id      int64  `json:"user_id" v:"required"`
	RefreshToken string `json:"refreshToken" v:"required"`
	Scope        string `json:"Scope" v:"required"`
}

type CreateRes struct {
	Status string `json:"status"`
}

type UpdateReq struct {
	g.Meta       `path:"/session/{id}" method:"patch" tags:"session" summary:"Update session"`
	Id           int64  `json:"id" v:"required"`
	User_id      int64  `json:"user_id" v:"required"`
	RefreshToken string `json:"refreshToken" v:"required"`
	Scope        string `json:"Scope" v:"required"`
}

type UpdateRes struct {
	Id           int64  `json:"id"`
	User_id      int64  `json:"user_id"`
	RefreshToken string `json:"refreshToken"`
	Scope        string `json:"scope"`
}

type DeleteReq struct {
	g.Meta `path:"/session/{id}" method:"delete" tags:"session" summary:"Delete session"`
	Scope  string `in:"header" name:"Scope" default:"network" summary:"Scope"`
	Id     int64  `json:"id" v:"required"`
}

type DeleteRes struct {
	Status string `json:"status"`
}
