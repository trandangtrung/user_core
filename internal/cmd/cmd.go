package cmd

import (
	"context"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"demo/internal/controller/auth"
	"demo/internal/controller/hello"
	"demo/internal/controller/platform"
	"demo/internal/controller/role"
	"demo/internal/controller/token"
	"demo/internal/controller/user"
	"demo/internal/controller/userPlatform"
	"demo/internal/controller/userRole"
	"demo/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			middle := middleware.NewMiddleware()
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, middle.AuthMiddleware("", "", true))
				group.Bind(
					auth.NewV1().Login,
					auth.NewV1().Signup,
				)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(middle.AuthMiddleware("admin", "network", false))
					group.Bind(hello.NewV1(),
						user.NewV1(),
						auth.NewV1().RefreshToken,
						role.NewV1(),
						userPlatform.NewV1(),
						platform.NewV1(),
						userRole.NewV1(),
						token.NewV1())
				})

			})
			s.Run()
			return nil
		},
	}
)
