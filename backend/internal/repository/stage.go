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
	err := r.db.WithContext(ctx).Preload("Participations").Find(&stages, "slam_id = ?", slmaId).Error
	return stages, err
}

func (r *StageRepository) FindByID(ctx context.Context, id string) (*model.Stage, error) {
	var stage model.Stage
	err := r.db.WithContext(ctx).
		Preload("Participations").
		Preload("Participations.User").
		First(&stage, "id = ?", id).Error
	return &stage, err
}
