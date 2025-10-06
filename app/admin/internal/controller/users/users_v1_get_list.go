package users

import (
	"context"

	v1 "gf-vue-admin/app/admin/api/users/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/common"
	"gf-vue-admin/app/admin/internal/model/entity"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	model := dao.AdminUsers.Ctx(ctx)

	// 添加查询条件
	if req.Username != "" {
		model = model.WhereLike(dao.AdminUsers.Columns().Username, "%"+req.Username+"%")
	}
	if req.Phone != "" {
		model = model.WhereLike(dao.AdminUsers.Columns().Phone, "%"+req.Phone+"%")
	}
	if req.Status != "" {
		model = model.Where(dao.AdminUsers.Columns().Status, req.Status)
	}

	// 获取总数
	total, err := model.Count()
	if err != nil {
		return nil, err
	}

	// 获取分页参数
	current, size := req.CurrentReq.GetCurrentInfo()

	// 关联查询用户和角色信息
	type UserWithRole struct {
		entity.AdminUsers
		RoleName string `json:"roleName"`
	}

	var usersWithRole []UserWithRole
	err = model.
		LeftJoin(dao.AdminRoles.Table()+" r", "r.id = "+dao.AdminUsers.Columns().RoleId).
		Fields(dao.AdminUsers.Table()+".*", "r.display_name as role_name").
		Page(current, size).
		OrderDesc(dao.AdminUsers.Columns().Id).
		Scan(&usersWithRole)
	if err != nil {
		return nil, err
	}

	// 转换数据格式
	var list []v1.GetByIdRes
	for _, userRole := range usersWithRole {

		list = append(list, v1.GetByIdRes{
			Id:            userRole.Id,
			Username:      userRole.Username,
			Nickname:      userRole.Nickname,
			Email:         userRole.Email,
			Phone:         userRole.Phone,
			Avatar:        userRole.Avatar,
			Status:        userRole.Status,
			RoleId:        userRole.RoleId,
			RoleName:      userRole.RoleName,
			LastLoginIp:   userRole.LastLoginIp,
			LastLoginTime: userRole.LastLoginTime,
			CreatedAt:     userRole.CreatedAt,
		})
	}

	return &v1.GetListRes{
		List:       list,
		CurrentRes: common.BuildCurrentRes(total, current, size),
	}, nil
}
