// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRoles is the golang structure for table admin_roles.
type AdminRoles struct {
	Id          uint64      `json:"id"          orm:"id"           description:"主键ID"`                                 // 主键ID
	Name        string      `json:"name"        orm:"name"         description:"角色名称（英文唯一标识，如 admin, editor, viewer）"` // 角色名称（英文唯一标识，如 admin, editor, viewer）
	DisplayName string      `json:"displayName" orm:"display_name" description:"角色显示名称（用于展示，如 系统管理员）"`                 // 角色显示名称（用于展示，如 系统管理员）
	Description string      `json:"description" orm:"description"  description:"角色描述"`                                 // 角色描述
	Status      string      `json:"status"      orm:"status"       description:"状态: active=启用, disabled=禁用"`           // 状态: active=启用, disabled=禁用
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`                                 // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`                                 // 更新时间
}
