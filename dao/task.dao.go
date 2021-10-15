package dao

import (
	"savi8sant8s/gotasks/database"
	"savi8sant8s/gotasks/entity"
)

type TaskDao struct {
	db database.Database
}

func (this *TaskDao) Create(task entity.Task) {
	this.db.Instance().Select("user_id", "title", "message").Create(&task)
}

func (this *TaskDao) Exists(taskId uint) bool {
	count := int64(0)
	this.db.Instance().Take(&entity.Task{}).Where("id = ?", taskId).Count(&count)
	return count > 0
}

func (this *TaskDao) Update(task entity.Task) {
	this.db.Instance().Model(task).Where("id = ? AND user_id = ?", task.ID, task.UserID).Update("title", task.Title).Update("message", task.Message)
}

func (this *TaskDao) Delete(taskId uint) {
	this.db.Instance().Where("id = ? ", taskId).Delete(&entity.Task{})
}

func (this *TaskDao) FindAllById(userId uint) []entity.Task {
	var tasks []entity.Task
	this.db.Instance().Where("user_id = ? AND deleted_at IS NULL", userId).Find(&tasks)
	return tasks
}
