package service

import (
	"context"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type PerformanceService struct {
	perfRepo *repository.PerformanceRepository
}

func NewPerformanceService(performances *repository.PerformanceRepository) *PerformanceService {
	return &PerformanceService{perfRepo: performances}
}

func (s *PerformanceService) GetPerformances(ctx context.Context, stageId string) ([]api.Performance, error) {
	perfs, err := s.perfRepo.FindSortedByStageId(ctx, stageId)
	if err != nil {
		return nil, err
	}
	var apiPerfs []api.Performance

	if err = copier.Copy(&apiPerfs, &perfs); err != nil {
		return nil, err
	}
	return apiPerfs, nil
}

func (s *PerformanceService) GetPerformance(ctx context.Context, performanceId string) (*api.Performance, error) {
	perf, err := s.perfRepo.FindByID(ctx, performanceId)
	if err != nil {
		return nil, err
	}
	var apiPerf api.Performance

	if err = copier.Copy(&apiPerf, &perf); err != nil {
		return nil, err
	}
	return &apiPerf, nil
}

func (s *PerformanceService) CreatePerformance(ctx context.Context, stageId string, p api.PerformanceRequest) (*api.Performance, error) {
	modelPerf := model.Performance{}
	copier.Copy(&modelPerf, &p)
	modelPerf.Id = uuid.New().String()
	modelPerf.StageId = stageId

	if err := s.perfRepo.Create(ctx, &modelPerf); err != nil {
		return nil, err
	}

	apiPerf := api.Performance{}
	copier.Copy(&apiPerf, &modelPerf)
	return &apiPerf, nil
}

func (s *PerformanceService) UpdatePerformance(ctx context.Context, performanceId string, p api.PerformanceRequest) (*api.Performance, error) {
	modelPerf := model.Performance{}
	copier.Copy(&modelPerf, &p)
	modelPerf.Id = performanceId

	if err := s.perfRepo.Update(ctx, &modelPerf); err != nil {
		return nil, err
	}

	apiPerf := api.Performance{}
	copier.Copy(&apiPerf, &modelPerf)
	return &apiPerf, nil
}

func (s *PerformanceService) UpdatePerformanceOrder(ctx context.Context, stageId string, orderedIDs []string) error {
	return s.perfRepo.UpdateOrderTx(ctx, stageId, orderedIDs)
}

func (s *PerformanceService) DeletePerformance(ctx context.Context, performanceId string) error {
	return s.perfRepo.Delete(ctx, performanceId)

}
