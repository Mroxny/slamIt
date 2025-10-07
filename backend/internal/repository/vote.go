package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
)

type VoteRepository struct {
	votes []api.Vote
}

func NewVoteRepository() *VoteRepository {
	return &VoteRepository{votes: []api.Vote{}}
}

func (r *VoteRepository) GetByPerformanceID(performanceId string) ([]api.Vote, error) {

	return nil, errors.New("vote not found")
}

func (r *VoteRepository) Create(performanceId string, v api.VoteRequest) (*api.Vote, error) {

	return nil, errors.New("vote not found")
}

func (r *VoteRepository) Update(voteId string, updated api.VoteRequest) (*api.Vote, error) {

	return nil, errors.New("vote not found")
}

func (r *VoteRepository) Delete(voteId string) error {

	return errors.New("vote not found")
}
