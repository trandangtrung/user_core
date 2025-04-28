package entity

import "gorm.io/gorm"

type Token struct {
	gorm.Model

	RefreshToken string `json:"refresh_token" gorm:"type:text;not null"`
	Scope        string `json:"scope" gorm:"type:text;not null"`
	CreatedBy    uint   `json:"created_by" gorm:"not null"`
	UpdatedBy    uint   `json:"updated_by" gorm:"not null"`

	UserID uint  `json:"user_id" gorm:"not null"`

}
