package cmd

import (
	_authHandler "alqinsidev/auth-service/modules/auth/handlers"
	_authService "alqinsidev/auth-service/modules/auth/services"
	_userHandlers "alqinsidev/auth-service/modules/user/handlers"
	_userRepository "alqinsidev/auth-service/modules/user/repositories"
	_userService "alqinsidev/auth-service/modules/user/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	SetupUserHandlers(r, db)

	SetupAuthHandlers(r, db)

	return r
}

func SetupUserHandlers(r *gin.Engine, db *gorm.DB) {
	userRepository := _userRepository.NewUserRepository(db)
	userService := _userService.NewUserService(userRepository)
	_userHandlers.NewUserHandler(r, userService)
}

func SetupAuthHandlers(r *gin.Engine, db *gorm.DB) {
	userRepository := _userRepository.NewUserRepository(db)
	authService := _authService.NewAuthService(userRepository)
	_authHandler.NewAuthHandler(r, authService)
}
