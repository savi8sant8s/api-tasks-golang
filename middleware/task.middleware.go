package middleware

import (
	"savi8sant8s/gotasks/service"

	"github.com/gin-gonic/gin"
)

type TaskMiddleware struct {
	taskService service.TaskService
}

func (this *TaskMiddleware) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		if this.taskService.TaskExists(c) {
			c.Next()
		}
	}
}
