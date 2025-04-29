package initialize

import (
	"context"
	"strongbody-api/global"
	"strongbody-api/internal/config"
	"strongbody-api/utility/token"
)

func InitToken() {
	secretKey := config.GetConfig().JwtCfg.SecretKey
	maker, err := token.NewJWTMaker(secretKey)
	if err != nil {
		global.Logger.Error(context.TODO(), "init token error")
	}

	global.Token = maker
}
