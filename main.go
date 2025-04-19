package main

import (
	_ "strongbody-api/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"strongbody-api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
