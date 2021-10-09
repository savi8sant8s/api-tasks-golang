package entity

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID uint
	Title   string
	Message string
}