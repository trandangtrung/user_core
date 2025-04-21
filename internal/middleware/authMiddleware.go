package middleware

import (
	"context"
	"demo/global"
	"demo/internal/consts"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func (m *middlewareStr) AuthMiddleware(permission string, scope string, isPublic bool) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		authorization := r.GetHeader(consts.AuthorizationHeader)
		authorizationScope := r.GetHeader(consts.AuthorizationScope)

		if consts.CONFIG_SCOPE[authorizationScope] == "" {
			global.Logger.Error(context.Background(), "unauthorized platform")
			r.Response.WriteStatus(http.StatusForbidden, "unauthorized platform")
			return
		}

		if authorization != "" {
			if len(authorization) == 0 {
				global.Logger.Error(context.Background(), "please provide authorization")
				r.Response.WriteStatus(http.StatusForbidden)
				return
			}

			fields := strings.Fields(authorization)

			if len(fields) < 2 {
				global.Logger.Error(context.Background(), "invalid format header")
				r.Response.WriteStatus(http.StatusForbidden, "invalid format header")
			}

			if fields[0] != consts.AuthorizationType {
				global.Logger.Error(context.Background(), "invalid type header")
				r.Response.WriteStatus(http.StatusForbidden, "invalid type header")
				return
			}

			payload, err := global.Token.VerifyToken(fields[1])

			if err != nil {
				global.Logger.Error(context.Background(), "Verify token invalid")
				r.Response.WriteStatus(http.StatusForbidden, "Verify token invalid")
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
				r.Response.WriteStatus(http.StatusForbidden, "unauthorized")
				return
			}
		} else if isPublic {
			ctx := context.WithValue(r.Context(), consts.AuthorizationScope, authorizationScope)

			r.SetCtx(ctx)
			r.Middleware.Next()
		} else {
			global.Logger.Error(context.Background(), "unauthorized")
			r.Response.WriteStatus(http.StatusForbidden, "unauthorized")
			return
		}

		r.Middleware.Next()
	}

}
