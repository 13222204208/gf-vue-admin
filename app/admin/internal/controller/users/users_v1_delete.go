package users

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf-vue-admin/app/admin/api/users/v1"
	"gf-vue-admin/app/admin/internal/dao"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	// 检查用户是否存在
	count, err := dao.AdminUsers.Ctx(ctx).Where(dao.AdminUsers.Columns().Id, req.Id).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("用户不存在")
	}

	// 删除用户
	_, err = dao.AdminUsers.Ctx(ctx).Where(dao.AdminUsers.Columns().Id, req.Id).Delete()
	if err != nil {
		return nil, err
	}

	return &v1.DeleteRes{}, nil
}
