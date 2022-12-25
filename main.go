package main

import (
	_ "FullTextSearch/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"FullTextSearch/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
