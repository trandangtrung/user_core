// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table Users.
type Users struct {
	Id           int64       `json:"id"           orm:"id"            description:""` //
	Email        string      `json:"email"        orm:"email"         description:""` //
	PasswordHash string      `json:"passwordHash" orm:"password_hash" description:""` //
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""` //
}
