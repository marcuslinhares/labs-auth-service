package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcuslinhares/auth-service/services"
)

var authService *services.AuthService

func SetAuthService(svc *services.AuthService) {
	authService = svc
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Register godoc
// @Summary Cadastra um novo usuário
// @Description Registra um usuário com nome, email e senha
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Dados do usuário"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	err := authService.Register(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

// Login godoc
// @Summary Realiza login de um usuário
// @Description Autentica usuário e retorna um token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Credenciais do usuário"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	token, err := authService.Login(ctx, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Profile godoc
// @Summary Retorna o perfil do usuário autenticado
// @Description Requer um token JWT válido no header Authorization
// @Tags Auth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /api/profile [get]
// @Security ApiKeyAuth
func Profile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": userID})
}
