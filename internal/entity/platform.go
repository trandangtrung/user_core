package entity

import "gorm.io/gorm"

type Platform struct {
	gorm.Model
	Name      string `json:"name" gorm:"column:name;type:varchar(255);not null;unique" description:""`
	Config    string `json:"config" gorm:"column:config;type:text;not null;" description:""`
	CreatedBy int    `json:"created_by" gorm:"column:created_by;not null;" description:""`
	UpdatedBy int    `json:"updated_by" gorm:"column:updated_by;not null;" description:""`
}
