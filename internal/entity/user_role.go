package entity

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserId    uint  `json:"user_id" gorm:"column:user_id" description:""` // Thêm
	User      *User `json:"user"    gorm:"foreignKey:UserId;references:ID" description:"User"`
	RoleId    int   `json:"role_id" gorm:"column:role_id" description:""`
	CreatedBy int   `json:"created_by" gorm:"column:created_by" description:""`
	UpdatedBy int   `json:"updated_by" gorm:"column:updated_by" description:""`
}
