package controller

import (
	"savi8sant8s/gotasks/entity"
	"savi8sant8s/gotasks/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func (this *AuthController) DoRegister(c *gin.Context) {
	body := entity.User{}
	c.ShouldBindJSON(&body)
	c.JSON(this.authService.RegisterUser(body))
}

func (this *AuthController) DoLogin(c *gin.Context) {
	body := entity.User{}
	c.ShouldBindJSON(&body)
	c.JSON(this.authService.CreateSession(body))
}
