package main

import (
	_ "gf-vue-admin/app/admin/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-vue-admin/app/admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
