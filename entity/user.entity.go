package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Sessions []Session `gorm:" ForeignKey: UserID"`
	Tasks []Task `gorm:" ForeignKey: UserID"`
}
