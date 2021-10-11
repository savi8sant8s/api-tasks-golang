package dao

import (
	"savi8sant8s/api/database"
	"savi8sant8s/api/entity"
)

type UserDao struct {
	db database.Database
}

func (this *UserDao) New(user entity.User) entity.User {
	this.db.Instance().Select("Name", "Email", "Password").Create(&user)
	return user
}

func (this *UserDao) Exists(email string) bool {
	valid := this.db.Instance().Where("email = ?", email).Take(&entity.User{})
	return valid.RowsAffected > 0
}

func (this *UserDao) GetHashByEmail(email string) string {
	user := new(entity.User)
	this.db.Instance().Where("email = ?", email).Select("password").Find(&user)
	return user.Password
}

func (this *UserDao) GetIdByEmail(email string) uint {
	user := new(entity.User)
	this.db.Instance().Where("email = ?", email).Select("id").Find(&user)
	return user.ID
}