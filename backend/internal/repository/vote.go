package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
)

type VodeRepository struct {
	performances []api.Vote
	nextID       int
}

func NewVoteRepository() *VodeRepository {
	return &VodeRepository{
		performances: []api.Vote{},
		nextID:       1,
	}
}

func (r *VodeRepository) GetByPerformanceID(performanceId int) ([]api.Vote, error) {

	return nil, errors.New("stage not found")
}

func (r *VodeRepository) Create(performanceId int, p api.Vote) (*api.Vote, error) {

	return nil, errors.New("stage not found")
}

func (r *VodeRepository) Update(performanceId int, updated api.Vote) (*api.Vote, error) {

	return nil, errors.New("stage not found")
}

func (r *VodeRepository) Delete(voteId int) error {

	return errors.New("stage not found")
}
