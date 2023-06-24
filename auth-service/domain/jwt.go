package domain

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTClaims struct {
	ID  uuid.UUID `json:"id"`
	NIK string    `json:"nik"`
	jwt.StandardClaims
}
