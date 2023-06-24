package utils

import (
	"alqinsidev/auth-service/domain"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func GenerateToken(user *domain.User) (string, error) {
	secretKey := []byte(viper.GetString("JWT_ACCESS_SECRET_KEY"))
	AccessTokenTTL := time.Second * time.Duration(viper.GetInt("JWT_ACCESS_LIFETIME"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &domain.JWTClaims{
		ID:  user.ID,
		NIK: user.Nik,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenTTL).Unix(), // Token expires in 24 hours
		},
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	secretKey := []byte(viper.GetString("JWT_ACCESS_SECRET_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
