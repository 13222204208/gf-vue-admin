package main

import (
	_ "gf-vue-admin/app/admin/internal/packed"

	"gf-vue-admin/app/admin/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
