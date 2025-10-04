// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUsers is the golang structure of table admin_users for DAO operations like Where/Data.
type AdminUsers struct {
	g.Meta        `orm:"table:admin_users, do:true"`
	Id            any         // 主键ID
	Username      any         // 用户名（唯一）
	Password      any         // 密码（加密后存储）
	Nickname      any         // 昵称
	Email         any         // 邮箱
	Phone         any         // 手机号
	Avatar        any         // 头像URL
	Status        any         // 账号状态：active=启用，disabled=禁用
	RoleId        any         // 角色ID（可扩展关联角色表）
	LastLoginIp   any         // 最后登录IP
	LastLoginTime *gtime.Time // 最后登录时间
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
