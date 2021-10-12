package route

func (this *App) PrepareTaskRoutes() {
	privateGroup := this.app.Group("/api/v1/tasks")
	privateGroup.Use(this.authMiddleware.Run())
	privateGroup.GET("", this.taskController.DoGet)
	privateGroup.POST("", this.taskController.DoCreate)
	privateGroup.Use(this.taskMiddleware.Run())
	privateGroup.PATCH("", this.taskController.DoUpdate)
	privateGroup.DELETE("", this.taskController.DoDelete)
}