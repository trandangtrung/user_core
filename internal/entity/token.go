package entity

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	UserId       uint   `json:"user_id" gorm:"column:user_id" description:""` // Thêm
	User         *User  `json:"user"    gorm:"foreignKey:UserId;references:ID" description:"User"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token;type:text" description:""`
	Scope        string `json:"scope"         gorm:"column:scope;type:text" description:""`
	CreatedBy    int    `json:"created_by"    gorm:"column:created_by" description:""`
	UpdatedBy    int    `json:"updated_by"    gorm:"column:updated_by" description:""`
}
