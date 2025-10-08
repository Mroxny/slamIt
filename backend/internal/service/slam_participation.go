package service

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/repository"
)

type SlamParticipationService struct {
	usersRepo *repository.UserRepository
	slamsRepo *repository.SlamRepository
	partRepo  *repository.SlamParticipationRepository
}

func NewSlamParticipationService(users *repository.UserRepository, slams *repository.SlamRepository, participations *repository.SlamParticipationRepository) *SlamParticipationService {
	return &SlamParticipationService{usersRepo: users, slamsRepo: slams, partRepo: participations}
}

func (s *SlamParticipationService) Join(userID string, slamID string) error {
	if _, err := s.usersRepo.GetByID(userID); err != nil {
		return errors.New("user not found")
	}
	if _, err := s.slamsRepo.GetByID(slamID); err != nil {
		return errors.New("slam not found")
	}
	return s.partRepo.Add(userID, slamID)
}

func (s *SlamParticipationService) Leave(userID string, slamID string) error {
	return s.partRepo.Remove(userID, slamID)
}

func (s *SlamParticipationService) GetSlamsForUser(userID string) ([]api.Slam, error) {
	ids := s.partRepo.GetSlamsForUser(userID)
	slams := []api.Slam{}
	for _, id := range ids {
		if slam, err := s.slamsRepo.GetByID(id); err == nil {
			slams = append(slams, *slam)
		}
	}
	return slams, nil
}

func (s *SlamParticipationService) UpdateParticipation(slamID string, userID string, p api.ParticipationUpdateRequest) (*api.Participation, error) {
	return s.partRepo.UpdateParticipation(slamID, userID, p)
}

func (s *SlamParticipationService) GetUsersForSlam(slamID string) ([]api.User, error) {
	ids := s.partRepo.GetUsersForSlam(slamID)
	users := []api.User{}
	for _, id := range ids {
		if user, err := s.usersRepo.GetByID(id); err == nil {
			u := api.User{
				Id:    &user.Id,
				Email: &user.Email,
				Name:  &user.Name,
			}
			users = append(users, u)
		}
	}
	return users, nil
}
