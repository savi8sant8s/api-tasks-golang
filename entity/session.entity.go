package entity

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserID uint
	Token   string
	Expired bool
}

