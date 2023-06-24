package cmd

import (
	_authHandler "alqinsidev/auth-service/modules/auth/handlers"
	_authService "alqinsidev/auth-service/modules/auth/services"
	_userHandlers "alqinsidev/auth-service/modules/user/handlers"
	_userRepository "alqinsidev/auth-service/modules/user/repositories"
	_userService "alqinsidev/auth-service/modules/user/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRoute := r.Group("/users")
	SetupUserHandlers(userRoute, db)

	authRoute := r.Group("/auth")
	SetupAuthHandlers(authRoute, db)

	return r
}

func SetupUserHandlers(userRoute *gin.RouterGroup, db *gorm.DB) {
	userRepository := _userRepository.NewUserRepository(db)
	userService := _userService.NewUserService(userRepository)
	_userHandlers.NewUserHandler(userRoute, userService)
}

func SetupAuthHandlers(authRoute *gin.RouterGroup, db *gorm.DB) {
	userRepository := _userRepository.NewUserRepository(db)
	authService := _authService.NewAuthService(userRepository)
	_authHandler.NewAuthHandler(authRoute, authService)
}
