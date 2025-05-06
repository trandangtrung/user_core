package main

import (
	"fmt"
	"os"

	"github.com/quannv/strongbody-api/internal/initialize"
	_ "github.com/quannv/strongbody-api/internal/packed"
	"github.com/quannv/strongbody-api/internal/storage/postgres"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/quannv/strongbody-api/internal/cmd"
)

func main() {
	if len(os.Args) > 1 {
		env := os.Args[1]

		initialize.InitLogger(env)
	} else {
		fmt.Println("Please specify the environment (e.g., dev, test, prod) as the first argument.")
	}
	initialize.InitToken()
	initialize.InitGmailAndTemplate()

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
