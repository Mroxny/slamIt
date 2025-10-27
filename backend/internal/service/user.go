package service

import (
	"context"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAll(ctx context.Context, page, pageSize int) (*api.UserPagination, error) {
	offset := (page - 1) * pageSize
	users, err := s.userRepo.FindAllPaginated(ctx, pageSize, offset)
	if err != nil {
		return nil, err
	}
	var apiUsers []api.User

	if err = copier.Copy(&apiUsers, &users); err != nil {
		return nil, err
	}

	pag := api.UserPagination{
		Page:     &page,
		PageSize: &pageSize,
		Items:    &apiUsers,
	}
	return &pag, nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (*api.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	var apiUser api.User

	if err = copier.Copy(&apiUser, &user); err != nil {
		return nil, err
	}
	return &apiUser, nil
}

func (s *UserService) CreateTmpUser(ctx context.Context, u api.UserRequest) (*api.User, error) {
	modelUser := model.User{}
	copier.Copy(&modelUser, &u)
	modelUser.Id = uuid.New().String()
	val := true
	modelUser.TmpUser = &val

	if err := s.userRepo.Create(ctx, &modelUser); err != nil {
		return nil, err
	}

	apiUser := api.User{}
	copier.Copy(&apiUser, &modelUser)
	return &apiUser, nil
}

func (s *UserService) FindUserByEmail(ctx context.Context, email string) (*api.User, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	var apiUser api.User

	if err = copier.Copy(&apiUser, &user); err != nil {
		return nil, err
	}
	return &apiUser, nil
}

func (s *UserService) Update(ctx context.Context, id string, u api.UserRequest) (*api.User, error) {
	modelUser := model.User{}
	copier.Copy(&modelUser, &u)
	modelUser.Id = id

	if err := s.userRepo.Update(ctx, &modelUser); err != nil {
		return nil, err
	}

	apiUser := api.User{}
	copier.Copy(&apiUser, &modelUser)
	return &apiUser, nil
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.userRepo.Delete(ctx, id)
}
