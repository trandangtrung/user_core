package cmd

import (
	"context"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"

	"github.com/quannv/strongbody-api/internal/config"
	"github.com/quannv/strongbody-api/internal/router"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetPort(config.GetConfig().ServerCfg.Port)

			// Additional Swagger UI settings
			s.GetOpenApi().Info.Title = "StrongBody API"
			s.GetOpenApi().Info.Description = "API documentation for StrongBody"
			s.GetOpenApi().Info.Version = "1.0.0"
			s.GetOpenApi().Servers = &goai.Servers{
				{
					URL:         "http://localhost:" + gconv.String(config.GetConfig().ServerCfg.Port),
					Description: "Local server",
				},
			}
			s.GetOpenApi().Security = &goai.SecurityRequirements{
				{
					"bearerAuth": []string{},
				},
			}
			s.GetOpenApi().Components.SecuritySchemes = goai.SecuritySchemes{
				"bearerAuth": goai.SecuritySchemeRef{
					Value: &goai.SecurityScheme{
						Type:        "http",
						Scheme:      "bearer",
						Description: "Bearer token authentication",
						Name:        "Authorization",
						In:          "header",
					},
				},
			}

			// Ready Swagger UI template
			htmlContent := gfile.GetBytes(gfile.Join(gfile.MainPkgPath(), "resource", "template", "swagger.html"))
			if htmlContent == nil {
				return gerror.New("failed to load Swagger UI template")
			}
			s.SetSwaggerUITemplate(string(htmlContent))
			s.SetOpenApiPath("/api.json")
			s.SetSwaggerPath("/swagger")

			r := s.Group("/api/v1")
			router.Router(r)

			s.Run()
			return nil
		},
	}
)
