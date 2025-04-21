package initialize

import (
	"context"
	"demo/global"
	"demo/utility/token"
)

func InitToken() {
	secretKey := global.VariableEnv.SecretKey
	maker, err := token.NewJWTMaker(secretKey)
	if err != nil {
		global.Logger.Error(context.TODO(), "init token error")
	}

	global.Token = maker
}
