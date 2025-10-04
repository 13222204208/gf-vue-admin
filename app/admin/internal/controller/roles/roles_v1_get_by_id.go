package roles

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf-vue-admin/app/admin/api/roles/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/entity"
)

func (c *ControllerV1) GetById(ctx context.Context, req *v1.GetByIdReq) (res *v1.GetByIdRes, err error) {
	var role entity.AdminRoles
	err = dao.AdminRoles.Ctx(ctx).Where(dao.AdminRoles.Columns().Id, req.Id).Scan(&role)
	if err != nil {
		return nil, err
	}

	// 检查角色是否存在
	if role.Id == 0 {
		return nil, gerror.New("角色不存在")
	}

	return &v1.GetByIdRes{
		Id:          role.Id,
		Name:        role.Name,
		DisplayName: role.DisplayName,
		Description: role.Description,
		Status:      role.Status,
		UserCount:   0, // TODO: 计算使用该角色的用户数量
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}