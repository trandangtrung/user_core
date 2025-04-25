package authBuyerRouter

import (
	"demo/api/auth"
	"demo/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.RouterGroup, middleware middleware.Middleware, authC auth.IAuthV1) {
	r.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.AuthMiddleware("", "", true), ghttp.MiddlewareHandlerResponse)
		group.Bind(
			authC,
		)
	})
}
