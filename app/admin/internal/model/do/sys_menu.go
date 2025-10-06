// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta        `orm:"table:sys_menu, do:true"`
	Id            any         // 主键ID
	ParentId      any         // 父级菜单ID
	Title         any         // 菜单标题
	Icon          any         // 菜单图标
	Path          any         // 路由路径
	Component     any         // 前端组件路径
	Redirect      any         // 重定向路径
	IsLink        any         // 是否外部链接
	LinkUrl       any         // 外部链接地址
	IsIframe      any         // 是否为iframe页面
	IsCache       any         // 是否缓存页面
	IsHide        any         // 是否在菜单中隐藏
	IsHideTab     any         // 是否在标签页中隐藏
	IsFullPage    any         // 是否为全屏页面
	IsFixedTab    any         // 是否固定标签页
	IsFirstLevel  any         // 是否为一级菜单
	IsAuthButton  any         // 是否为权限按钮
	AuthMark      any         // 权限标识（如 user:add）
	AuthList      any         // 操作权限列表（如 [{"title":"新增","authMark":"user:add"}]）
	Roles         any         // 角色权限列表（如 ["admin","editor"]）
	ShowBadge     any         // 是否显示徽章
	ShowTextBadge any         // 文本徽章内容
	ActivePath    any         // 激活菜单路径
	OrderNum      any         // 菜单排序
	Status        any         // 状态
	Remark        any         // 备注
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
