package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
)

type StageRepository struct {
	stages []api.Stage
	nextID int
}

func NewStageRepository() *StageRepository {
	return &StageRepository{
		stages: []api.Stage{},
		nextID: 1,
	}
}

func (r *StageRepository) GetBySlamID(slamId string) ([]api.Stage, error) {

	return nil, errors.New("stage not found")
}

func (r *StageRepository) Create(slamId string, s api.Stage) (*api.Stage, error) {

	return nil, errors.New("stage not found")
}

func (r *StageRepository) Update(slamId string, updated api.Stage) (*api.Stage, error) {

	return nil, errors.New("stage not found")
}

func (r *StageRepository) Delete(stageId string) error {

	return errors.New("stage not found")
}
