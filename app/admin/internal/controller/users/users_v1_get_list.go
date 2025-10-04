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
	if req.Nickname != "" {
		model = model.WhereLike(dao.AdminUsers.Columns().Nickname, "%"+req.Nickname+"%")
	}
	if req.Status != "" {
		model = model.Where(dao.AdminUsers.Columns().Status, req.Status)
	}
	if req.RoleId > 0 {
		model = model.Where(dao.AdminUsers.Columns().RoleId, req.RoleId)
	}

	// 获取总数
	total, err := model.Count()
	if err != nil {
		return nil, err
	}
	// 获取分页参数
	current, size := req.CurrentReq.GetcurrentInfo()
	// 分页查询
	var users []entity.AdminUsers
	err = model.Page(current, size).OrderDesc(dao.AdminUsers.Columns().Id).Scan(&users)
	if err != nil {
		return nil, err
	}

	// 转换数据格式
	var list []v1.GetByIdRes
	for _, user := range users {
		// 模拟用户角色，实际项目中应该从角色表查询
		userRoles := []string{"admin"}
		if user.RoleId == 2 {
			userRoles = []string{"user"}
		}

		list = append(list, v1.GetByIdRes{
			Id:            user.Id,
			Username:      user.Username,
			Nickname:      user.Nickname,
			Email:         user.Email,
			Phone:         user.Phone,
			Avatar:        user.Avatar,
			Status:        user.Status,
			RoleId:        user.RoleId,
			UserRoles:     userRoles,
			LastLoginIp:   user.LastLoginIp,
			LastLoginTime: user.LastLoginTime,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
		})
	}

	return &v1.GetListRes{
		List:       list,
		CurrentRes: common.BuildCurrentRes(total, current, size),
	}, nil
}
