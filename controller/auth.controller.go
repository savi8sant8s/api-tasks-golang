package controller

import (
	"savi8sant8s/api/data"
	"savi8sant8s/api/entity"
	"savi8sant8s/api/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func (this *AuthController) DoRegister(c *gin.Context) {
	body := entity.User{}
	c.ShouldBindJSON(&body)
	c.JSON(this.authService.Register(body))
}

func (this *AuthController) DoLogin(c *gin.Context) {
	body := data.Login{}
	c.ShouldBindJSON(&body)
	c.JSON(this.authService.CreateSession(body))
}

func (this *AuthController) DoLogout(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	token := strings.Fields(authorizationHeader)[1]
	c.JSON(this.authService.CloseSession(token))
}
