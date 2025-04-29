package router

import (
	"strongbody-api/internal/controller/app"
	"strongbody-api/internal/controller/auth"
	"strongbody-api/internal/controller/role"
	"strongbody-api/internal/controller/user"
	"strongbody-api/internal/middleware"
	"strongbody-api/internal/repository"
	adminRouter "strongbody-api/internal/router/admin"
	buyerRouter "strongbody-api/internal/router/buyer"
	sellerRouter "strongbody-api/internal/router/seller"
	"strongbody-api/internal/service"
	"strongbody-api/internal/storage/postgres"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Router(r *ghttp.RouterGroup) {

	// init database
	db := postgres.GetDatabaseConnection().Connection

	// init repository
	appRepo := repository.NewAppRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	// userAppRepo := repository.NewUserAppRepository(db)
	// userRoleRepo := repository.NewUserRoleRepository(db)
	userRepo := repository.NewUserRepository(db)

	// init logic
	authService := service.NewAuthService(userRepo, roleRepo, tokenRepo)
	appService := service.NewAppService(appRepo)
	roleService := service.NewRoleService(roleRepo)
	userService := service.NewUserService(userRepo)

	// init controller
	authController := auth.NewV1(authService)
	appController := app.NewV1(appService)
	roleController := role.NewV1(roleService)
	userController := user.NewV1(userService)

	// init middleware
	middleware := middleware.NewMiddleware()

	r.Middleware(ghttp.MiddlewareHandlerResponse)

	// register router
	adminRouter.Register(r)
	buyerRouter.Register(r, middleware, authController, userController, roleController,
		appController)
	sellerRouter.Register(r, middleware)
}
