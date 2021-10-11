package service

import (
	"net/http"
	"savi8sant8s/api/dao"
	"savi8sant8s/api/data"
	"savi8sant8s/api/entity"
	"savi8sant8s/api/utils"
	"savi8sant8s/api/validation"
)

type TaskService struct {
	sessionDao dao.SessionDao
	taskDao dao.TaskDao
}

func (this *TaskService) Create(userToken string, task entity.Task) (int, interface{}) {
	valid, messageError := validation.ValidTask(task)
	if !valid {
		return http.StatusBadRequest, data.Message {
			ApiStatus: utils.API_CREATE_TASK_INCORRECT_FIELDS, 
			Message: messageError,
		}
	} 
	task.UserID = this.sessionDao.UserID(userToken)
	this.taskDao.Create(task)
	return http.StatusOK, data.Message {
		ApiStatus: utils.API_CREATE_TASK_SUCCESS, 
		Message: utils.CREATE_TASK_SUCCESS,
	}
}

func (this *TaskService) All(userToken string) (int, interface{}) {
	userId := this.sessionDao.UserID(userToken)
	json := data.MessageTasks{
		Message: data.Message{
			ApiStatus: utils.API_ALL_TASKS_SUCCESS,
			Message: utils.ALL_TASKS_SUCCESS,
		},
		Tasks:   this.taskDao.FindAllById(userId),
	}
	return http.StatusOK, json
}

func (this *TaskService) Delete(userToken string, taskId string) (int, interface{}) {
	userId := this.sessionDao.UserID(userToken)
	this.taskDao.Delete(userId, taskId)
	return http.StatusOK, data.Message{
		ApiStatus: utils.API_DELETE_TASK_SUCCESS,
		Message: utils.DELETE_TASK_SUCCESS,
	}
}

func (this *TaskService) Update(userToken string, taskId string, task entity.Task) (int, interface{}) {
	valid, messageError := validation.ValidTask(task)
	if !valid {
		return http.StatusBadRequest, data.Message {
			ApiStatus: utils.API_UPDATE_TASK_INCORRECT_FIELDS, 
			Message: messageError,
		}
	} 
	userId := this.sessionDao.UserID(userToken)
	this.taskDao.Update(userId, taskId, task)
	return http.StatusOK, data.Message {
		ApiStatus: utils.API_UPDATE_TASK_SUCCESS, 
		Message: utils.UPDATE_TASK_SUCCESS,
	}
}