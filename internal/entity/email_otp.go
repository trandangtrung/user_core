package entity

import "time"

type EmailOTP struct {
	BaseEntity
	Email    string    `json:"email" gorm:"not null"`
	OTP      string    `json:"otp" gorm:"not null"`
	Used     bool      `json:"used" gorm:"not null;default:false"`
	ExpireAt time.Time `json:"expire_at" gorm:"not null"`
}
