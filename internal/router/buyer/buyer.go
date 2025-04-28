package buyerRouter

import (
	"demo/api/app"
	"demo/api/auth"
	"demo/api/role"
	"demo/api/user"
	"demo/internal/middleware"
	appBuyerRouter "demo/internal/router/buyer/app"
	authBuyerRouter "demo/internal/router/buyer/auth"
	roleBuyerRouter "demo/internal/router/buyer/role"
	userBuyerRouter "demo/internal/router/buyer/user"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.RouterGroup, middleware middleware.Middleware, authC auth.IAuthV1, userC user.IUserV1, roleC role.IRoleV1, appC app.IAppV1) {
	userBuyerRouter.Register(r, middleware, userC)
	authBuyerRouter.Register(r, middleware, authC)
	appBuyerRouter.Register(r, middleware, appC)
	roleBuyerRouter.Register(r, middleware, roleC)
}
