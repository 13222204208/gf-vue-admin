package v1

import (
	"gf-vue-admin/app/admin/internal/model/common"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CreateReq 创建菜单请求
type CreateReq struct {
	g.Meta        `path:"/menus" tags:"Menus" method:"post" summary:"创建菜单"`
	ParentId      uint64 `json:"parentId" v:"min:0#父级菜单ID不能小于0" dc:"父级菜单ID"`
	Title         string `json:"title" v:"required|length:1,50#菜单标题不能为空|菜单标题长度为1-50位" dc:"菜单标题"`
	Icon          string `json:"icon" dc:"菜单图标"`
	Path          string `json:"path" v:"required|length:1,200#路由路径不能为空|路由路径长度为1-200位" dc:"路由路径"`
	Component     string `json:"component" dc:"前端组件路径"`
	Redirect      string `json:"redirect" dc:"重定向路径"`
	IsLink        int    `json:"isLink" v:"in:0,1#是否外部链接值不正确" dc:"是否外部链接"`
	LinkUrl       string `json:"linkUrl" dc:"外部链接地址"`
	IsIframe      int    `json:"isIframe" v:"in:0,1#是否为iframe页面值不正确" dc:"是否为iframe页面"`
	IsCache       int    `json:"isCache" v:"in:0,1#是否缓存页面值不正确" dc:"是否缓存页面"`
	IsHide        int    `json:"isHide" v:"in:0,1#是否在菜单中隐藏值不正确" dc:"是否在菜单中隐藏"`
	IsHideTab     int    `json:"isHideTab" v:"in:0,1#是否在标签页中隐藏值不正确" dc:"是否在标签页中隐藏"`
	IsFullPage    int    `json:"isFullPage" v:"in:0,1#是否为全屏页面值不正确" dc:"是否为全屏页面"`
	IsFixedTab    int    `json:"isFixedTab" v:"in:0,1#是否固定标签页值不正确" dc:"是否固定标签页"`
	IsFirstLevel  int    `json:"isFirstLevel" v:"in:0,1#是否为一级菜单值不正确" dc:"是否为一级菜单"`
	IsAuthButton  int    `json:"isAuthButton" v:"in:0,1#是否为权限按钮值不正确" dc:"是否为权限按钮"`
	AuthMark      string `json:"authMark" dc:"权限标识（如 user:add）"`
	AuthList      string `json:"authList" dc:"操作权限列表"`
	Roles         string `json:"roles" dc:"角色权限列表"`
	ShowBadge     int    `json:"showBadge" v:"in:0,1#是否显示徽章值不正确" dc:"是否显示徽章"`
	ShowTextBadge string `json:"showTextBadge" dc:"文本徽章内容"`
	ActivePath    string `json:"activePath" dc:"激活菜单路径"`
	OrderNum      int    `json:"orderNum" v:"min:0#菜单排序不能小于0" dc:"菜单排序"`
	Status        string `json:"status" v:"in:active,disabled#状态值不正确" dc:"状态"`
	Remark        string `json:"remark" v:"max-length:500#备注长度不能超过500字符" dc:"备注"`
}

type CreateRes struct {
	g.Meta `mime:"application/json"`
	Id     uint64 `json:"id" dc:"菜单ID"`
}

// UpdateReq 更新菜单请求
type UpdateReq struct {
	g.Meta        `path:"/menus/{id}" tags:"Menus" method:"put" summary:"更新菜单"`
	Id            uint64 `json:"id" v:"required|min:1#菜单ID不能为空|菜单ID必须大于0" dc:"菜单ID"`
	ParentId      uint64 `json:"parentId" v:"min:0#父级菜单ID不能小于0" dc:"父级菜单ID"`
	Title         string `json:"title" v:"length:1,50#菜单标题长度为1-50位" dc:"菜单标题"`
	Icon          string `json:"icon" dc:"菜单图标"`
	Path          string `json:"path" v:"length:1,200#路由路径长度为1-200位" dc:"路由路径"`
	Component     string `json:"component" dc:"前端组件路径"`
	Redirect      string `json:"redirect" dc:"重定向路径"`
	IsLink        int    `json:"isLink" v:"in:0,1#是否外部链接值不正确" dc:"是否外部链接"`
	LinkUrl       string `json:"linkUrl" dc:"外部链接地址"`
	IsIframe      int    `json:"isIframe" v:"in:0,1#是否为iframe页面值不正确" dc:"是否为iframe页面"`
	IsCache       int    `json:"isCache" v:"in:0,1#是否缓存页面值不正确" dc:"是否缓存页面"`
	IsHide        int    `json:"isHide" v:"in:0,1#是否在菜单中隐藏值不正确" dc:"是否在菜单中隐藏"`
	IsHideTab     int    `json:"isHideTab" v:"in:0,1#是否在标签页中隐藏值不正确" dc:"是否在标签页中隐藏"`
	IsFullPage    int    `json:"isFullPage" v:"in:0,1#是否为全屏页面值不正确" dc:"是否为全屏页面"`
	IsFixedTab    int    `json:"isFixedTab" v:"in:0,1#是否固定标签页值不正确" dc:"是否固定标签页"`
	IsFirstLevel  int    `json:"isFirstLevel" v:"in:0,1#是否为一级菜单值不正确" dc:"是否为一级菜单"`
	IsAuthButton  int    `json:"isAuthButton" v:"in:0,1#是否为权限按钮值不正确" dc:"是否为权限按钮"`
	AuthMark      string `json:"authMark" dc:"权限标识（如 user:add）"`
	AuthList      string `json:"authList" dc:"操作权限列表"`
	Roles         string `json:"roles" dc:"角色权限列表"`
	ShowBadge     int    `json:"showBadge" v:"in:0,1#是否显示徽章值不正确" dc:"是否显示徽章"`
	ShowTextBadge string `json:"showTextBadge" dc:"文本徽章内容"`
	ActivePath    string `json:"activePath" dc:"激活菜单路径"`
	OrderNum      int    `json:"orderNum" v:"min:0#菜单排序不能小于0" dc:"菜单排序"`
	Status        string `json:"status" v:"in:active,disabled#状态值不正确" dc:"状态"`
	Remark        string `json:"remark" v:"max-length:500#备注长度不能超过500字符" dc:"备注"`
}

type UpdateRes struct {
	g.Meta `mime:"application/json"`
}

// DeleteReq 删除菜单请求
type DeleteReq struct {
	g.Meta `path:"/menus/{id}" tags:"Menus" method:"delete" summary:"删除菜单"`
	Id     uint64 `json:"id" v:"required|min:1#菜单ID不能为空|菜单ID必须大于0" dc:"菜单ID"`
}

type DeleteRes struct {
	g.Meta `mime:"application/json"`
}

// GetByIdReq 根据ID获取菜单请求
type GetByIdReq struct {
	g.Meta `path:"/menus/{id}" tags:"Menus" method:"get" summary:"根据ID获取菜单"`
	Id     uint64 `json:"id" v:"required|min:1#菜单ID不能为空|菜单ID必须大于0" dc:"菜单ID"`
}

type GetByIdRes struct {
	g.Meta        `mime:"application/json"`
	Id            uint64      `json:"id" dc:"主键ID"`
	ParentId      uint64      `json:"parentId" dc:"父级菜单ID"`
	Title         string      `json:"title" dc:"菜单标题"`
	Icon          string      `json:"icon" dc:"菜单图标"`
	Path          string      `json:"path" dc:"路由路径"`
	Component     string      `json:"component" dc:"前端组件路径"`
	Redirect      string      `json:"redirect" dc:"重定向路径"`
	IsLink        int         `json:"isLink" dc:"是否外部链接"`
	LinkUrl       string      `json:"linkUrl" dc:"外部链接地址"`
	IsIframe      int         `json:"isIframe" dc:"是否为iframe页面"`
	IsCache       int         `json:"isCache" dc:"是否缓存页面"`
	IsHide        int         `json:"isHide" dc:"是否在菜单中隐藏"`
	IsHideTab     int         `json:"isHideTab" dc:"是否在标签页中隐藏"`
	IsFullPage    int         `json:"isFullPage" dc:"是否为全屏页面"`
	IsFixedTab    int         `json:"isFixedTab" dc:"是否固定标签页"`
	IsFirstLevel  int         `json:"isFirstLevel" dc:"是否为一级菜单"`
	IsAuthButton  int         `json:"isAuthButton" dc:"是否为权限按钮"`
	AuthMark      string      `json:"authMark" dc:"权限标识"`
	AuthList      string      `json:"authList" dc:"操作权限列表"`
	Roles         string      `json:"roles" dc:"角色权限列表"`
	ShowBadge     int         `json:"showBadge" dc:"是否显示徽章"`
	ShowTextBadge string      `json:"showTextBadge" dc:"文本徽章内容"`
	ActivePath    string      `json:"activePath" dc:"激活菜单路径"`
	OrderNum      int         `json:"orderNum" dc:"菜单排序"`
	Status        string      `json:"status" dc:"状态"`
	Remark        string      `json:"remark" dc:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// GetListReq 获取菜单列表请求
type GetListReq struct {
	g.Meta `path:"/menus" tags:"Menus" method:"get" summary:"获取菜单列表"`
	common.CurrentReq
	Title        string `json:"title" dc:"菜单标题（模糊搜索）"`
	Path         string `json:"path" dc:"路由路径（模糊搜索）"`
	Status       string `json:"status" v:"in:active,disabled#状态值不正确" dc:"状态筛选"`
	ParentId     uint64 `json:"parentId" dc:"父级菜单ID"`
	IsAuthButton int    `json:"isAuthButton" v:"in:0,1#是否为权限按钮值不正确" dc:"是否为权限按钮"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	common.CurrentRes
	List []GetByIdRes `json:"list" dc:"菜单列表"`
}

// GetTreeReq 获取菜单树请求
type GetTreeReq struct {
	g.Meta `path:"/menus/tree" tags:"Menus" method:"get" summary:"获取菜单树"`
}

type GetTreeRes struct {
	g.Meta `mime:"application/json"`
	List   []MenuTreeItem `json:"list" dc:"菜单树列表"`
}

// MenuTreeItem 菜单树项
type MenuTreeItem struct {
	GetByIdRes
	Children []MenuTreeItem `json:"children" dc:"子菜单列表"`
}

// GetUserMenusReq 获取用户菜单请求
type GetUserMenusReq struct {
	g.Meta `path:"/menus/user" tags:"Menus" method:"get" summary:"获取用户菜单"`
}

type GetUserMenusRes struct {
	g.Meta `mime:"application/json"`
	List   []MenuTreeItem `json:"list" dc:"用户菜单树"`
}
