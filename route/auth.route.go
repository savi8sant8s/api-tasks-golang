package route

import (
	"savi8sant8s/api/controller"
	"savi8sant8s/api/middleware"
)

func (this *AppRoute) PrepareAuthRoutes() {
	controller := new(controller.AuthController)
	
	publicGroup := route.Group("/api/v1/auth")
	publicGroup.POST("/register", controller.DoRegister)
	publicGroup.POST("/login", controller.DoLogin)

	privateGroup := route.Group("/api/v1/auth")
	privateGroup.Use(new(middleware.AuthMiddleware).Run())
	privateGroup.POST("/logout", controller.DoLogout)
}
