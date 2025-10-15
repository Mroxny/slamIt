package repository

import (
	"context"

	"github.com/Mroxny/slamIt/internal/model"
	"gorm.io/gorm"
)

type PerformanceRepository struct {
	*Repository[model.Performance]
}

func NewPerformanceRepository(db *gorm.DB) *PerformanceRepository {
	return &PerformanceRepository{
		Repository: NewRepository[model.Performance](db),
	}
}

func (r *PerformanceRepository) FindByStageId(ctx context.Context, stageId string) ([]model.Performance, error) {
	var performances []model.Performance
	err := r.db.WithContext(ctx).Preload("Particiaptions").Find(&performances, "stage_id = ?", stageId).Error
	return performances, err
}
