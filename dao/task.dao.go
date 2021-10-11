package dao

import (
	"savi8sant8s/api/database"
	"savi8sant8s/api/entity"
)

type TaskDao struct {
	db database.Database
}

func (this *TaskDao) Create(task entity.Task) entity.Task {
	this.db.Instance().Select("UserID", "Title", "Message").Create(&task)
	return task
}

func (this *TaskDao) Exists(taskId string) bool {
	count := int64(0) 
	this.db.Instance().Take(&entity.Session{}).Where("id = ?", taskId).Count(&count)
	return count > 0
}

func (this *TaskDao) Update(userId uint, id string, task entity.Task) {
	this.db.Instance().Model(&entity.Task{}).Where("user_id = ? AND id = ?", userId, id).Update("title", task.Title).Update("message", task.Message)
}

func (this *TaskDao) Delete(userId uint, id string) {
	this.db.Instance().Where("user_id = ? AND id = ?", userId, id).Delete(&entity.Task{})
}

func (this *TaskDao) FindAllById(userId uint) []entity.Task{
	var tasks []entity.Task 
	this.db.Instance().Where("user_id = ? AND deleted_at IS NULL", userId).Find(&tasks)
	return tasks
}