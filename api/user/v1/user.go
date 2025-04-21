package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta   `path:"/user" method:"post" tags:"User" summary:"Create user"`
	Email    string `v:"required"`
	Password string `v:"required"`
}

type CreateRes struct {
}
