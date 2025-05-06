package router

import (
	"github.com/quannv/strongbody-api/global"
	"github.com/quannv/strongbody-api/internal/controller/app"
	"github.com/quannv/strongbody-api/internal/controller/auth"
	"github.com/quannv/strongbody-api/internal/controller/role"
	"github.com/quannv/strongbody-api/internal/controller/user"
	"github.com/quannv/strongbody-api/internal/middleware"
	"github.com/quannv/strongbody-api/internal/repository"
	adminRouter "github.com/quannv/strongbody-api/internal/router/admin"
	buyerRouter "github.com/quannv/strongbody-api/internal/router/buyer"
	sellerRouter "github.com/quannv/strongbody-api/internal/router/seller"
	"github.com/quannv/strongbody-api/internal/service"
	"github.com/quannv/strongbody-api/internal/storage/postgres"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Router(r *ghttp.RouterGroup) {

	// init database
	db := postgres.GetDatabaseConnection().Connection

	// init repository
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	appRepo := repository.NewAppRepository(db)

	// init logic
	mailService := service.NewGmailService(global.Gmail, global.Template)
	authService := service.NewAuthService(userRepo, roleRepo, mailService)
	userService := service.NewUserService(userRepo, roleRepo, appRepo)
	appService := service.NewAppService(appRepo)
	roleService := service.NewRoleService(roleRepo)

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
	buyerRouter.Register(r, middleware, authController, userController, roleController, appController)
	sellerRouter.Register(r, middleware)
}
