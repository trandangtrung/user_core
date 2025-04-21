// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Session is the golang structure of table Session for DAO operations like Where/Data.
type Session struct {
	g.Meta       `orm:"table:Session, do:true"`
	Id           interface{} //
	UserId       interface{} //
	RefreshToken interface{} //
	Scope        interface{} //
	CreateBy     interface{} //
	UpdateBy     interface{} //
	CreateAt     *gtime.Time //
	UpdateAt     *gtime.Time //
}
