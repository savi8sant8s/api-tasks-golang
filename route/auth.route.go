package route

func (this *App) PrepareAuthRoutes() {
	publicGroup := this.app.Group("/api/v1/auth")
	publicGroup.POST("/register", this.authController.DoRegister)
	publicGroup.POST("/login", this.authController.DoLogin)

	privateGroup := this.app.Group("/api/v1/auth")
	privateGroup.Use(this.authMiddleware.Run())
}
