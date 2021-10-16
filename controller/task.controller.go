package controller

import (
	"savi8sant8s/gotasks/entity"
	"savi8sant8s/gotasks/service"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService service.TaskService
}

func (this *TaskController) DoCreate(c *gin.Context) {
	body := entity.Task{}
	c.ShouldBindJSON(&body)
	userId, _ := c.Get("userId")
	body.UserID = userId.(uint)
	this.taskService.CreateTask(c, body)
}

func (this *TaskController) DoGet(c *gin.Context) {
	userId, _ := c.Get("userId")
	this.taskService.GetAllTasks(c, userId.(uint))
}

func (this *TaskController) DoDelete(c *gin.Context) {
	taskId, _ := c.Get("taskId")
	this.taskService.DeleteTask(c, taskId.(uint))
}

func (this *TaskController) DoUpdate(c *gin.Context) {
	body := entity.Task{}
	c.ShouldBindJSON(&body)
	userId, _ := c.Get("userId")
	taskId, _ := c.Get("taskId")
	body.ID = taskId.(uint)
	body.UserID = userId.(uint)
	this.taskService.UpdateTask(c, body)
}
