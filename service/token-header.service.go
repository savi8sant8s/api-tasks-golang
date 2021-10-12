package service

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTokenFromBearerAuthentication(c *gin.Context) string {
	authorizationHeader := c.Request.Header.Get("Authorization")
	return strings.Fields(authorizationHeader)[1]
}