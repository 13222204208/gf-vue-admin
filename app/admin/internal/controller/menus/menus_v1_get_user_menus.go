package menus

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"gf-vue-admin/app/admin/api/menus/v1"
)

func (c *ControllerV1) GetUserMenus(ctx context.Context, req *v1.GetUserMenusReq) (res *v1.GetUserMenusRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
