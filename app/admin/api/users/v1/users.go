package v1

import (
	"gf-vue-admin/app/admin/internal/model/common"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CreateReq 创建用户请求
type CreateReq struct {
	g.Meta   `path:"/users" tags:"Users" method:"post" summary:"创建用户"`
	Username string `json:"username" v:"required|length:3,20#用户名不能为空|用户名长度为3-20位" dc:"用户名"`
	Password string `json:"password" v:"required|length:6,20#密码不能为空|密码长度为6-20位" dc:"密码"`
	Nickname string `json:"nickname" v:"required|length:2,20#昵称不能为空|昵称长度为2-20位" dc:"昵称"`
	Email    string `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
	Phone    string `json:"phone" v:"phone#手机号格式不正确" dc:"手机号"`
	Avatar   string `json:"avatar" dc:"头像URL"`
	Status   string `json:"status" v:"in:active,disabled#状态值不正确" dc:"账号状态"`
	RoleId   int64  `json:"roleId" dc:"角色ID"`
}

type CreateRes struct {
	g.Meta `mime:"application/json"`
	Id     uint64 `json:"id" dc:"用户ID"`
}

// UpdateReq 更新用户请求
type UpdateReq struct {
	g.Meta   `path:"/users/{id}" tags:"Users" method:"put" summary:"更新用户"`
	Id       uint64 `json:"id" v:"required|min:1#用户ID不能为空|用户ID必须大于0" dc:"用户ID"`
	Username string `json:"username" v:"length:3,20#用户名长度为3-20位" dc:"用户名"`
	Nickname string `json:"nickname" v:"length:0,20#昵称长度为2-20位" dc:"昵称"`
	Email    string `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
	Phone    string `json:"phone" v:"phone#手机号格式不正确" dc:"手机号"`
	Avatar   string `json:"avatar" dc:"头像URL"`
	Status   string `json:"status" v:"in:active,disabled#状态值不正确" dc:"账号状态"`
	RoleId   int64  `json:"roleId" dc:"角色ID"`
}

type UpdateRes struct {
	g.Meta `mime:"application/json"`
}

// DeleteReq 删除用户请求
type DeleteReq struct {
	g.Meta `path:"/users/{id}" tags:"Users" method:"delete" summary:"删除用户"`
	Id     uint64 `json:"id" v:"required|min:1#用户ID不能为空|用户ID必须大于0" dc:"用户ID"`
}

type DeleteRes struct {
	g.Meta `mime:"application/json"`
}

// GetByIdReq 根据ID获取用户请求
type GetByIdReq struct {
	g.Meta `path:"/users/{id}" tags:"Users" method:"get" summary:"根据ID获取用户"`
	Id     uint64 `json:"id" v:"required|min:1#用户ID不能为空|用户ID必须大于0" dc:"用户ID"`
}

type GetByIdRes struct {
	g.Meta        `mime:"application/json"`
	Id            uint64      `json:"id" dc:"主键ID"`
	Username      string      `json:"username" dc:"用户名"`
	Nickname      string      `json:"nickname" dc:"昵称"`
	Email         string      `json:"email" dc:"邮箱"`
	Phone         string      `json:"phone" dc:"手机号"`
	Avatar        string      `json:"avatar" dc:"头像URL"`
	Status        string      `json:"status" dc:"账号状态"`
	RoleId        int64       `json:"roleId" dc:"角色ID"`
	RoleName      string      `json:"roleName" dc:"角色名称"`
	LastLoginIp   string      `json:"lastLoginIp" dc:"最后登录IP"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" dc:"最后登录时间"`
	CreatedAt     *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// GetListReq 获取用户列表请求
type GetListReq struct {
	g.Meta `path:"/users" tags:"Users" method:"get" summary:"获取用户列表"`
	common.CurrentReq
	Username string `json:"username" dc:"用户名（模糊搜索）"`
	Phone    string `json:"phone" dc:"手机号（模糊搜索）"`
	Status   string `json:"status" v:"in:active,disabled#状态值不正确" dc:"账号状态"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	common.CurrentRes
	List []GetByIdRes `json:"list" dc:"用户列表"`
}
