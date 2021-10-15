package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Tasks []Task `json:"tasks" gorm:" ForeignKey: UserID"`
}
