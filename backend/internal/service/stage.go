package service

import (
	"context"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type StageService struct {
	stageRepo *repository.StageRepository
}

func NewStageService(stages *repository.StageRepository) *StageService {
	return &StageService{stageRepo: stages}
}

func (s *StageService) GetStages(ctx context.Context, slamId string) ([]api.Stage, error) {
	stages, err := s.stageRepo.FindBySlamId(ctx, slamId)
	if err != nil {
		return nil, err
	}
	var apiStages []api.Stage

	if err = copier.Copy(&apiStages, &stages); err != nil {
		return nil, err
	}
	return apiStages, nil
}

func (s *StageService) GetStage(ctx context.Context, stageId string) (*api.Stage, error) {
	stage, err := s.stageRepo.FindByID(ctx, stageId)
	if err != nil {
		return nil, err
	}
	var apiStage api.Stage

	if err = copier.Copy(&apiStage, &stage); err != nil {
		return nil, err
	}
	return &apiStage, nil
}

func (s *StageService) CreateStage(ctx context.Context, slamId string, stage api.StageRequest) (*api.Stage, error) {
	modelStage := model.Stage{}
	copier.Copy(&modelStage, &stage)
	modelStage.Id = uuid.New().String()
	modelStage.SlamId = slamId

	if err := s.stageRepo.Create(ctx, &modelStage); err != nil {
		return nil, err
	}

	apiStage := api.Stage{}
	copier.Copy(&apiStage, &modelStage)
	return &apiStage, nil
}

func (s *StageService) UpdateStage(ctx context.Context, stageId string, stage api.StageRequest) (*api.Stage, error) {
	modelStage := model.Stage{}
	copier.Copy(&modelStage, &stage)
	modelStage.Id = stageId

	if err := s.stageRepo.Update(ctx, &modelStage); err != nil {
		return nil, err
	}

	apiStage := api.Stage{}
	copier.Copy(&apiStage, &modelStage)
	return &apiStage, nil
}

func (s *StageService) DeleteStage(ctx context.Context, stageId string) error {
	return s.stageRepo.Delete(ctx, stageId)
}
