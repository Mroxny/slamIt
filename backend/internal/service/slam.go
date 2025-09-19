package service

import (
	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/repository"
)

type SlamService struct {
	repo *repository.SlamRepository
}

func NewSlamService(repo *repository.SlamRepository) *SlamService {
	return &SlamService{repo: repo}
}

func (s *SlamService) GetAll() []api.Slam {
	return s.repo.GetAll()
}

func (s *SlamService) GetByID(id int) (*api.Slam, error) {
	return s.repo.GetByID(id)
}

func (s *SlamService) Create(slam api.Slam) (api.Slam, error) {
	return s.repo.Create(slam)
}

func (s *SlamService) Update(id int, slam api.Slam) (*api.Slam, error) {
	return s.repo.Update(id, slam)
}

func (s *SlamService) Delete(id int) error {
	return s.repo.Delete(id)
}
