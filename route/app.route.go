package route

import "github.com/gin-gonic/gin"

type AppRoute struct {}

var route *gin.Engine = gin.Default()

func (br *AppRoute) PrepareRoutes(){
	br.PrepareAuthRoutes()
}

func (br *AppRoute) Run(){
	route.Run()
}