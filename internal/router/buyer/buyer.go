package buyerRouter

import (
	"demo/api/app"
	"demo/api/auth"
	"demo/api/role"
	"demo/api/token"
	"demo/api/user"
	"demo/api/userApp"
	"demo/api/userRole"
	"demo/internal/middleware"
	appBuyerRouter "demo/internal/router/buyer/app"
	authBuyerRouter "demo/internal/router/buyer/auth"
	roleBuyerRouter "demo/internal/router/buyer/role"
	tokenBuyerRouter "demo/internal/router/buyer/token"
	userBuyerRouter "demo/internal/router/buyer/user"
	userAppBuyerRouter "demo/internal/router/buyer/user-app"
	userRoleBuyerRouter "demo/internal/router/buyer/user-role"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.RouterGroup, middleware middleware.Middleware, authC auth.IAuthV1, userC user.IUserV1, roleC role.IRoleV1, userRoleC userRole.IUserRoleV1, tokenC token.ITokenV1, appC app.IAppV1, userAppC userApp.IUserAppV1) {
	userBuyerRouter.Register(r, middleware, userC)
	authBuyerRouter.Register(r, middleware, authC)
	appBuyerRouter.Register(r, middleware, appC)
	roleBuyerRouter.Register(r, middleware, roleC)
	tokenBuyerRouter.Register(r, middleware, tokenC)
	userAppBuyerRouter.Register(r, middleware, userAppC)
	userRoleBuyerRouter.Register(r, middleware, userRoleC)
}
