package entity

import "gorm.io/gorm"

type UserPlatform struct {
	gorm.Model
	UserId     uint  `json:"user_id" gorm:"column:user_id" description:""` // Thêm
	User       *User `json:"user"    gorm:"foreignKey:UserId;references:ID" description:"User"`
	PlatformId int   `json:"platform_id" gorm:"column:platform_id" description:""`
	CreatedBy  int   `json:"created_by" gorm:"column:created_by" description:""`
	UpdatedBy  int   `json:"updated_by" gorm:"column:updated_by" description:""`
}
