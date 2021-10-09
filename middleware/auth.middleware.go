package middleware

import (
	"net/http"
	"savi8sant8s/api/data"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func (c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		token := strings.Fields(authorizationHeader)
		if len(token) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, data.ErrorMessage{Erro: true, Message: "Token não especificado."})
		}
		if len(token[1]) < 30 {
			c.AbortWithStatusJSON(http.StatusBadRequest, data.ErrorMessage{Erro: true, Message: "Token não especificado corretamente."})
		} else {
			c.Next()
		}
    }
}