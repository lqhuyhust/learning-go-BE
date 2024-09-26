package models

import (
	"gorm.io/gorm"
)

type Reaction struct {
	gorm.Model
	UserID         uint `gorm:"foreignKey:UserID"`
	PostID         uint `gorm:"foreignKey:PostID"`
	ReactionTypeID uint `gorm:"foreignKey:ReactionTypeID"`
}
