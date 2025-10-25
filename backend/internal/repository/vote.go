package repository

import (
	"context"

	"github.com/Mroxny/slamIt/internal/model"
	"gorm.io/gorm"
)

type VoteRepository struct {
	*Repository[model.Vote]
}

func NewVoteRepository(db *gorm.DB) *VoteRepository {
	return &VoteRepository{
		Repository: NewRepository[model.Vote](db),
	}
}

func (r *VoteRepository) FindAllByPerformanceID(ctx context.Context, performanceID string) ([]model.Vote, error) {
	var votes []model.Vote
	err := r.db.WithContext(ctx).Where("performance_id = ?", performanceID).Find(&votes).Error
	return votes, err
}
