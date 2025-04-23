package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string `json:"name" gorm:"type:varchar(255);not null"`
	Email          string `json:"email" gorm:"type:varchar(255);not null;unique"`
	PasswordHashed string `json:"password_hashed" gorm:"type:varchar(255);not null"`
	CreatedBy      uint   `json:"created_by" gorm:"not null"`
	UpdatedBy      uint   `json:"updated_by" gorm:"not null"`
}
