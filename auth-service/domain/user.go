package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Nik      string    `json:"nik" gorm:"unique"`
	Password string    `json:"password"`
	Role     int       `json:"role"`
}

type CreateUserDTO struct {
	Nik  string `json:"nik" validate:"required"`
	Role int    `json:"role" validate:"required"`
}

type CreateUserResponse struct {
	Nik      string `json:"nik"`
	Role     int    `json:"role"`
	Password string `json:"password"`
}
