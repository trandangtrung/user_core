package entity

type Token struct {
	BaseEntity
	RefreshToken string `json:"refresh_token" gorm:"type:text;not null"`
	Scope        string `json:"scope" gorm:"type:text;not null"`

	UserID uint `json:"user_id" gorm:"not null"`
}
