package middleware

import (
	"fmt"
	"net/http"
	"savi8sant8s/api/service"
	"savi8sant8s/api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	service service.AuthService
}

func (this *AuthMiddleware) Run() gin.HandlerFunc {
    return func (c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			c.AbortWithStatusJSON(service.MakeRes(http.StatusBadRequest, service.API_INCORRECT_AUTH_HEADER, utils.MESSAGE_ERROR_INCORRECT_AUTH_HEADER))
		} else {
			token := strings.Fields(authorizationHeader)[1]
			valid, userId := this.service.ValidToken(token)
			if !valid {
				c.AbortWithStatusJSON(service.MakeRes(http.StatusBadRequest, service.API_INVALID_TOKEN, utils.MESSAGE_ERROR_INVALID_TOKEN))
			} 
			c.Request.Header.Add("UserId", fmt.Sprintf("%v",userId))
			c.Next()
		}
		
    }
}