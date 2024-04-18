package main

import (
	_ "jonog/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"jonog/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
