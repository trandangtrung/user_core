package userRoleBuyerRouter

import (
	"demo/api/userRole"
	"demo/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.RouterGroup, middleware middleware.Middleware, userRoleC userRole.IUserRoleV1) {
	r.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.AuthMiddleware("admin", "network", false))
		group.Bind(
			userRoleC,
		)
	})
}
