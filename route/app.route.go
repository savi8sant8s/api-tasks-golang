package route

import (
	"github.com/gin-gonic/gin"
	"savi8sant8s/api/controller"
	"savi8sant8s/api/middleware"
)

type AppRoute struct {
	route *gin.Engine
	authController controller.AuthController
	taskController controller.TaskController
	authMiddleware middleware.AuthMiddleware
	taskMiddleware middleware.TaskMiddleware
}

func (this *AppRoute) PrepareRoutes(){
	this.route = gin.New()
	this.route.Use(gin.Logger())
	this.route.Use(gin.Recovery())
	
	this.PrepareAuthRoutes()
	this.PrepareTaskRoutes()
}

func (this *AppRoute) Run(){
	this.route.Run(":8080")
}