package service

import (
	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/repository"
)

type SlamService struct {
	slamRepo *repository.SlamRepository
}

func NewSlamService(slams *repository.SlamRepository) *SlamService {
	return &SlamService{slamRepo: slams}
}

func (s *SlamService) GetAll() []api.Slam {
	return s.slamRepo.GetAll()
}

func (s *SlamService) GetByID(id string) (*api.Slam, error) {
	return s.slamRepo.GetByID(id)
}

func (s *SlamService) Create(slam api.Slam) (api.Slam, error) {
	return s.slamRepo.Create(slam)
}

func (s *SlamService) Update(id string, slam api.Slam) (*api.Slam, error) {
	return s.slamRepo.Update(id, slam)
}

func (s *SlamService) Delete(id string) error {
	return s.slamRepo.Delete(id)
}
