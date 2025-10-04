package users

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf-vue-admin/app/admin/api/users/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/do"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	// 检查用户是否存在
	count, err := dao.AdminUsers.Ctx(ctx).Where("id", req.Id).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("用户不存在")
	}

	// 如果更新用户名，检查是否重复
	if req.Username != "" {
		count, err = dao.AdminUsers.Ctx(ctx).Where(dao.AdminUsers.Columns().Id, req.Username).WhereNot(dao.AdminUsers.Columns().Id, req.Id).Count()
		if err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, gerror.New("用户名已存在")
		}
	}

	// 构建更新数据
	updateData := do.AdminUsers{}
	if req.Username != "" {
		updateData.Username = req.Username
	}
	if req.Nickname != "" {
		updateData.Nickname = req.Nickname
	}
	if req.Email != "" {
		updateData.Email = req.Email
	}
	if req.Phone != "" {
		updateData.Phone = req.Phone
	}
	if req.Avatar != "" {
		updateData.Avatar = req.Avatar
	}
	if req.Status != "" {
		updateData.Status = req.Status
	}
	if req.RoleId > 0 {
		updateData.RoleId = req.RoleId
	}

	// 更新数据
	_, err = dao.AdminUsers.Ctx(ctx).Where("id", req.Id).Data(updateData).Update()
	if err != nil {
		return nil, err
	}

	return &v1.UpdateRes{}, nil
}
