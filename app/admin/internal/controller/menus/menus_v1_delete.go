package menus

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf-vue-admin/app/admin/api/menus/v1"
	"gf-vue-admin/app/admin/internal/dao"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	// 检查是否存在子菜单
	count, err := dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().ParentId, req.Id).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询子菜单失败")
	}
	if count > 0 {
		return nil, gerror.New("请先删除子菜单")
	}

	// 删除菜单
	_, err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, gerror.Wrap(err, "删除菜单失败")
	}
	return

}
