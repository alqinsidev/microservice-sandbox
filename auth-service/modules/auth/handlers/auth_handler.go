package handlers

import (
	"alqinsidev/auth-service/domain"
	"alqinsidev/auth-service/middlewares"
	"alqinsidev/auth-service/modules/auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(r *gin.RouterGroup, authService *services.AuthService) {
	handler := &AuthHandler{authService: authService}

	r.POST("/login", handler.Login)
	r.GET("/validate", middlewares.JWTMiddleware(), handler.ValidateToken)

}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginDto domain.LoginDTO

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&loginDto); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.authService.Login(&loginDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})

}

func (h *AuthHandler) ValidateToken(c *gin.Context) {
	claims, _ := c.Get("user")
	claimsData := claims.(*domain.JWTClaims)
	result, err := h.authService.ValidateToken(claimsData) // Pass the pointer to claimsData
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}
