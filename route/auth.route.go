package route

func (this *AppRoute) PrepareAuthRoutes() {
	publicGroup := this.route.Group("/api/v1/auth")
	publicGroup.POST("/register", this.authController.DoRegister)
	publicGroup.POST("/login", this.authController.DoLogin)

	privateGroup := this.route.Group("/api/v1/auth")
	privateGroup.Use(this.authMiddleware.Run())
	privateGroup.POST("/logout", this.authController.DoLogout)
}
