package route

import (
	"savi8sant8s/api/controller"
)

func (ur *AppRoute) PrepareAuthRoutes() {
	authController := new(controller.AuthController)

	route.POST("/api/v1/auth/register", authController.Register)
	route.POST("/api/v1/auth/login", authController.Login)
	route.POST("/api/v1/auth/logout", authController.Logout)
}