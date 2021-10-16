package dao

import (
	"savi8sant8s/gotasks/database"
	"savi8sant8s/gotasks/entity"
)

type UserDao struct {
	db database.Database
}

func (this *UserDao) Create(user entity.User) entity.User {
	this.db.Instance().Select("name", "email", "password").Create(&user)
	return user
}

func (this *UserDao) Exists(email string) bool {
	valid := this.db.Instance().Where("email = ?", email).Take(&entity.User{})
	return valid.RowsAffected > 0
}

func (this *UserDao) FindUserByEmail(email string) entity.User {
	var user entity.User
	this.db.Instance().Where("email = ?", email).Find(&user)
	return user
}
