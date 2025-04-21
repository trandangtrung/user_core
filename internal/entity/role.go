package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	PlatformId  uint      `json:"platform_id" gorm:"column:platform_id;not null;" description:""` // Thêm field khóa ngoại
	Platform    *Platform `json:"platform"    gorm:"foreignKey:PlatformId;references:ID" description:""`
	Name        string    `json:"name"        gorm:"column:name;type:varchar(255);not null;" description:""`
	Description string    `json:"description" gorm:"column:description;type:varchar(255);not null;" description:""`
	CreatedBy   int       `json:"createdBy"   gorm:"column:created_by;not null;" description:""`
	UpdatedBy   int       `json:"updatedBy"   gorm:"column:updated_by;not null;" description:""`
}
