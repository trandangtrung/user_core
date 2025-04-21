// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRole is the golang structure for table UserRole.
type UserRole struct {
	Id        int64       `json:"id"        orm:"id"         description:""` //
	UserId    int         `json:"userId"    orm:"user_id"    description:""` //
	RoleId    int         `json:"roleId"    orm:"role_id"    description:""` //
	CreatedBy int         `json:"createdBy" orm:"created_by" description:""` //
	UpdatedBy int         `json:"updatedBy" orm:"updated_by" description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
}
