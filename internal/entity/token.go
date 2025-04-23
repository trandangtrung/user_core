package entity

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	UserID uint  `json:"user_id" gorm:"not null"`
	User   *User `json:"user" gorm:"foreignKey:UserID"`

	RefreshToken string `json:"refresh_token" gorm:"type:text;not null"`
	Scope        string `json:"scope" gorm:"type:text;not null"`

	CreatedBy uint `json:"created_by" gorm:"not null"`
	UpdatedBy uint `json:"updated_by" gorm:"not null"`
}
