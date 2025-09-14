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

func (s *UserService) GetByID(id string) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Update(id string, u model.User) (*model.User, error) {
	return s.repo.Update(id, u)
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}
