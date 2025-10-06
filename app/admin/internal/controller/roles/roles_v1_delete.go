package roles

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf-vue-admin/app/admin/api/roles/v1"
	"gf-vue-admin/app/admin/internal/dao"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	// 检查角色是否存在
	count, err := dao.AdminRoles.Ctx(ctx).Where(dao.AdminRoles.Columns().Id, req.Id).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("角色不存在")
	}

	// TODO: 检查角色是否被用户使用
	// 这里可以添加检查逻辑，查询是否有用户关联此角色
	userCount, err := dao.AdminUsers.Ctx(ctx).Where("role_id", req.Id).Count()
	if err != nil {
		return nil, err
	}
	if userCount > 0 {
		return nil, gerror.New("角色正在被用户使用，无法删除")
	}

	// 执行删除
	_, err = dao.AdminRoles.Ctx(ctx).Where(dao.AdminRoles.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, err
	}

	return &v1.DeleteRes{}, nil
}
