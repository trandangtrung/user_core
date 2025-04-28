package router

import (
	"demo/internal/controller/app"
	"demo/internal/controller/auth"
	"demo/internal/controller/role"
	"demo/internal/controller/token"
	"demo/internal/controller/user"
	"demo/internal/controller/userApp"
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
	appRepo := repository.NewAppRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	userAppRepo := repository.NewUserAppRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)
	userRepo := repository.NewUserRepository(db)

	// init logic
	authService := service.NewAuthService(userRepo, roleRepo, tokenRepo)
	appService := service.NewAppService(appRepo)
	roleService := service.NewRoleService(roleRepo)
	tokenService := service.NewTokenService(tokenRepo)
	userAppService := service.NewUserAppService(userAppRepo)
	userRoleService := service.NewUserRoleService(userRoleRepo)
	userService := service.NewUserService(userRepo)

	// init controller
	authController := auth.NewV1(authService)
	appController := app.NewV1(appService)
	roleController := role.NewV1(roleService)
	tokenController := token.NewV1(tokenService)
	userController := user.NewV1(userService)
	userAppController := userApp.NewV1(userAppService)
	userRoleController := userRole.NewV1(userRoleService)

	// init middleware
	middleware := middleware.NewMiddleware()

	r.Middleware(ghttp.MiddlewareHandlerResponse)

	// register router
	adminRouter.Register(r)
	buyerRouter.Register(r, middleware, authController, userController, roleController,
		userRoleController, tokenController, appController, userAppController)
	sellerRouter.Register(r, middleware)
}
