package dao

import (
	"savi8sant8s/api/entity"
)

type UserDao struct {}

func (dao *UserDao) FindById(id uint) entity.User {
	user := new(entity.User)
	db.First(&user, id)
	return *user
}

func (dao *UserDao) New(user entity.User) entity.User {
	db.Select("Name", "Email", "Password").Create(&user)
	return user
}

func (dao *UserDao) Exists(email string) bool {
	valid := db.Where("email = ?", email).Take(&entity.User{})
	return valid.RowsAffected > 0
}

func (dao *UserDao) GetHashByEmail(email string) string {
	user := new(entity.User)
	db.Where("email = ?", email).Select("password").Find(&user)
	return user.Password
}

func (dao *UserDao) GetIdByEmail(email string) uint {
	user := new(entity.User)
	db.Where("email = ?", email).Select("id").Find(&user)
	return user.ID
}