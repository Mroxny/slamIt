package service

import (
	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/repository"
)

type VoteService struct {
	voteRepo *repository.VoteRepository
}

func NewVoteService(votes *repository.VoteRepository) *VoteService {
	return &VoteService{voteRepo: votes}
}

func (s *VoteService) GetVotes(performanceId string) ([]api.Vote, error) {
	return s.voteRepo.GetByPerformanceID(performanceId)
}

func (s *VoteService) CreateVote(performanceId string, vote api.VoteRequest) (*api.Vote, error) {
	return s.voteRepo.Create(performanceId, vote)
}

func (s *VoteService) UpdateVote(voteId string, vote api.VoteRequest) (*api.Vote, error) {
	return s.voteRepo.Update(voteId, vote)
}

func (s *VoteService) DeleteVotes(voteId string) error {
	return s.voteRepo.Delete(voteId)
}
