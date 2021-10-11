package middleware

import (
	"net/http"
	"savi8sant8s/api/dao"
	"savi8sant8s/api/data"
	"savi8sant8s/api/utils"

	"github.com/gin-gonic/gin"
)

type TaskMiddleware struct {
	taskDao dao.TaskDao
}

func (this *TaskMiddleware) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
			taskId := c.Query("id")
			exists := this.taskDao.Exists(taskId)
			if !exists {
				c.AbortWithStatusJSON(http.StatusNotFound, data.Message{
					ApiStatus: utils.API_NOT_FOUND_TASK,
					Message: utils.NOT_FOUND_TASK_SUCCESS,
				})
			} else {
				c.Next()
			}
		}
}
