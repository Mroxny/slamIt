package service

import (
	"errors"
	"time"

	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(name, email, password string) (*model.User, error) {
	if u, _ := s.userRepo.GetByEmail(email); u != nil {
		return nil, errors.New("user with email already exists")
	}

	hash, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		ID:           uuid.New().String(),
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
	}

	u, err := s.userRepo.Create(&user)
	if err != nil {
		return nil, err
	}

	return u, nil
}

type LoginResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (s *AuthService) Login(email, password string) (*LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials (no email)")
	}

	if !PasswordHashMatch(user.PasswordHash, password) {
		return nil, errors.New("invalid credentials (wrong password)")
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		return nil, errors.New("error when creating the auth token")
	}

	res := &LoginResponse{
		ID:    user.ID,
		Token: token,
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

var jwtKey = []byte("supersecretkey") // move to env

func GenerateJWT(userID string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)
	return claims.Subject, nil
}
