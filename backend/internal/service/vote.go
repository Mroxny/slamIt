package service

import (
	"context"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type VoteService struct {
	voteRepo *repository.VoteRepository
}

func NewVoteService(votes *repository.VoteRepository) *VoteService {
	return &VoteService{voteRepo: votes}
}

func (s *VoteService) GetVotes(ctx context.Context, performanceId string, page, pageSize int) (*api.VotePagination, error) {
	offset := (page - 1) * pageSize
	votes, err := s.voteRepo.FindAllByPerformanceID(ctx, performanceId, pageSize, offset)
	if err != nil {
		return nil, err
	}
	var apiVotes []api.Vote

	if err = copier.Copy(&apiVotes, &votes); err != nil {
		return nil, err
	}

	pag := api.VotePagination{
		Page:     &page,
		PageSize: &pageSize,
		Items:    &apiVotes,
	}
	return &pag, nil
}

func (s *VoteService) CreateVote(ctx context.Context, performanceId string, vote api.VoteRequest) (*api.Vote, error) {
	modelVote := model.Vote{}
	copier.Copy(&modelVote, &vote)
	modelVote.Id = uuid.New().String()
	modelVote.PerformanceId = performanceId

	if err := s.voteRepo.Create(ctx, &modelVote); err != nil {
		return nil, err
	}

	apiVote := api.Vote{}
	copier.Copy(&apiVote, &modelVote)
	return &apiVote, nil
}

func (s *VoteService) UpdateVote(ctx context.Context, voteId string, vote api.VoteRequest) (*api.Vote, error) {
	modelVote := model.Vote{}
	copier.Copy(&modelVote, &vote)
	modelVote.Id = voteId

	if err := s.voteRepo.Update(ctx, &modelVote); err != nil {
		return nil, err
	}

	apiVote := api.Vote{}
	copier.Copy(&apiVote, &modelVote)
	return &apiVote, nil
}

func (s *VoteService) DeleteVotes(ctx context.Context, voteId string) error {
	return s.voteRepo.Delete(ctx, voteId)
}
