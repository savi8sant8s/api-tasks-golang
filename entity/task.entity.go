package entity

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID uint `json:"userId"`
	Title   string `json:"title"`
	Message string `json:"message"`
}