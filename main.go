package main

import (
	_ "goframe/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
