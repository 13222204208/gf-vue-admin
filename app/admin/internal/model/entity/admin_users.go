// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUsers is the golang structure for table admin_users.
type AdminUsers struct {
	Id            uint64      `json:"id"            orm:"id"              description:"主键ID"`                       // 主键ID
	Username      string      `json:"username"      orm:"username"        description:"用户名（唯一）"`                    // 用户名（唯一）
	Password      string      `json:"password"      orm:"password"        description:"密码（加密后存储）"`                  // 密码（加密后存储）
	Nickname      string      `json:"nickname"      orm:"nickname"        description:"昵称"`                         // 昵称
	Email         string      `json:"email"         orm:"email"           description:"邮箱"`                         // 邮箱
	Phone         string      `json:"phone"         orm:"phone"           description:"手机号"`                        // 手机号
	Avatar        string      `json:"avatar"        orm:"avatar"          description:"头像URL"`                      // 头像URL
	Status        string      `json:"status"        orm:"status"          description:"账号状态：active=启用，disabled=禁用"` // 账号状态：active=启用，disabled=禁用
	RoleId        int64       `json:"roleId"        orm:"role_id"         description:"角色ID（可扩展关联角色表）"`             // 角色ID（可扩展关联角色表）
	LastLoginIp   string      `json:"lastLoginIp"   orm:"last_login_ip"   description:"最后登录IP"`                     // 最后登录IP
	LastLoginTime *gtime.Time `json:"lastLoginTime" orm:"last_login_time" description:"最后登录时间"`                     // 最后登录时间
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`                       // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`                       // 更新时间
}
