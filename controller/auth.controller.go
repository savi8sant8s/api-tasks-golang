package controller

import (
	"net/http"
	"savi8sant8s/api/data"

	"github.com/gin-gonic/gin"
)

type AuthController struct { }

func (ac *AuthController) Register(c *gin.Context) {
	body := data.Register{}
	c.ShouldBindJSON(&body)
	c.JSON(http.StatusOK, body)
}

func (ac *AuthController) Login(c *gin.Context) {
	body := data.Login{}
	c.ShouldBindJSON(&body)
	c.JSON(http.StatusOK, body)
}

func (ac *AuthController) Logout(c *gin.Context) {
	bearerToken := c.Request.Header.Get("Authorization")
	c.JSON(http.StatusOK, bearerToken)
}