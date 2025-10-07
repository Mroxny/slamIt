package service

import (
	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/repository"
)

type PerformanceService struct {
	perfRepo *repository.PerformanceRepository
}

func NewPerformanceService(performances *repository.PerformanceRepository) *PerformanceService {
	return &PerformanceService{perfRepo: performances}
}

func (s *PerformanceService) GetPerformances(stageId string) ([]api.Performance, error) {
	return s.perfRepo.GetByStageID(stageId)
}

func (s *PerformanceService) CreatePerformance(stageId string, p api.PerformanceRequest) (*api.Performance, error) {
	return s.perfRepo.Create(stageId, p)
}

func (s *PerformanceService) UpdatePerformance(performanceId string, p api.PerformanceRequest) (*api.Performance, error) {
	return s.perfRepo.Update(performanceId, p)
}

func (s *PerformanceService) DeletePerformance(performanceId string) error {
	return s.perfRepo.Delete(performanceId)
}
