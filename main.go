package main

import (
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

	cmd.Main.Run(gctx.GetInitCtx())
}
