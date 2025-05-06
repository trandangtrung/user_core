package entity

import (
	"time"
)

type User struct {
	BaseEntity
	Email            string     `json:"email" gorm:"type:varchar(255);not null;unique"`
	Mobile           string     `json:"mobile" gorm:"type:varchar(20);unique"`
	PasswordHashed   string     `json:"password_hashed" gorm:"type:varchar(255);not null"`
	UserName         string     `json:"user_name" gorm:"type:varchar(50)"`
	FirstName        string     `json:"first_name" gorm:"type:varchar(50)"`
	MiddleName       string     `json:"middle_name" gorm:"type:varchar(50)"`
	LastName         string     `json:"last_name" gorm:"type:varchar(50)"`
	BirthDate        *time.Time `json:"birth_date" gorm:"type:date"`
	Language         string     `json:"language" gorm:"type:varchar(20)"`
	Currency         string     `json:"currency" gorm:"type:varchar(20)"`
	Address          string     `json:"address" gorm:"type:varchar(20)"`
	City             string     `json:"city" gorm:"type:varchar(20)"`
	State            string     `json:"state" gorm:"type:varchar(20)"`
	Country          string     `json:"country" gorm:"type:varchar(20)"`
	Bio              string     `json:"bio" gorm:"type:varchar(255)"`
	ProfilePicture   string     `json:"profile_picture" gorm:"type:varchar(255)"`
	IsEmailVerified  bool       `json:"email_verified" gorm:"not null;default:false"`
	IsMobileVerified bool       `json:"mobile_verified" gorm:"not null;default:false"`

	Tokens []Token `gorm:"foreignKey:UserID"`
	Apps   []App   `gorm:"many2many:user_apps;"`
	Roles  []Role  `gorm:"many2many:user_roles;"`
}
