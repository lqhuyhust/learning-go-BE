package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content  string    `json:"content" gorm:"not null"`
	UserID   uint      `json:"user_id" gorm:"not null"`
	PostID   uint      `json:"post_id" gorm:"not null"`
	ParentID *uint     `json:"parent_id"`
	Replies  []Comment `gorm:"foreignKey:ParentID"`
}
