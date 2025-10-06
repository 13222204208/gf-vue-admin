// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id            uint64      `json:"id"            orm:"id"              description:"主键ID"`                                                     // 主键ID
	ParentId      uint64      `json:"parentId"      orm:"parent_id"       description:"父级菜单ID"`                                                   // 父级菜单ID
	Title         string      `json:"title"         orm:"title"           description:"菜单标题"`                                                     // 菜单标题
	Icon          string      `json:"icon"          orm:"icon"            description:"菜单图标"`                                                     // 菜单图标
	Path          string      `json:"path"          orm:"path"            description:"路由路径"`                                                     // 路由路径
	Component     string      `json:"component"     orm:"component"       description:"前端组件路径"`                                                   // 前端组件路径
	Redirect      string      `json:"redirect"      orm:"redirect"        description:"重定向路径"`                                                    // 重定向路径
	IsLink        int         `json:"isLink"        orm:"is_link"         description:"是否外部链接"`                                                   // 是否外部链接
	LinkUrl       string      `json:"linkUrl"       orm:"link_url"        description:"外部链接地址"`                                                   // 外部链接地址
	IsIframe      int         `json:"isIframe"      orm:"is_iframe"       description:"是否为iframe页面"`                                              // 是否为iframe页面
	IsCache       int         `json:"isCache"       orm:"is_cache"        description:"是否缓存页面"`                                                   // 是否缓存页面
	IsHide        int         `json:"isHide"        orm:"is_hide"         description:"是否在菜单中隐藏"`                                                 // 是否在菜单中隐藏
	IsHideTab     int         `json:"isHideTab"     orm:"is_hide_tab"     description:"是否在标签页中隐藏"`                                                // 是否在标签页中隐藏
	IsFullPage    int         `json:"isFullPage"    orm:"is_full_page"    description:"是否为全屏页面"`                                                  // 是否为全屏页面
	IsFixedTab    int         `json:"isFixedTab"    orm:"is_fixed_tab"    description:"是否固定标签页"`                                                  // 是否固定标签页
	IsFirstLevel  int         `json:"isFirstLevel"  orm:"is_first_level"  description:"是否为一级菜单"`                                                  // 是否为一级菜单
	IsAuthButton  int         `json:"isAuthButton"  orm:"is_auth_button"  description:"是否为权限按钮"`                                                  // 是否为权限按钮
	AuthMark      string      `json:"authMark"      orm:"auth_mark"       description:"权限标识（如 user:add）"`                                         // 权限标识（如 user:add）
	AuthList      string      `json:"authList"      orm:"auth_list"       description:"操作权限列表（如 [{\"title\":\"新增\",\"authMark\":\"user:add\"}]）"` // 操作权限列表（如 [{"title":"新增","authMark":"user:add"}]）
	Roles         string      `json:"roles"         orm:"roles"           description:"角色权限列表（如 [\"admin\",\"editor\"]）"`                         // 角色权限列表（如 ["admin","editor"]）
	ShowBadge     int         `json:"showBadge"     orm:"show_badge"      description:"是否显示徽章"`                                                   // 是否显示徽章
	ShowTextBadge string      `json:"showTextBadge" orm:"show_text_badge" description:"文本徽章内容"`                                                   // 文本徽章内容
	ActivePath    string      `json:"activePath"    orm:"active_path"     description:"激活菜单路径"`                                                   // 激活菜单路径
	OrderNum      int         `json:"orderNum"      orm:"order_num"       description:"菜单排序"`                                                     // 菜单排序
	Status        string      `json:"status"        orm:"status"          description:"状态"`                                                       // 状态
	Remark        string      `json:"remark"        orm:"remark"          description:"备注"`                                                       // 备注
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`                                                     // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`                                                     // 更新时间
}
