package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type User struct {
	Email string `json:"email"`
	Role  string
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"Auth" summary:"Login user"`
	Email    string `json:"email" v:"required|email"`
	Password string `json:"password" v:"password"`
}

type LoginRes struct {
	User  User  `json:"user"`
	Token Token `json:"token"`
}

type SignupReq struct {
	g.Meta   `path:"/signup" method:"post" tags:"Auth" summary:"Sign up user"`
	Email    string `json:"email" v:"required|email"`
	Password string `json:"password" v:"password"`
}

type SignupRes struct {
	Status string
}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh-token" method:"get" tags:"Auth" summary:"refresh token"`
}

type RefreshTokenRes struct {
	AccessToken string
}
