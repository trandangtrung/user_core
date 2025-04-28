package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(255);not null"`

	AppID uint `json:"app_id" gorm:"not null"`
	App   *App `json:"app" gorm:"foreignKey:AppID"`

	CreatedBy *uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by" gorm:"not null"`

	Users []User `gorm:"many2many:user_role;"`
}
