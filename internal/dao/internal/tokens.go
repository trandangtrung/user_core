// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TokensDao is the data access object for the table Tokens.
type TokensDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TokensColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TokensColumns defines and stores column names for the table Tokens.
type TokensColumns struct {
	Id           string //
	UserId       string //
	RefreshToken string //
	Scope        string //
	CreatedBy    string //
	UpdatedBy    string //
	CreatedAt    string //
	UpdatedAt    string //
}

// tokensColumns holds the columns for the table Tokens.
var tokensColumns = TokensColumns{
	Id:           "id",
	UserId:       "user_id",
	RefreshToken: "refresh_token",
	Scope:        "scope",
	CreatedBy:    "created_by",
	UpdatedBy:    "updated_by",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewTokensDao creates and returns a new DAO object for table data access.
func NewTokensDao(handlers ...gdb.ModelHandler) *TokensDao {
	return &TokensDao{
		group:    "default",
		table:    "Tokens",
		columns:  tokensColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TokensDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TokensDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TokensDao) Columns() TokensColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TokensDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TokensDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TokensDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
