package service

import (
	"net/http"
	"savi8sant8s/gotasks/dao"
	"savi8sant8s/gotasks/data"
	"savi8sant8s/gotasks/entity"
	"savi8sant8s/gotasks/utils"
	"savi8sant8s/gotasks/validation"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskService struct {
	taskDao dao.TaskDao
}

func (this *TaskService) CreateTask(c *gin.Context, task entity.Task) {
	valid := this.ValidTaskBody(c, task)
	if valid {
		this.taskDao.Create(task)
		c.JSON(http.StatusOK, data.Message{
			ApiStatus: utils.API_CREATE_TASK_SUCCESS,
			Message:   utils.CREATE_TASK_SUCCESS,
		})
	}
}

func (this *TaskService) GetAllTasks(c *gin.Context, userId uint) {
	c.JSON(http.StatusOK, data.MessageTasks{
		Message: data.Message{
			ApiStatus: utils.API_ALL_TASKS_SUCCESS,
			Message:   utils.ALL_TASKS_SUCCESS,
		},
		Tasks: this.taskDao.FindAllById(userId),
	})
}

func (this *TaskService) DeleteTask(c *gin.Context, taskId uint) {
	this.taskDao.Delete(taskId)
	c.JSON(http.StatusOK, data.Message{
		ApiStatus: utils.API_DELETE_TASK_SUCCESS,
		Message:   utils.DELETE_TASK_SUCCESS,
	})
}

func (this *TaskService) UpdateTask(c *gin.Context, task entity.Task) {
	valid := this.ValidTaskBody(c, task)
	if valid {
		this.taskDao.Update(task)
		c.JSON(http.StatusOK, data.Message{
			ApiStatus: utils.API_UPDATE_TASK_SUCCESS,
			Message:   utils.UPDATE_TASK_SUCCESS,
		})
	}	
}

func (this *TaskService) ValidTaskBody(c *gin.Context, task entity.Task) bool {
	valid, messageError := validation.ValidTask(task)
	if !valid {
		c.JSON(http.StatusBadRequest, data.Message{
			ApiStatus: utils.API_TASK_INCORRECT_FIELDS,
			Message:   messageError,
		})
		return false
	}
	return true
}

func (this *TaskService) TaskExists(c *gin.Context) bool {
	taskId, _ := strconv.ParseUint(c.Param("taskId"), 10, 64)
	exists := this.taskDao.Exists(uint(taskId))
	c.Set("taskId", uint(taskId))
	if exists {
		return true
	} 
	c.JSON(http.StatusNotFound, data.Message{
		ApiStatus: utils.API_NOT_FOUND_TASK,
		Message:   utils.NOT_FOUND_TASK_SUCCESS,
	})
	return false
}