package controller

import (
	"net/http"
	"savi8sant8s/api/data"

	"github.com/gin-gonic/gin"
)

type TaskController struct { }

func (ac *TaskController) New(c *gin.Context) {
	body := data.Task{}
	c.ShouldBindJSON(&body)
	c.JSON(http.StatusOK, body)
}

func (ac *TaskController) Delete(c *gin.Context) {
	taskid := c.Query("taskid")
	c.JSON(http.StatusOK, taskid)
}

func (ac *TaskController) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, "Todas as tarefas")
}

func (ac *TaskController) UpdateMessage(c *gin.Context) {
	taskid := c.Query("taskid")
	body := data.Task{}
	c.ShouldBindJSON(&body)
	body.Title = body.Title + taskid
	c.JSON(http.StatusOK, body)
}