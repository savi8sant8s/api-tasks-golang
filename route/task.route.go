package route

import (
	"savi8sant8s/api/controller"
	"savi8sant8s/api/middleware"
)

func (ur *AppRoute) PrepareTaskRoutes() {
	taskController := new(controller.TaskController)

	privateGroup := route.Group("/api/v1/tasks")
	privateGroup.Use(new(middleware.AuthMiddleware).Run())
	privateGroup.GET("", taskController.GetAll)
	privateGroup.POST("", taskController.New)
	privateGroup.PATCH("", taskController.Update)
	privateGroup.DELETE("", taskController.Delete)
}