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

	// This is a placeholder for the actual seeding logic.
	// You can replace this with your actual seeding function.
	err := postgres.GetDatabaseConnection().Seed()
	if err != nil {
		fmt.Println("Error seeding database:", err)
		return
	}

	cmd.Main.Run(gctx.GetInitCtx())
	defer postgres.GetDatabaseConnection().Close()
}
