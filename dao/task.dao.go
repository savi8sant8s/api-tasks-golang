package dao

import (
	"savi8sant8s/api/entity"
)

type TaskDao struct {}

func (td *TaskDao) Create(task entity.Task) entity.Task {
	db.Select("UserID", "Title", "Message").Create(&task)
	return task
}

func (td *TaskDao) UpdateMessage(userid uint, message string) {
	db.Model(&entity.Task{}).Where("user_id = ?", userid).Update("message", message)
}

func (td *TaskDao) Delete(id uint) {
	db.Delete(&entity.Task{}, id)
}

func (td *TaskDao) FindAllById(id uint) []entity.Task{
	var user = new(entity.User)
	db.First(&user, 10)
	return user.Tasks
}