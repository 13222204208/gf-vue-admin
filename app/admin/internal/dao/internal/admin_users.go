// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminUsersDao is the data access object for the table admin_users.
type AdminUsersDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AdminUsersColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AdminUsersColumns defines and stores column names for the table admin_users.
type AdminUsersColumns struct {
	Id            string // 主键ID
	Username      string // 用户名（唯一）
	Password      string // 密码（加密后存储）
	Nickname      string // 昵称
	Email         string // 邮箱
	Phone         string // 手机号
	Avatar        string // 头像URL
	Status        string // 账号状态：active=启用，disabled=禁用
	RoleId        string // 角色ID（可扩展关联角色表）
	LastLoginIp   string // 最后登录IP
	LastLoginTime string // 最后登录时间
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// adminUsersColumns holds the columns for the table admin_users.
var adminUsersColumns = AdminUsersColumns{
	Id:            "id",
	Username:      "username",
	Password:      "password",
	Nickname:      "nickname",
	Email:         "email",
	Phone:         "phone",
	Avatar:        "avatar",
	Status:        "status",
	RoleId:        "role_id",
	LastLoginIp:   "last_login_ip",
	LastLoginTime: "last_login_time",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewAdminUsersDao creates and returns a new DAO object for table data access.
func NewAdminUsersDao(handlers ...gdb.ModelHandler) *AdminUsersDao {
	return &AdminUsersDao{
		group:    "default",
		table:    "admin_users",
		columns:  adminUsersColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminUsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminUsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminUsersDao) Columns() AdminUsersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminUsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminUsersDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminUsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
