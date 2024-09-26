package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	RefreshToken string    `json:"refresh_token" gorm:"not null"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	ExpiredAt    time.Time `json:"expired_at" gorm:"not null"`
}
