package middleware

import (
	"net/http"
	"savi8sant8s/gotasks/dao"
	"savi8sant8s/gotasks/data"
	"savi8sant8s/gotasks/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskMiddleware struct {
	taskDao dao.TaskDao
}

func (this *TaskMiddleware) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskId, _ := strconv.ParseUint(c.Param("taskId"), 10, 64)
		exists := this.taskDao.Exists(uint(taskId))
		c.Set("taskId", uint(taskId))
		if !exists {
			c.AbortWithStatusJSON(http.StatusNotFound, data.Message{
				ApiStatus: utils.API_NOT_FOUND_TASK,
				Message:   utils.NOT_FOUND_TASK_SUCCESS,
			})
		} else {
			c.Next()
		}
	}
}
