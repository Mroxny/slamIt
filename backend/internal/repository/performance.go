package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
)

type PerformanceRepository struct {
	performances []api.Performance
}

func NewPerformanceRepository() *PerformanceRepository {
	return &PerformanceRepository{performances: []api.Performance{}}
}

func (r *PerformanceRepository) GetByStageID(stageId string) ([]api.Performance, error) {

	return nil, errors.New("performance not found")
}

func (r *PerformanceRepository) Create(stageId string, p api.PerformanceRequest) (*api.Performance, error) {

	return nil, errors.New("performance not found")
}

func (r *PerformanceRepository) Update(performanceId string, updated api.PerformanceRequest) (*api.Performance, error) {

	return nil, errors.New("performance not found")
}

func (r *PerformanceRepository) Delete(performanceId string) error {

	return errors.New("performance not found")
}
