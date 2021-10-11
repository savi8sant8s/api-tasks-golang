package route

func (this *AppRoute) PrepareTaskRoutes() {
	privateGroup := this.route.Group("/api/v1/tasks")
	privateGroup.Use(this.authMiddleware.Run())
	privateGroup.GET("", this.taskController.All)
	privateGroup.POST("", this.taskController.Create)
	privateGroup.Use(this.taskMiddleware.Run())
	privateGroup.PATCH("", this.taskController.Update)
	privateGroup.DELETE("", this.taskController.Delete)
}