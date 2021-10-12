package middleware

import (
	"net/http"
	"savi8sant8s/api/dao"
	"savi8sant8s/api/data"
	"savi8sant8s/api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	sessionDao dao.SessionDao
}

func (this *AuthMiddleware) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusBadRequest, data.Message {
				ApiStatus: utils.API_INCORRECT_AUTH_HEADER, 
				Message: utils.ERROR_INCORRECT_AUTH_HEADER,
			})
		} else {
			token := strings.Fields(authorizationHeader)[1]
			valid := this.sessionDao.ValidToken(token)
			if !valid {
				c.AbortWithStatusJSON(http.StatusBadRequest, data.Message {
					ApiStatus: utils.API_INVALID_TOKEN, 
					Message: utils.ERROR_INVALID_TOKEN,
				})
			}
			c.Next()
		}
	}
}
