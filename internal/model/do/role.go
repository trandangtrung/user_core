// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table Role for DAO operations like Where/Data.
type Role struct {
	g.Meta         `orm:"table:Role, do:true"`
	Id             interface{} //
	UserPlatformId interface{} //
	Name           interface{} //
	Description    interface{} //
	CreatedBy      interface{} //
	UpdatedBy      interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
