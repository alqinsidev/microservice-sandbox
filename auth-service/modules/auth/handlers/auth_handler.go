package handlers

import (
	"alqinsidev/auth-service/domain"
	"alqinsidev/auth-service/middlewares"
	"alqinsidev/auth-service/modules/auth/services"
	"alqinsidev/auth-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		errorData := utils.ErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorData)
		return
	}

	validate := validator.New()
	if err := validate.Struct(&loginDto); err != nil {
		errorData := utils.ErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorData)
		return
	}

	result, err := h.authService.Login(&loginDto)

	if err != nil {
		errorData := utils.ErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorData)
		return
	}

	response := utils.SuccessResponse(result, http.StatusOK, "success login")
	c.JSON(http.StatusOK, response)

}

func (h *AuthHandler) ValidateToken(c *gin.Context) {
	claims, _ := c.Get("user")
	claimsData := claims.(*domain.JWTClaims)
	result, err := h.authService.ValidateToken(claimsData) // Pass the pointer to claimsData
	if err != nil {
		errorData := utils.ErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorData)
		return
	}
	response := utils.SuccessResponse(result, http.StatusOK, "token are valid")
	c.JSON(http.StatusOK, response)
}
