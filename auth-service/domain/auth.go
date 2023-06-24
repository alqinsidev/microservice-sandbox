package domain

import (
	"time"

	"github.com/google/uuid"
)

type LoginDTO struct {
	Nik      string `json:"nik" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	ID    uuid.UUID `json:"id"`
	NIK   string    `json:"nik"`
	Role  int       `json:"role"`
	Token string    `json:"access_token"`
}

type ValidateTokenResponse struct {
	ID          uuid.UUID `json:"id"`
	NIK         string    `json:"nik"`
	ExpiredDate time.Time `json:"expired_date_token"`
}
