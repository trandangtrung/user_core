package initialize

import (
	"context"

	"github.com/quannv/strongbody-api/global"
	"github.com/quannv/strongbody-api/internal/config"
	"github.com/quannv/strongbody-api/utility/token"
)

func InitToken() {
	secretKey := config.GetConfig().JwtCfg.SecretKey
	maker, err := token.NewJWTMaker(secretKey)
	if err != nil {
		global.Logger.Error(context.TODO(), "init token error")
	}

	global.Token = maker
}
