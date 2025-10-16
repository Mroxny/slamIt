package service

import (
	"context"
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ParticipationService struct {
	partRepo *repository.ParticipationRepository
}

func NewParticipationService(partRepo *repository.ParticipationRepository) *ParticipationService {
	return &ParticipationService{
		partRepo: partRepo,
	}
}

func (s *ParticipationService) AddUserToSlam(ctx context.Context, userID, slamID string, role api.ParticipationRoleEnum) (*api.Participation, error) {
	_, err := s.partRepo.FindBySlamAndUser(ctx, slamID, userID)
	if err == nil {
		return nil, errors.New("user is already participating in this slam")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	participation := model.Participation{

		Participation: api.Participation{
			Id:     uuid.New().String(),
			Role:   role,
			UserId: userID,
			SlamId: slamID,
		},
	}

	if err := s.partRepo.Create(ctx, &participation); err != nil {
		return nil, errors.New("failed to add user to slam, check if user and slam exist")
	}

	apiPart := api.Participation{}
	copier.Copy(&apiPart, &participation.Participation)
	return &apiPart, nil
}

func (s *ParticipationService) RemoveUserFromSlam(ctx context.Context, userID, slamID string) error {
	err := s.partRepo.DeleteBySlamAndUser(ctx, slamID, userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user is not participating in this slam")
	}
	return err
}

func (s *ParticipationService) GetUsersForSlam(ctx context.Context, slamID string) ([]api.Participation, error) {
	modelUsers, err := s.partRepo.FindParticipatingUsersBySlamID(ctx, slamID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("slam not found")
	}
	if err != nil {
		return nil, err
	}

	var apiUsers []api.Participation
	copier.Copy(&apiUsers, &modelUsers)
	return apiUsers, nil
}

func (s *ParticipationService) GetSlamsForUser(ctx context.Context, userID string) ([]api.Participation, error) {
	modelSlams, err := s.partRepo.FindParticipatedSlamsByUserID(ctx, userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	var apiSlams []api.Participation
	copier.Copy(&apiSlams, &modelSlams)
	return apiSlams, nil
}

func (s *ParticipationService) UpdateParticipation(ctx context.Context, slamID, userID string, req api.ParticipationUpdateRequest) (*api.Participation, error) {
	p, err := s.partRepo.FindBySlamAndUser(ctx, slamID, userID)
	if err != nil {
		return nil, errors.New("participation record not found")
	}

	if err := copier.Copy(&p, &req); err != nil {
		return nil, err
	}

	if err := s.partRepo.Update(ctx, &p); err != nil {
		return nil, err
	}

	apiPart := api.Participation{}
	copier.Copy(&apiPart, &p.Participation)
	return &apiPart, nil
}
