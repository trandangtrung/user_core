package entity

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserID uint  `json:"user_id" gorm:"not null"`
	User   *User `json:"user" gorm:"foreignKey:UserID"`

	RoleID uint  `json:"role_id" gorm:"not null"`
	Role   *Role `json:"role" gorm:"foreignKey:RoleID"`

	CreatedBy uint `json:"created_by" gorm:"not null"`
	UpdatedBy uint `json:"updated_by" gorm:"not null"`
}
