package buyerRouter

import (
	"github.com/quannv/strongbody-api/api/app"
	"github.com/quannv/strongbody-api/api/auth"
	"github.com/quannv/strongbody-api/api/role"
	"github.com/quannv/strongbody-api/api/user"
	"github.com/quannv/strongbody-api/internal/middleware"
	appBuyerRouter "github.com/quannv/strongbody-api/internal/router/buyer/app"
	authBuyerRouter "github.com/quannv/strongbody-api/internal/router/buyer/auth"
	roleBuyerRouter "github.com/quannv/strongbody-api/internal/router/buyer/role"
	userBuyerRouter "github.com/quannv/strongbody-api/internal/router/buyer/user"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.RouterGroup, middleware middleware.Middleware, authC auth.IAuthV1, userC user.IUserV1, roleC role.IRoleV1, appC app.IAppV1) {
	userBuyerRouter.Register(r, middleware, userC)
	authBuyerRouter.Register(r, middleware, authC)
	appBuyerRouter.Register(r, middleware, appC)
	roleBuyerRouter.Register(r, middleware, roleC)
}
