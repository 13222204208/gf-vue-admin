package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "gf-vue-admin/app/admin/api/auth/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/do"
	"gf-vue-admin/app/admin/internal/model/entity"
	"gf-vue-admin/utility/jwt"
	"gf-vue-admin/utility/password"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	// 查找用户
	var user *entity.AdminUsers
	err = dao.AdminUsers.Ctx(ctx).Where(dao.AdminUsers.Columns().Username, req.UserName).Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("用户名或密码错误")
	}

	// 检查用户状态
	if user.Status != "active" {
		return nil, gerror.New("账号已被禁用")
	}

	// 验证密码
	if !password.CheckPasswordHash(req.Password, user.Password) {
		return nil, gerror.New("用户名或密码错误")
	}

	// 更新最后登录信息
	clientIp := g.RequestFromCtx(ctx).GetClientIp()
	_, err = dao.AdminUsers.Ctx(ctx).Where(dao.AdminUsers.Columns().Id, user.Id).Data(do.AdminUsers{
		LastLoginIp:   clientIp,
		LastLoginTime: gtime.Now(),
	}).Update()
	if err != nil {
		return nil, err
	}

	// 使用JWT生成token
	accessToken, refreshToken, err := jwt.GenerateTokens(user.Id, user.Username, user.Email)
	if err != nil {
		return nil, gerror.New("生成令牌失败: " + err.Error())
	}

	return &v1.LoginRes{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}
