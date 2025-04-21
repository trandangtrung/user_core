// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Platform is the golang structure for table Platform.
type Platform struct {
	Id        int64       `json:"id"        orm:"id"         description:""` //
	Name      string      `json:"name"      orm:"name"       description:""` //
	Config    string      `json:"config"    orm:"config"     description:""` //
	CreatedBy int         `json:"createdBy" orm:"created_by" description:""` //
	UpdatedBy int         `json:"updatedBy" orm:"updated_by" description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
}
