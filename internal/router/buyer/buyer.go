package buyerRouter

import (
	"demo/api/auth"
	"demo/api/platform"
	"demo/api/role"
	"demo/api/token"
	"demo/api/user"
	"demo/api/userPlatform"
	"demo/api/userRole"
	"demo/internal/middleware"
	authBuyerRouter "demo/internal/router/buyer/auth"
	platformBuyerRouter "demo/internal/router/buyer/platform"
	roleBuyerRouter "demo/internal/router/buyer/role"
	tokenBuyerRouter "demo/internal/router/buyer/token"
	userBuyerRouter "demo/internal/router/buyer/user"
	userPlatformBuyerRouter "demo/internal/router/buyer/user-platform"
	userRoleBuyerRouter "demo/internal/router/buyer/user-role"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(r *ghttp.RouterGroup, middleware middleware.Middleware, authC auth.IAuthV1, userC user.IUserV1, roleC role.IRoleV1, userRoleC userRole.IUserRoleV1, tokenC token.ITokenV1, platformC platform.IPlatformV1, userPlatformC userPlatform.IUserPlatformV1) {
	userBuyerRouter.Register(r, middleware, userC)
	authBuyerRouter.Register(r, middleware, authC)
	platformBuyerRouter.Register(r, middleware, platformC)
	roleBuyerRouter.Register(r, middleware, roleC)
	tokenBuyerRouter.Register(r, middleware, tokenC)
	userPlatformBuyerRouter.Register(r, middleware, userPlatformC)
	userRoleBuyerRouter.Register(r, middleware, userRoleC)
}
