package authBuyerRouter

import (
	"github.com/quannv/strongbody-api/api/auth"
	"github.com/quannv/strongbody-api/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.RouterGroup, middleware middleware.Middleware, authC auth.IAuthV1) {
	r.Bind(authC.Signup)

	r.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.AuthMiddleware("", "", true))
		group.Bind(
			authC.Login,
			authC.LoginByToken,
			authC.RefreshToken,
		)
	})
}
