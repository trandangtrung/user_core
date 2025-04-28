package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	CreatedBy *uint          `json:"created_by"`
	UpdatedBy *uint          `json:"updated_by"`
}
