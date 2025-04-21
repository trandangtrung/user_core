// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Session is the golang structure for table Session.
type Session struct {
	Id           int64       `json:"id"           orm:"id"            description:""` //
	UserId       int         `json:"userId"       orm:"user_id"       description:""` //
	RefreshToken string      `json:"refreshToken" orm:"refresh_token" description:""` //
	Scope        string      `json:"scope"        orm:"scope"         description:""` //
	CreateBy     int         `json:"createBy"     orm:"createBy"      description:""` //
	UpdateBy     int         `json:"updateBy"     orm:"updateBy"      description:""` //
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"     description:""` //
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"     description:""` //
}
