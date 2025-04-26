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
	"demo/internal/repository"
	adminRouter "demo/internal/router/admin"
	buyerRouter "demo/internal/router/buyer"
	sellerRouter "demo/internal/router/seller"
	"demo/internal/service"
	"demo/internal/storage/postgres"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Router(r *ghttp.RouterGroup) {

	// init database
	db := postgres.GetDatabaseConnection().Connection

	// init repository
	platformRepo := repository.NewPlatformRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	userPlatformRepo := repository.NewUserPlatformRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)
	userRepo := repository.NewUserRepository(db)

	// init logic
	authService := service.NewAuthService(userRepo, roleRepo, tokenRepo)
	platformService := service.NewPlatformService(platformRepo)
	roleService := service.NewRoleService(roleRepo)
	tokenService := service.NewTokenService(tokenRepo)
	userPlatformService := service.NewUserPlatformService(userPlatformRepo)
	userRoleService := service.NewUserRoleService(userRoleRepo)
	userService := service.NewUserService(userRepo)

	// init controller
	authController := auth.NewV1(authService)
	platformController := platform.NewV1(platformService)
	roleController := role.NewV1(roleService)
	tokenController := token.NewV1(tokenService)
	userController := user.NewV1(userService)
	userPlatformController := userPlatform.NewV1(userPlatformService)
	userRoleController := userRole.NewV1(userRoleService)

	// init middleware
	middleware := middleware.NewMiddleware()

	r.Middleware(ghttp.MiddlewareHandlerResponse)
	// register router
	adminRouter.Register(r)
	buyerRouter.Register(r, middleware, authController, userController, roleController,
		userRoleController, tokenController, platformController, userPlatformController)
	sellerRouter.Register(r, middleware)
}
