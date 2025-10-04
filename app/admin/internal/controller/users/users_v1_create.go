package users

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "gf-vue-admin/app/admin/api/users/v1"
	"gf-vue-admin/app/admin/internal/consts"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/do"
	"gf-vue-admin/utility/password"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 检查用户名是否已存在
	count, err := dao.AdminUsers.Ctx(ctx).Where(dao.AdminUsers.Columns().Username, req.Username).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("用户名已存在")
	}

	// 密码加密
	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 设置默认值
	if req.Status == "" {
		req.Status = consts.UserDefaultStatus
	}

	// 插入数据
	insertId, err := dao.AdminUsers.Ctx(ctx).Data(do.AdminUsers{
		Username: req.Username,
		Password: hashedPassword,
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Avatar:   req.Avatar,
		Status:   req.Status,
		RoleId:   req.RoleId,
	}).InsertAndGetId()

	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{
		Id: gconv.Uint64(insertId),
	}, nil
}
