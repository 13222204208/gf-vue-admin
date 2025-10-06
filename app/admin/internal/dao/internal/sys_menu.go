// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenuDao is the data access object for the table sys_menu.
type SysMenuDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysMenuColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysMenuColumns defines and stores column names for the table sys_menu.
type SysMenuColumns struct {
	Id            string // 主键ID
	ParentId      string // 父级菜单ID
	Title         string // 菜单标题
	Icon          string // 菜单图标
	Path          string // 路由路径
	Component     string // 前端组件路径
	Redirect      string // 重定向路径
	IsLink        string // 是否外部链接
	LinkUrl       string // 外部链接地址
	IsIframe      string // 是否为iframe页面
	IsCache       string // 是否缓存页面
	IsHide        string // 是否在菜单中隐藏
	IsHideTab     string // 是否在标签页中隐藏
	IsFullPage    string // 是否为全屏页面
	IsFixedTab    string // 是否固定标签页
	IsFirstLevel  string // 是否为一级菜单
	IsAuthButton  string // 是否为权限按钮
	AuthMark      string // 权限标识（如 user:add）
	AuthList      string // 操作权限列表（如 [{"title":"新增","authMark":"user:add"}]）
	Roles         string // 角色权限列表（如 ["admin","editor"]）
	ShowBadge     string // 是否显示徽章
	ShowTextBadge string // 文本徽章内容
	ActivePath    string // 激活菜单路径
	OrderNum      string // 菜单排序
	Status        string // 状态
	Remark        string // 备注
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// sysMenuColumns holds the columns for the table sys_menu.
var sysMenuColumns = SysMenuColumns{
	Id:            "id",
	ParentId:      "parent_id",
	Title:         "title",
	Icon:          "icon",
	Path:          "path",
	Component:     "component",
	Redirect:      "redirect",
	IsLink:        "is_link",
	LinkUrl:       "link_url",
	IsIframe:      "is_iframe",
	IsCache:       "is_cache",
	IsHide:        "is_hide",
	IsHideTab:     "is_hide_tab",
	IsFullPage:    "is_full_page",
	IsFixedTab:    "is_fixed_tab",
	IsFirstLevel:  "is_first_level",
	IsAuthButton:  "is_auth_button",
	AuthMark:      "auth_mark",
	AuthList:      "auth_list",
	Roles:         "roles",
	ShowBadge:     "show_badge",
	ShowTextBadge: "show_text_badge",
	ActivePath:    "active_path",
	OrderNum:      "order_num",
	Status:        "status",
	Remark:        "remark",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewSysMenuDao creates and returns a new DAO object for table data access.
func NewSysMenuDao(handlers ...gdb.ModelHandler) *SysMenuDao {
	return &SysMenuDao{
		group:    "default",
		table:    "sys_menu",
		columns:  sysMenuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysMenuDao) Columns() SysMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysMenuDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
