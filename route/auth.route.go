package route

import (
	"savi8sant8s/api/controller"
	"savi8sant8s/api/middleware"
)

func (ur *AppRoute) PrepareAuthRoutes() {
	authController := new(controller.AuthController)
	
	publicGroup := route.Group("/api/v1/auth")
	publicGroup.POST("/register", authController.Register)
	publicGroup.POST("/login", authController.Login)

	privateGroup := route.Group("/api/v1/auth")
	privateGroup.Use(new(middleware.AuthMiddleware).Run())
	privateGroup.POST("/logout", authController.Logout)
}
