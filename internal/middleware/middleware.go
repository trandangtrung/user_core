package middleware

import "github.com/gogf/gf/v2/net/ghttp"

type Middleware interface {
	AuthMiddleware(permission string, scope string, isPublic bool) func(r *ghttp.Request)
}

type middlewareStr struct {
}

func NewMiddleware() Middleware {
	return &middlewareStr{}
}
