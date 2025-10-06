package roles

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "gf-vue-admin/app/admin/api/roles/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/common"
	"gf-vue-admin/app/admin/internal/model/entity"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	// 构建查询条件
	query := dao.AdminRoles.Ctx(ctx)

	// 按名称筛选
	if req.Name != "" {
		query = query.WhereLike(dao.AdminRoles.Columns().Name, "%"+req.Name+"%")
	}

	// 按显示名称筛选
	if req.DisplayName != "" {
		query = query.WhereLike(dao.AdminRoles.Columns().DisplayName, "%"+req.DisplayName+"%")
	}

	// 按状态筛选
	if req.Status != "" {
		query = query.Where(dao.AdminRoles.Columns().Status, req.Status)
	}

	// 获取总数
	total, err := query.Count()
	if err != nil {
		return nil, err
	}

	// 分页查询
	current, size := req.GetCurrentInfo()

	var roles []entity.AdminRoles
	err = query.Page(current, size).OrderDesc(dao.AdminRoles.Columns().Id).Scan(&roles)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	var list []v1.GetByIdRes
	for _, role := range roles {
		list = append(list, v1.GetByIdRes{
			Id:          role.Id,
			Name:        role.Name,
			DisplayName: role.DisplayName,
			Description: role.Description,
			Status:      role.Status,
			CreatedAt:   role.CreatedAt,
		})
	}

	return &v1.GetListRes{
		CurrentRes: common.BuildCurrentRes(gconv.Int(total), current, size),
		List:       list,
	}, nil
}
