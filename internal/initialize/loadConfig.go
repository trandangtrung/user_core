package initialize

import (
	"context"
	"demo/global"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func InitLoadConfig() {
	err := godotenv.Load()
	if err != nil {
		global.Logger.Error(context.TODO(), "Error loading .env file")
	}

	global.VariableEnv.SecretKey = os.Getenv("SECRET_KEY")

	log.Println(os.Getenv("TIME_ACCESS"))
	
	global.VariableEnv.TimeAccess, err = time.ParseDuration(os.Getenv("TIME_ACCESS"))

	if err != nil {
		global.Logger.Error(context.TODO(), err)
	}
	global.VariableEnv.TimeRefresh, err = time.ParseDuration(os.Getenv("TIME_REFRESH"))

	if err != nil {
		global.Logger.Error(context.TODO(), err)
	}

}
