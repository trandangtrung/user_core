package entity

import "gorm.io/gorm"

type UserPlatform struct {
	gorm.Model
	UserID uint  `json:"user_id" gorm:"not null"`
	User   *User `json:"user" gorm:"foreignKey:UserID"`

	PlatformID uint      `json:"platform_id" gorm:"not null"`
	Platform   *Platform `json:"platform" gorm:"foreignKey:PlatformID"`

	CreatedBy uint `json:"created_by" gorm:"not null"`
	UpdatedBy uint `json:"updated_by" gorm:"not null"`
}
