package services

import (
	"alqinsidev/auth-service/domain"
	"alqinsidev/auth-service/modules/user/repositories"
	"alqinsidev/auth-service/utils"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	ur *repositories.UserRepository
}

func NewAuthService(ur *repositories.UserRepository) *AuthService {
	return &AuthService{ur: ur}
}

func ValidateHash(password string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, err
}

func (s *AuthService) Login(loginDTO *domain.LoginDTO) (*domain.LoginResponse, error) {

	matchUser, err := s.ur.GetByNIK(loginDTO.Nik)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	_, err = ValidateHash(loginDTO.Password, matchUser.Password)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("password didn't match")
	}

	accessToken, err := utils.GenerateToken(matchUser)
	if err != nil {
		return nil, err
	}

	res := &domain.LoginResponse{
		ID:    matchUser.ID,
		NIK:   matchUser.Nik,
		Role:  matchUser.Role,
		Token: accessToken,
	}

	return res, nil
}

func (s *AuthService) ValidateToken(claims *domain.JWTClaims) (*domain.ValidateTokenResponse, error) {

	expiredDateToken := time.Unix(claims.ExpiresAt, 0)
	res := &domain.ValidateTokenResponse{
		ID:          claims.ID,
		NIK:         claims.NIK,
		ExpiredDate: expiredDateToken,
	}
	return res, nil
}
