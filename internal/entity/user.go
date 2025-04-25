package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `json:"email" gorm:"type:varchar(255);not null;unique"`
	PasswordHashed string `json:"password_hashed" gorm:"type:varchar(255);not null"`
}
