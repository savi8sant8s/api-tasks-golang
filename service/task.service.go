package service

import (
	"net/http"
	"savi8sant8s/gotasks/dao"
	"savi8sant8s/gotasks/data"
	"savi8sant8s/gotasks/entity"
	"savi8sant8s/gotasks/utils"
	"savi8sant8s/gotasks/validation"
)

type TaskService struct {
	taskDao dao.TaskDao
}

func (this *TaskService) CreateTask(task entity.Task) (int, interface{}) {
	valid, messageError := validation.ValidTask(task)
	if !valid {
		return http.StatusBadRequest, data.Message{
			ApiStatus: utils.API_CREATE_TASK_INCORRECT_FIELDS,
			Message:   messageError,
		}
	}
	this.taskDao.Create(task)
	return http.StatusOK, data.Message{
		ApiStatus: utils.API_CREATE_TASK_SUCCESS,
		Message:   utils.CREATE_TASK_SUCCESS,
	}
}

func (this *TaskService) GetAllTasks(userId uint) (int, interface{}) {
	json := data.MessageTasks{
		Message: data.Message{
			ApiStatus: utils.API_ALL_TASKS_SUCCESS,
			Message:   utils.ALL_TASKS_SUCCESS,
		},
		Tasks: this.taskDao.FindAllById(userId),
	}
	return http.StatusOK, json
}

func (this *TaskService) DeleteTask(taskId uint) (int, interface{}) {
	this.taskDao.Delete(taskId)
	return http.StatusOK, data.Message{
		ApiStatus: utils.API_DELETE_TASK_SUCCESS,
		Message:   utils.DELETE_TASK_SUCCESS,
	}
}

func (this *TaskService) UpdateTask(task entity.Task) (int, interface{}) {
	valid, messageError := validation.ValidTask(task)
	if !valid {
		return http.StatusBadRequest, data.Message{
			ApiStatus: utils.API_UPDATE_TASK_INCORRECT_FIELDS,
			Message:   messageError,
		}
	}
	this.taskDao.Update(task)
	return http.StatusOK, data.Message{
		ApiStatus: utils.API_UPDATE_TASK_SUCCESS,
		Message:   utils.UPDATE_TASK_SUCCESS,
	}
}
