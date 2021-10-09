package route

import (
	"github.com/gin-gonic/gin"
)

type AppRoute struct {}

var route *gin.Engine = gin.New()

func (appRoute *AppRoute) PrepareRoutes(){
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	
	appRoute.PrepareAuthRoutes()
	appRoute.PrepareTaskRoutes()
}

func (appRoute *AppRoute) Run(){
	route.Run(":8080")
}