package service

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
)

type SlamParticipationService struct {
	usersRepo *repository.UserRepository
	slamsRepo *repository.SlamRepository
	repo      *repository.SlamParticipationRepository
}

func NewSlamParticipationService(users *repository.UserRepository, slams *repository.SlamRepository, repo *repository.SlamParticipationRepository) *SlamParticipationService {
	return &SlamParticipationService{usersRepo: users, slamsRepo: slams, repo: repo}
}

func (s *SlamParticipationService) Join(userID string, slamID int) error {
	if _, err := s.usersRepo.GetByID(userID); err != nil {
		return errors.New("user not found")
	}
	if _, err := s.slamsRepo.GetByID(slamID); err != nil {
		return errors.New("slam not found")
	}
	return s.repo.Add(userID, slamID)
}

func (s *SlamParticipationService) Leave(userID string, slamID int) error {
	return s.repo.Remove(userID, slamID)
}

func (s *SlamParticipationService) GetSlamsForUser(userID string) ([]model.Slam, error) {
	ids := s.repo.GetSlamsForUser(userID)
	slams := []model.Slam{}
	for _, id := range ids {
		if slam, err := s.slamsRepo.GetByID(id); err == nil {
			slams = append(slams, *slam)
		}
	}
	return slams, nil
}

func (s *SlamParticipationService) GetUsersForSlam(slamID int) ([]model.User, error) {
	ids := s.repo.GetUsersForSlam(slamID)
	users := []model.User{}
	for _, id := range ids {
		if user, err := s.usersRepo.GetByID(id); err == nil {
			users = append(users, *user)
		}
	}
	return users, nil
}
