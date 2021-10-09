package controller

import (
	"net/http"
	"savi8sant8s/api/data"
	"savi8sant8s/api/entity"
	"savi8sant8s/api/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service service.AuthService
 }

func (controller *AuthController) Register(c *gin.Context) {
	body := entity.User{}
	c.ShouldBindJSON(&body)
	c.JSON(controller.service.RegisterUser(body))
}

func (controller *AuthController) Login(c *gin.Context) {
	body := data.Login{}
	c.ShouldBindJSON(&body)
	c.JSON(controller.service.Login(body))
}

func (controller *AuthController) Logout(c *gin.Context) {
	bearerToken := c.Request.Header.Get("UserId")
	c.Request.Header.Del("UserId")
	c.JSON(http.StatusOK, bearerToken)
}