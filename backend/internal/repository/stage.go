package repository

import (
	"context"
	"errors"

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

func (r *StageRepository) FindBySlamId(ctx context.Context, slamId string, limit, offset int) ([]model.Stage, error) {
	var slamCheck model.Slam
	if err := r.db.WithContext(ctx).Select("id").First(&slamCheck, "id = ?", slamId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("slam not found")
		}
		return nil, err
	}

	var stages []model.Stage
	err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Preload("Participations").
		Find(&stages, "slam_id = ?", slamId).Error
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

func (r *StageRepository) Create(ctx context.Context, stage *model.Stage) error {
	var slamCheck model.Slam
	if err := r.db.WithContext(ctx).Select("id").First(&slamCheck, "id = ?", stage.SlamId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("slam not found")
		}
		return err
	}
	return r.db.WithContext(ctx).Create(stage).Error
}
