package users

import (
	"context"

	v1 "gf-vue-admin/app/admin/api/users/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetById(ctx context.Context, req *v1.GetByIdReq) (res *v1.GetByIdRes, err error) {
	var user *entity.AdminUsers
	err = dao.AdminUsers.Ctx(ctx).Where(dao.AdminUsers.Columns().Id, req.Id).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("用户不存在")
	}

	return &v1.GetByIdRes{
		Id:            user.Id,
		Username:      user.Username,
		Nickname:      user.Nickname,
		Email:         user.Email,
		Phone:         user.Phone,
		Avatar:        user.Avatar,
		Status:        user.Status,
		RoleId:        user.RoleId,
		LastLoginIp:   user.LastLoginIp,
		LastLoginTime: user.LastLoginTime,
		CreatedAt:     user.CreatedAt,
	}, nil
}
