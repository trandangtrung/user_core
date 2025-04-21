// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRole is the golang structure of table UserRole for DAO operations like Where/Data.
type UserRole struct {
	g.Meta    `orm:"table:UserRole, do:true"`
	Id        interface{} //
	UserId    interface{} //
	RoleId    interface{} //
	CreatedBy interface{} //
	UpdatedBy interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
