package handlers

import (
	"alqinsidev/auth-service/domain"
	"alqinsidev/auth-service/modules/user/services"
	"alqinsidev/auth-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UsersHandler struct {
	userService *services.UserService
}

func NewUserHandler(r *gin.Engine, userService *services.UserService) {
	handler := &UsersHandler{userService: userService}

	r.POST("/register", handler.CreateUser)
}

func (h *UsersHandler) CreateUser(c *gin.Context) {
	var createUserDTO domain.CreateUserDTO

	if err := c.ShouldBindJSON(&createUserDTO); err != nil {
		errorData := utils.ErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorData)
		return
	}

	validate := validator.New()
	validate.RegisterValidation("validRole", domain.ValidateRole)
	if err := validate.Struct(&createUserDTO); err != nil {
		errorData := utils.ErrorResponse(http.StatusBadRequest, "plesae provide valid data")
		c.JSON(http.StatusBadRequest, errorData)
		return
	}

	result, err := h.userService.CreateUser(&createUserDTO)
	if err != nil {
		errorData := utils.ErrorResponse(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, errorData)
		return
	}

	response := utils.SuccessResponse(result, http.StatusCreated, "user has been created")
	c.JSON(http.StatusCreated, response)
}
