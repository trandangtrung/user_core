// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlatformDao is the data access object for the table Platform.
type PlatformDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PlatformColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PlatformColumns defines and stores column names for the table Platform.
type PlatformColumns struct {
	Id        string //
	Name      string //
	Config    string //
	CreatedBy string //
	UpdatedBy string //
	CreatedAt string //
	UpdatedAt string //
}

// platformColumns holds the columns for the table Platform.
var platformColumns = PlatformColumns{
	Id:        "id",
	Name:      "name",
	Config:    "config",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPlatformDao creates and returns a new DAO object for table data access.
func NewPlatformDao(handlers ...gdb.ModelHandler) *PlatformDao {
	return &PlatformDao{
		group:    "default",
		table:    "Platform",
		columns:  platformColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PlatformDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PlatformDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PlatformDao) Columns() PlatformColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PlatformDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PlatformDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PlatformDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
