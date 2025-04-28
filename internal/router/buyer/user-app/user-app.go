package userAppBuyerRouter

import (
	"demo/api/userApp"
	"demo/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.RouterGroup, middleware middleware.Middleware, userAppC userApp.IUserAppV1) {
	r.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.AuthMiddleware("admin", "network", false))
		group.Bind(
			userAppC,
		)
	})
}
