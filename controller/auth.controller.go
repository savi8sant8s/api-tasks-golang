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
	this.authService.RegisterUser(c, body)
}

func (this *AuthController) DoLogin(c *gin.Context) {
	body := entity.User{}
	c.ShouldBindJSON(&body)
	this.authService.CreateSession(c, body)
}
