package repository

import (
	"context"

	"github.com/Mroxny/slamIt/internal/model"
	"gorm.io/gorm"
)

type StageRepository struct {
	*Repository[model.Stage]
}

func NewStageRepository(db *gorm.DB) *StageRepository {
	return &StageRepository{
		Repository: NewRepository[model.Stage](db),
	}
}

func (r *StageRepository) FindBySlamId(ctx context.Context, slmaId string) ([]model.Stage, error) {
	var stages []model.Stage
	err := r.db.WithContext(ctx).Preload("Performances").Find(&stages, "slam_id = ?", slmaId).Error
	return stages, err
}
