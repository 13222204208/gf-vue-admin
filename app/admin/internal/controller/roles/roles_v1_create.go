package roles

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "gf-vue-admin/app/admin/api/roles/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/do"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 检查角色名称是否已存在
	count, err := dao.AdminRoles.Ctx(ctx).Where(dao.AdminRoles.Columns().Name, req.Name).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("角色名称已存在")
	}

	// 设置默认状态
	status := req.Status
	if status == "" {
		status = "active"
	}

	// 创建角色
	result, err := dao.AdminRoles.Ctx(ctx).Data(do.AdminRoles{
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Description: req.Description,
		Status:      status,
	}).Insert()
	if err != nil {
		return nil, err
	}

	// 获取插入的ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{
		Id: gconv.Uint64(id),
	}, nil
}