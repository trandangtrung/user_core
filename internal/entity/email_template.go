package entity

type EmailTemplate struct {
	BaseEntity
	Name        string `json:"name" gorm:"type:varchar(255);not null;unique"`
	Key         string `json:"key" gorm:"type:varchar(255);not null;unique"`
	Subject     string `json:"subject" gorm:"type:varchar(255);not null"`
	Body        string `json:"body" gorm:"type:text;not null"`
	Description string `json:"description" gorm:"type:text"`
	IsActive    bool   `json:"is_active" gorm:"type:boolean;default:true"`
	AppID       uint   `json:"app_id" gorm:"not null"`
	App         App    `json:"app" gorm:"foreignKey:AppID"`
}
