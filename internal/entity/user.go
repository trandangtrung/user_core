package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string `json:"email" gorm:"column:email;type:varchar(255);not null;unique" description:""`
	PasswordHashed string `json:"password_hashed" gorm:"column:password_hashed;type:varchar(255);not null;" description:""`
}
