package controller

import (
	"savi8sant8s/api/entity"
	"savi8sant8s/api/service"
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
	c.JSON(this.taskService.CreateTask(body))
}

func (this *TaskController) DoGet(c *gin.Context) {
	userId, _ := c.Get("userId")
	c.JSON(this.taskService.GetAllTasks(userId.(uint)))
}

func (this *TaskController) DoDelete(c *gin.Context) {
	taskId, _ := c.Get("taskId")
	c.JSON(this.taskService.DeleteTask(taskId.(uint)))
}

func (this *TaskController) DoUpdate(c *gin.Context) {
	body := entity.Task{}
	c.ShouldBindJSON(&body)
	userId, _ := c.Get("userId")
	taskId, _ := c.Get("taskId")
	body.ID = taskId.(uint)
	body.UserID = userId.(uint)
	c.JSON(this.taskService.UpdateTask(body))
}