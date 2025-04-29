package buyerRouter

import (
	"strongbody-api/api/app"
	"strongbody-api/api/auth"
	"strongbody-api/api/role"
	"strongbody-api/api/user"
	"strongbody-api/internal/middleware"
	appBuyerRouter "strongbody-api/internal/router/buyer/app"
	authBuyerRouter "strongbody-api/internal/router/buyer/auth"
	roleBuyerRouter "strongbody-api/internal/router/buyer/role"
	userBuyerRouter "strongbody-api/internal/router/buyer/user"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.RouterGroup, middleware middleware.Middleware, authC auth.IAuthV1, userC user.IUserV1, roleC role.IRoleV1, appC app.IAppV1) {
	userBuyerRouter.Register(r, middleware, userC)
	authBuyerRouter.Register(r, middleware, authC)
	appBuyerRouter.Register(r, middleware, appC)
	roleBuyerRouter.Register(r, middleware, roleC)
}
