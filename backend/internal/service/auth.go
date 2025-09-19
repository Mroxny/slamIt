package service

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/Mroxny/slamIt/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(name, email, password string) (*api.User, error) {
	if u, _ := s.userRepo.GetByEmail(email); u != nil {
		return nil, errors.New("user with email already exists")
	}

	// hash, err := HashPassword(password)
	// if err != nil {
	// 	return nil, err
	// }

	newId := uuid.New().String()
	user := api.User{
		Id:    &newId,
		Name:  &name,
		Email: &email,
	}

	u, err := s.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *AuthService) Login(email, password string) (*api.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials (no email)")
	}

	// if !PasswordHashMatch(user.PasswordHash, password) {
	// 	return nil, errors.New("invalid credentials (wrong password)")
	// }

	token, err := utils.GenerateJWT(*user.Id)
	if err != nil {
		return nil, errors.New("error when creating the auth token")
	}

	res := &api.LoginResponse{
		UserId: user.Id,
		Token:  &token,
	}

	return res, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func PasswordHashMatch(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
