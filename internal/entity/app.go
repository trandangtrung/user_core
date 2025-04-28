package entity

import "gorm.io/gorm"

type App struct {
	gorm.Model
	Name      string `json:"name" gorm:"type:varchar(255);not null;unique"`
	Config    string `json:"config" gorm:"type:text;not null"`
	CreatedBy *uint  `json:"created_by"`
	UpdatedBy uint   `json:"updated_by" gorm:"not null"`

	Users []User `gorm:"many2many:user_app;"`
}
