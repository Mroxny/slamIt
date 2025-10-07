package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
)

type PerformanceRepository struct {
	performances []api.Performance
	nextID       int
}

func NewPerformanceRepository() *PerformanceRepository {
	return &PerformanceRepository{
		performances: []api.Performance{},
		nextID:       1,
	}
}

func (r *PerformanceRepository) GetByStageID(stageId int) ([]api.Performance, error) {

	return nil, errors.New("performance not found")
}

func (r *PerformanceRepository) Create(stageId int, p api.Performance) (*api.Performance, error) {

	return nil, errors.New("performance not found")
}

func (r *PerformanceRepository) Update(stageId int, updated api.Performance) (*api.Performance, error) {

	return nil, errors.New("performance not found")
}

func (r *PerformanceRepository) Delete(performanceId int) error {

	return errors.New("performance not found")
}
