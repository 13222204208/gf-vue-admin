// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRoles is the golang structure of table admin_roles for DAO operations like Where/Data.
type AdminRoles struct {
	g.Meta      `orm:"table:admin_roles, do:true"`
	Id          any         // 主键ID
	Name        any         // 角色名称（英文唯一标识，如 admin, editor, viewer）
	DisplayName any         // 角色显示名称（用于展示，如 系统管理员）
	Description any         // 角色描述
	Status      any         // 状态: active=启用, disabled=禁用
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
