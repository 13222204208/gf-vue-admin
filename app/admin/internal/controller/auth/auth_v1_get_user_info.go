package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "gf-vue-admin/app/admin/api/auth/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/entity"
)

func (c *ControllerV1) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	// 从上下文中获取用户信息（由JWT中间件设置）
	request := g.RequestFromCtx(ctx)
	userIDVar := request.GetCtxVar("user_id")
	usernameVar := request.GetCtxVar("username")
	emailVar := request.GetCtxVar("email")

	if userIDVar == nil || usernameVar == nil || emailVar == nil {
		return nil, gerror.New("用户信息获取失败")
	}

	userID := userIDVar.Uint64()
	username := usernameVar.String()
	email := emailVar.String()

	// 根据用户ID查询完整用户信息
	var user *entity.AdminUsers
	err = dao.AdminUsers.Ctx(ctx).Where(dao.AdminUsers.Columns().Id, userID).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("用户不存在")
	}

	// 构建用户信息响应
	return &v1.GetUserInfoRes{
		Buttons:  []string{"add", "edit", "delete", "export"}, // 模拟按钮权限
		Roles:    []string{"R_SUPER"},                         // 模拟角色权限
		UserId:   userID,
		UserName: username,
		Email:    email,
		Avatar:   user.Avatar,
	}, nil
}
