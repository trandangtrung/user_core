package initialize

import (
	"context"
	"demo/global"
	"demo/internal/config"
	"demo/utility/token"
)

func InitToken() {
	secretKey := config.GetConfig().JwtCfg.SecretKey
	maker, err := token.NewJWTMaker(secretKey)
	if err != nil {
		global.Logger.Error(context.TODO(), "init token error")
	}

	global.Token = maker
}
