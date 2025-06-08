package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcuslinhares/auth-service/config"
	"github.com/marcuslinhares/auth-service/controllers"
	"github.com/marcuslinhares/auth-service/middleware"
	"github.com/marcuslinhares/auth-service/repositories"
	"github.com/marcuslinhares/auth-service/services"
)

func main() {
	config.Init()

	userRepo := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepo)

	controllers.SetAuthService(authService)

	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("/profile", controllers.Profile)
	}

	r.Run(":8080")
}
