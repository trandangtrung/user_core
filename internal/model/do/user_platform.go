// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserPlatform is the golang structure of table UserPlatform for DAO operations like Where/Data.
type UserPlatform struct {
	g.Meta     `orm:"table:UserPlatform, do:true"`
	Id         interface{} //
	UserId     interface{} //
	PlatformId interface{} //
	CreatedBy  interface{} //
	UpdatedBy  interface{} //
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
