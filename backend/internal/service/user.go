package service

import (
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() []model.User {
	return s.repo.GetAll()
}

func (s *UserService) GetByID(id int) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Create(user model.User) model.User {
	return s.repo.Create(user)
}
