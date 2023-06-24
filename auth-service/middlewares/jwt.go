package middlewares

import (
	"alqinsidev/auth-service/domain"
	"alqinsidev/auth-service/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		claims := &domain.JWTClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			secret := viper.GetString("JWT_ACCESS_SECRET_KEY")
			accessSecretKey := []byte(secret)
			return accessSecretKey, nil
		})
		if err != nil {
			errorData := utils.ErrorResponse(http.StatusUnauthorized, "unauthorized")
			c.JSON(http.StatusUnauthorized, errorData)
			c.Abort()
			return
		}

		if !token.Valid {
			errorData := utils.ErrorResponse(http.StatusUnauthorized, "unauthorized")
			c.JSON(http.StatusUnauthorized, errorData)
			c.Abort()
		}

		c.Set("user", claims)
		c.Next()
	}
}
