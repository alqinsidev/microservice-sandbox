package services

import (
	"alqinsidev/auth-service/domain"
	"alqinsidev/auth-service/modules/user/repositories"
	"alqinsidev/auth-service/utils"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ur *repositories.UserRepository
}

func GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidateHash(password string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, err
}

func GenerateRandomPassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	chars := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, length)

	result[0] = chars[rand.Intn(10)]
	result[1] = chars[10+rand.Intn(len(chars)-10)]

	for i := 2; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	for i := range result {
		j := rand.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func NewUserService(ur *repositories.UserRepository) *UserService {
	return &UserService{ur}
}

func (s *UserService) CreateUser(createUserDTO *domain.CreateUserDTO) (*domain.CreateUserResponse, error) {
	randomPassword := GenerateRandomPassword(6)
	logrus.Info(randomPassword)
	hashedPassword, err := GenerateHash(randomPassword)
	if err != nil {
		fmt.Println(err.Error())
	}

	user := &domain.User{
		Nik:      createUserDTO.Nik,
		Password: hashedPassword,
		Role:     createUserDTO.Role,
	}

	storedUser, err := s.ur.Create(user)
	if err != nil {
		return nil, err
	}

	createdUser := &domain.CreateUserResponse{
		Nik:      storedUser.Nik,
		Password: randomPassword,
		Role:     storedUser.Role,
	}
	return createdUser, nil
}

func (s *UserService) Login(loginDTO *domain.LoginDTO) (*domain.LoginResponse, error) {

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
