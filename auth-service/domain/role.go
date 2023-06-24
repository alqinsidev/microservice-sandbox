package domain

import "github.com/go-playground/validator/v10"

const (
	ROLE_SUPER_ADMIN int64 = 1
	ROLE_USER        int64 = 2
)

func ValidateRole(fl validator.FieldLevel) bool {
	role := fl.Field().Int()
	return role == ROLE_SUPER_ADMIN || role == ROLE_USER
}
