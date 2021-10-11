package controller

import (
	"savi8sant8s/api/entity"
	"savi8sant8s/api/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService service.TaskService
}

func (this *TaskController) Create(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	token := strings.Fields(authorizationHeader)[1]
	body := entity.Task{}
	c.ShouldBindJSON(&body)
	c.JSON(this.taskService.Create(token, body))
}

func (this *TaskController) All(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	token := strings.Fields(authorizationHeader)[1]
	c.JSON(this.taskService.All(token))
}

func (this *TaskController) Delete(c *gin.Context) {
	taskId := c.Query("id")
	authorizationHeader := c.Request.Header.Get("Authorization")
	token := strings.Fields(authorizationHeader)[1]
	c.JSON(this.taskService.Delete(token, taskId))
}

func (this *TaskController) Update(c *gin.Context) {
	taskId := c.Query("taskid")
	authorizationHeader := c.Request.Header.Get("Authorization")
	token := strings.Fields(authorizationHeader)[1]
	body := entity.Task{}
	c.ShouldBindJSON(&body)
	c.JSON(this.taskService.Update(token, taskId, body))
}