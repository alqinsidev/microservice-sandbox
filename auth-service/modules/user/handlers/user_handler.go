package handlers

import (
	"alqinsidev/auth-service/domain"
	"alqinsidev/auth-service/middlewares"
	"alqinsidev/auth-service/modules/user/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UsersHandler struct {
	userService *services.UserService
}

func NewUserHandler(r *gin.RouterGroup, userService *services.UserService) {
	handler := &UsersHandler{userService: userService}

	r.POST("/", middlewares.JWTMiddleware(), handler.CreateUser)
}

func (h *UsersHandler) CreateUser(c *gin.Context) {
	var createUserDTO domain.CreateUserDTO

	if err := c.ShouldBindJSON(&createUserDTO); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&createUserDTO); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.CreateUser(&createUserDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": result})
}
