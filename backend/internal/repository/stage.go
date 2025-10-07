package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
)

type StageRepository struct {
	stages []api.Stage
}

func NewStageRepository() *StageRepository {
	return &StageRepository{stages: []api.Stage{}}
}

func (r *StageRepository) GetBySlamID(slamId string) ([]api.Stage, error) {

	return nil, errors.New("stage not found")
}

func (r *StageRepository) Create(slamId string, s api.StageRequest) (*api.Stage, error) {

	return nil, errors.New("stage not found")
}

func (r *StageRepository) Update(stageId string, updated api.StageRequest) (*api.Stage, error) {

	return nil, errors.New("stage not found")
}

func (r *StageRepository) Delete(stageId string) error {

	return errors.New("stage not found")
}
