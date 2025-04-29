package middleware

import (
	"context"
	"strings"
	"strongbody-api/global"
	"strongbody-api/internal/consts"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (m *middlewareStr) AuthMiddleware(permission string, scope string, isPublic bool) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		authorization := r.GetHeader(consts.AuthorizationHeader)
		authorizationScope := r.GetHeader(consts.AuthorizationScope)

		if consts.CONFIG_SCOPE[authorizationScope] == "" {
			global.Logger.Error(context.Background(), "unauthorized platform")
			r.Response.WriteJsonExit(g.Map{
				"code":    403,
				"message": "unauthorized platform",
				"data":    nil,
			})
			return
		}

		if authorization != "" {
			if len(authorization) == 0 {
				global.Logger.Error(context.Background(), "please provide authorization")
				r.Response.WriteJsonExit(g.Map{
					"code":    403,
					"message": "please provide authorization",
					"data":    nil,
				})
				return
			}

			fields := strings.Fields(authorization)

			if len(fields) < 2 {
				global.Logger.Error(context.Background(), "invalid format header")
				r.Response.WriteJsonExit(g.Map{
					"code":    403,
					"message": "invalid format header",
					"data":    nil,
				})
				return
			}

			if fields[0] != consts.AuthorizationType {
				global.Logger.Error(context.Background(), "invalid type header")
				r.Response.WriteJsonExit(g.Map{
					"code":    403,
					"message": "invalid type header",
					"data":    nil,
				})

				return
			}

			payload, err := global.Token.VerifyToken(fields[1])

			if err != nil {
				global.Logger.Error(context.Background(), "verify token invalid")
				r.Response.WriteJsonExit(g.Map{
					"code":    403,
					"message": "verify token invalid",
					"data":    nil,
				})
				return
			}

			// payload check role

			if payload.Permissions == permission || payload.Permissions == consts.CONFIG_PERMISSIONS["ADMIN"] {
				ctx := context.WithValue(r.Context(), consts.AuthorizationKey, payload)

				ctx = context.WithValue(ctx, consts.AuthorizationScope, authorizationScope)

				r.SetCtx(ctx)
				r.Middleware.Next()
				return
			} else {
				global.Logger.Error(context.Background(), "unauthorized")
				r.Response.WriteJsonExit(g.Map{
					"code":    403,
					"message": "unauthorized",
					"data":    nil,
				})
				return
			}
		} else if isPublic {
			ctx := context.WithValue(r.Context(), consts.AuthorizationScope, authorizationScope)

			r.SetCtx(ctx)
			r.Middleware.Next()
		} else {
			global.Logger.Error(context.Background(), "unauthorized")
			r.Response.WriteJsonExit(g.Map{
				"code":    403,
				"message": "unauthorized",
				"data":    nil,
			})
			return
		}

		r.Middleware.Next()
	}

}
