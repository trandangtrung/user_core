package entity

type UserApps struct {
	BaseEntity
	UserID uint `json:"user_id" gorm:"not null"`
	AppID  uint `json:"app_id" gorm:"not null"`
}
