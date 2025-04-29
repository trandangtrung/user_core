package entity

type Role struct {
	BaseEntity
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Key         string `json:"key" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(255);not null"`
	AppID       uint   `json:"app_id" gorm:"not null"`
	App         *App   `json:"app" gorm:"foreignKey:AppID"`

	Users []User `gorm:"many2many:user_roles;"`
}
