package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcuslinhares/auth-service/config"
	"github.com/marcuslinhares/auth-service/controllers"
	_ "github.com/marcuslinhares/auth-service/docs"
	"github.com/marcuslinhares/auth-service/middleware"
	"github.com/marcuslinhares/auth-service/repositories"
	"github.com/marcuslinhares/auth-service/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth Service API
// @version 1.0
// @description Documentação da API de autenticação
// @host auth-service-byz1bp-c95241-204-12-199-113.traefik.me
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
		auth.GET("/verify", middleware.JWTAuthMiddleware(), controllers.Profile)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
