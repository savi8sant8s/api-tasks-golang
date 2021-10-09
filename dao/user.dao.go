package dao

import (
	"savi8sant8s/api/entity"
)

type UserDao struct {}

func (ud *UserDao) FindById(id uint) entity.User {
	user := new(entity.User)
	db.First(&user, id)
	return *user
}

func (ud *UserDao) Save(user entity.User) entity.User {
	db.Select("Name", "Email", "Password").Create(&user)
	return user
}