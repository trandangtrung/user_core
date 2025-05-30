package userBuyerRouter

import (
	"github.com/quannv/strongbody-api/api/user"
	"github.com/quannv/strongbody-api/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.RouterGroup, middleware middleware.Middleware, userC user.IUserV1) {
	r.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.AuthMiddleware("admin", "network", false))
		group.Bind(
			userC,
		)
	})
}
