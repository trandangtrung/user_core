package main

import (
	"demo/config"
	"demo/internal/db"
	"demo/internal/initialize"
	_ "demo/internal/packed"
	"fmt"
	"os"

	"github.com/gogf/gf/v2/os/gctx"

	"demo/internal/cmd"
)

func main() {

	if len(os.Args) > 1 {
		env := os.Args[1]

		initialize.InitLogger(env)
	} else {
		fmt.Println("No environment argument provided")
	}

	initialize.InitLoadConfig()
	initialize.InitToken()

	// init config
	cfg := config.GetConfig()
	fmt.Printf("Config: %v\n", cfg)

	// init db
	db := db.GetDatabaseConnection()
	if db == nil {
		fmt.Println("Failed to connect to the database")
		return
	}
	defer db.Close()

	cmd.Main.Run(gctx.GetInitCtx())
}
