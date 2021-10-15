package service

import (
	"strings"
	"github.com/gin-gonic/gin"
)

func GetTokenFromBearerAuth(c *gin.Context) (string, bool) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer"){
		return "", true
	}
	return strings.Fields(authorizationHeader)[1], false
}