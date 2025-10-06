package roles

import (
	"context"

	v1 "gf-vue-admin/app/admin/api/roles/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/entity"
)

func (c *ControllerV1) GetActiveList(ctx context.Context, req *v1.GetActiveListReq) (res *v1.GetActiveListRes, err error) {
	// 查询所有启用状态的角色
	var roles []entity.AdminRoles
	err = dao.AdminRoles.Ctx(ctx).
		Where(dao.AdminRoles.Columns().Status, "active").
		OrderAsc(dao.AdminRoles.Columns().DisplayName).
		Scan(&roles)
	if err != nil {
		return nil, err
	}

	// 构建响应数据，只返回ID和显示名称
	var list []struct {
		Id          uint64 `json:"id" dc:"主键ID"`
		DisplayName string `json:"displayName" dc:"角色显示名称"`
	}
	
	for _, role := range roles {
		list = append(list, struct {
			Id          uint64 `json:"id" dc:"主键ID"`
			DisplayName string `json:"displayName" dc:"角色显示名称"`
		}{
			Id:          role.Id,
			DisplayName: role.DisplayName,
		})
	}

	return &v1.GetActiveListRes{
		List: list,
	}, nil
}
