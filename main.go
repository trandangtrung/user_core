package main

import (
	"demo/internal/initialize"
	_ "demo/internal/packed"
	"demo/internal/storage/postgres"
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
		fmt.Println("Please specify the environment (e.g., dev, test, prod) as the first argument.")
	}
	initialize.InitToken()
	postgres.GetDatabaseConnection().Seed()

	cmd.Main.Run(gctx.GetInitCtx())

	defer postgres.GetDatabaseConnection().Close()
}
