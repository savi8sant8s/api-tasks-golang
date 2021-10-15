package route

func (this *App) PrepareTaskRoutes() {
	privateGroup := this.app.Group("/api/v1/tasks")
	privateGroup.Use(this.authMiddleware.Run())
	privateGroup.GET("", this.taskController.DoGet)
	privateGroup.POST("", this.taskController.DoCreate)
	privateGroup.Use(this.taskMiddleware.Run())
	privateGroup.PATCH("/:taskId", this.taskController.DoUpdate)
	privateGroup.DELETE("/:taskId", this.taskController.DoDelete)
}