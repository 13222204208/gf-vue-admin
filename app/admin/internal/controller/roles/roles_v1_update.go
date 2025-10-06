package roles

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf-vue-admin/app/admin/api/roles/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/do"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	// 检查角色是否存在
	count, err := dao.AdminRoles.Ctx(ctx).Where(dao.AdminRoles.Columns().Id, req.Id).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("角色不存在")
	}

	// 检查角色名称是否已被其他角色使用
	if req.Name != "" {
		existCount, err := dao.AdminRoles.Ctx(ctx).
			Where(dao.AdminRoles.Columns().Name, req.Name).
			WhereNot(dao.AdminRoles.Columns().Id, req.Id).
			Count()
		if err != nil {
			return nil, err
		}
		if existCount > 0 {
			return nil, gerror.New("角色名称已被其他角色使用")
		}
	}

	// 构建更新数据
	updateData := do.AdminRoles{}
	if req.Name != "" {
		updateData.Name = req.Name
	}
	if req.DisplayName != "" {
		updateData.DisplayName = req.DisplayName
	}
	if req.Description != "" {
		updateData.Description = req.Description
	}
	if req.Status != "" {
		updateData.Status = req.Status
	}

	// 执行更新
	_, err = dao.AdminRoles.Ctx(ctx).Where(dao.AdminRoles.Columns().Id, req.Id).Data(updateData).Update()
	if err != nil {
		return nil, err
	}

	return &v1.UpdateRes{}, nil
}
