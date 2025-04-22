package router

import (
	"demo/internal/controller/auth"
	"demo/internal/controller/platform"
	"demo/internal/controller/role"
	"demo/internal/controller/token"
	"demo/internal/controller/user"
	"demo/internal/controller/userPlatform"
	"demo/internal/controller/userRole"
	"demo/internal/middleware"
	adminRouter "demo/internal/router/admin"
	buyerRouter "demo/internal/router/buyer"
	sellerRouter "demo/internal/router/seller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Router(r *ghttp.RouterGroup) {
	middleware := middleware.NewMiddleware()

	// init controller
	authC := auth.NewV1()
	userC := user.NewV1()
	roleC := role.NewV1()
	userRoleC := userRole.NewV1()
	tokenC := token.NewV1()
	platformC := platform.NewV1()
	userPlatformC := userPlatform.NewV1()

	// register router
	adminRouter.Register(r)
	buyerRouter.Register(r, middleware, authC, userC, roleC, userRoleC, tokenC, platformC, userPlatformC)
	sellerRouter.Register(r, middleware)
}
