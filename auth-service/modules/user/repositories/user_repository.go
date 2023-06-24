package repositories

import (
	"alqinsidev/auth-service/domain"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(u *domain.User) (*domain.User, error) {

	if err := ur.db.Create(&u).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("nik already exist")
		} else {
			return nil, err
		}
	}
	return u, nil
}

func (ur *UserRepository) GetByNIK(nik string) (*domain.User, error) {
	var user *domain.User
	if err := ur.db.Where("nik = ?", nik).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
