package appBuyerRouter

import (
	"strongbody-api/api/app"
	"strongbody-api/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.RouterGroup, middleware middleware.Middleware, appC app.IAppV1) {
	r.Bind(appC.Create)
	r.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.AuthMiddleware("admin", "network", false))
		group.Bind(
			appC.Update,
			appC.Delete,
			appC.Get,
		)
	})
}
