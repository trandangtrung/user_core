package entity

type App struct {
	BaseEntity
	Name   string `json:"name" gorm:"type:varchar(255);not null;unique"`
	Config string `json:"config" gorm:"type:text;not null"`

	Users []User `gorm:"many2many:user_apps;"`
}
