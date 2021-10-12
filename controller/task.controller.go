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
	token := service.GetTokenFromBearerAuthentication(c)
	body := entity.Task{}
	c.ShouldBindJSON(&body)
	c.JSON(this.taskService.Create(token, body))
}

func (this *TaskController) DoGet(c *gin.Context) {
	token := service.GetTokenFromBearerAuthentication(c)
	c.JSON(this.taskService.GetAll(token))
}

func (this *TaskController) DoDelete(c *gin.Context) {
	taskId := c.Query("id")
	token := service.GetTokenFromBearerAuthentication(c)
	c.JSON(this.taskService.Delete(token, taskId))
}

func (this *TaskController) DoUpdate(c *gin.Context) {
	taskId := c.Query("id")
	token := service.GetTokenFromBearerAuthentication(c)
	body := entity.Task{}
	c.ShouldBindJSON(&body)
	c.JSON(this.taskService.Update(token, taskId, body))
}