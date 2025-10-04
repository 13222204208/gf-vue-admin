package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LoginReq 登录请求
type LoginReq struct {
	g.Meta   `path:"/auth/login" tags:"Auth" method:"post" summary:"用户登录"`
	UserName string `json:"userName" v:"required|length:3,20#用户名不能为空|用户名长度为3-20位" dc:"用户名"`
	Password string `json:"password" v:"required|length:6,20#密码不能为空|密码长度为6-20位" dc:"密码"`
}

type LoginRes struct {
	g.Meta       `mime:"application/json"`
	Token        string `json:"token" dc:"访问令牌"`
	RefreshToken string `json:"refreshToken" dc:"刷新令牌"`
}

// GetUserInfoReq 获取用户信息请求
type GetUserInfoReq struct {
	g.Meta `path:"/user/info" tags:"Auth" method:"get" summary:"获取用户信息"`
}

type GetUserInfoRes struct {
	g.Meta   `mime:"application/json"`
	Buttons  []string `json:"buttons" dc:"按钮权限列表"`
	Roles    []string `json:"roles" dc:"角色列表"`
	UserId   uint64   `json:"userId" dc:"用户ID"`
	UserName string   `json:"userName" dc:"用户名"`
	Email    string   `json:"email" dc:"邮箱"`
	Avatar   string   `json:"avatar" dc:"头像URL"`
}

// RefreshTokenReq 刷新令牌请求
type RefreshTokenReq struct {
	g.Meta       `path:"/auth/refresh_token" tags:"Auth" method:"post" summary:"刷新访问令牌"`
	RefreshToken string `json:"refreshToken" v:"required#刷新令牌不能为空" dc:"刷新令牌"`
}

type RefreshTokenRes struct {
	g.Meta       `mime:"application/json"`
	Token        string `json:"token" dc:"新的访问令牌"`
	RefreshToken string `json:"refreshToken" dc:"新的刷新令牌"`
}