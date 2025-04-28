package entity

type UserRoles struct {
	BaseEntity
	UserID uint `json:"user_id" gorm:"not null"`
	RoleID uint `json:"role_id" gorm:"not null"`
}
