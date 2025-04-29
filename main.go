package main

import (
	"fmt"
	"os"
	"strongbody-api/internal/initialize"
	_ "strongbody-api/internal/packed"
	"strongbody-api/internal/storage/postgres"

	"github.com/gogf/gf/v2/os/gctx"

	"strongbody-api/internal/cmd"
)

func main() {
	if len(os.Args) > 1 {
		env := os.Args[1]

		initialize.InitLogger(env)
	} else {
		fmt.Println("Please specify the environment (e.g., dev, test, prod) as the first argument.")
	}
	initialize.InitToken()

	cmd.Main.Run(gctx.GetInitCtx())
	defer postgres.GetDatabaseConnection().Close()
}
