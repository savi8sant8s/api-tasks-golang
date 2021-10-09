package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string 
	Email string
	Password string
	Sessions []Session `gorm:" ForeignKey: UserID"`
	Tasks []Task `gorm:" ForeignKey: UserID"`
}
