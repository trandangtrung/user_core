// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tokens is the golang structure of table Tokens for DAO operations like Where/Data.
type Tokens struct {
	g.Meta       `orm:"table:Tokens, do:true"`
	Id           interface{} //
	UserId       interface{} //
	RefreshToken interface{} //
	Scope        interface{} //
	CreatedBy    interface{} //
	UpdatedBy    interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
