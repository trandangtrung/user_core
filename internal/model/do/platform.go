// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Platform is the golang structure of table Platform for DAO operations like Where/Data.
type Platform struct {
	g.Meta    `orm:"table:Platform, do:true"`
	Id        interface{} //
	Name      interface{} //
	Config    interface{} //
	CreatedBy interface{} //
	UpdatedBy interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
