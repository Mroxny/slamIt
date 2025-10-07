package service

import (
	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/repository"
)

type StageService struct {
	stageRepo *repository.StageRepository
}

func NewStageService(stages *repository.StageRepository) *StageService {
	return &StageService{stageRepo: stages}
}

func (s *StageService) GetStages(slamId string) ([]api.Stage, error) {
	return s.stageRepo.GetBySlamID(slamId)
}

func (s *StageService) CreateStage(slamId string, stage api.StageRequest) (*api.Stage, error) {
	return s.stageRepo.Create(slamId, stage)
}

func (s *StageService) UpdateStage(stageId string, stage api.StageRequest) (*api.Stage, error) {
	return s.stageRepo.Update(stageId, stage)
}

func (s *StageService) DeleteStage(stageId string) error {
	return s.stageRepo.Delete(stageId)
}
