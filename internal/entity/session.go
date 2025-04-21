package entity

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserId       uint   `json:"user_id" gorm:"column:user_id" description:""`
	User         *User  `json:"user"    gorm:"foreignKey:UserId;references:ID" description:"User"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token;type:text" description:""`
	Scope        string `json:"scope"         gorm:"column:scope;type:text" description:""`
	CreatedBy    int    `json:"create_by"     gorm:"column:create_by"      description:""`
	UpdatedBy    int    `json:"update_by"     gorm:"column:update_by"      description:""`
}
