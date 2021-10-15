package route

import (
	"io"
	"os"
	"savi8sant8s/api/controller"
	"savi8sant8s/api/middleware"
	"github.com/gin-gonic/gin"
)

type App struct {
	app *gin.Engine
	authController controller.AuthController
	taskController controller.TaskController
	authMiddleware middleware.AuthMiddleware
	taskMiddleware middleware.TaskMiddleware
}

func (this *App) ConfigApp() {
	gin.SetMode(gin.DebugMode)
	this.app = gin.New()
}

func (this *App) ConfigLogger() {
	f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)

	this.app.Use(gin.Logger())
	this.app.Use(gin.Recovery())
}

func (this *App) Prepare(){	
	this.ConfigApp()
	this.ConfigLogger()
	
	this.PrepareAuthRoutes()
	this.PrepareTaskRoutes()
}

func (this *App) Run(){
	this.app.Run(":8080")
}