package service

import (
	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/jinzhu/copier"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]api.User, error) {
	modelUsers, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	apiUsers := []api.User{}
	if err := copier.Copy(&apiUsers, &modelUsers); err != nil {
		return nil, err
	}

	return apiUsers, nil
}

func (s *UserService) GetByID(id string) (*api.User, error) {
	modelUser, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	apiUser := api.User{}
	copier.Copy(&apiUser, &modelUser)

	return &apiUser, nil
}

func (s *UserService) Update(id string, u api.User) (*api.User, error) {
	modelUser := model.User{}
	copier.Copy(&modelUser, &u)
	_, err := s.repo.Update(id, modelUser)
	return &u, err
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}
