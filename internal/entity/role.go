package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	PlatformID  uint      `json:"platform_id" gorm:"not null"`
	Platform    *Platform `json:"platform" gorm:"foreignKey:PlatformID"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:varchar(255);not null"`
	CreatedBy   uint      `json:"created_by" gorm:"not null"`
	UpdatedBy   uint      `json:"updated_by" gorm:"not null"`
}
