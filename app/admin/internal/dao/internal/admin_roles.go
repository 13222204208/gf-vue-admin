// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRolesDao is the data access object for the table admin_roles.
type AdminRolesDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminRolesColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminRolesColumns defines and stores column names for the table admin_roles.
type AdminRolesColumns struct {
	Id          string // 主键ID
	Name        string // 角色名称（英文唯一标识，如 admin, editor, viewer）
	DisplayName string // 角色显示名称（用于展示，如 系统管理员）
	Description string // 角色描述
	Status      string // 状态: active=启用, disabled=禁用
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// adminRolesColumns holds the columns for the table admin_roles.
var adminRolesColumns = AdminRolesColumns{
	Id:          "id",
	Name:        "name",
	DisplayName: "display_name",
	Description: "description",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewAdminRolesDao creates and returns a new DAO object for table data access.
func NewAdminRolesDao(handlers ...gdb.ModelHandler) *AdminRolesDao {
	return &AdminRolesDao{
		group:    "default",
		table:    "admin_roles",
		columns:  adminRolesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminRolesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminRolesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminRolesDao) Columns() AdminRolesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminRolesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminRolesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminRolesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
