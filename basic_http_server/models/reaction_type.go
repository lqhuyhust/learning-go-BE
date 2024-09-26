package models

import (
	"gorm.io/gorm"
)

type ReactionType struct {
	gorm.Model
	Name string `json:"name" gorm:"unique; not null"`
}
