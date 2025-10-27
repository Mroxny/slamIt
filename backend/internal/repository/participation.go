package repository

import (
	"context"
	"errors"

	"github.com/Mroxny/slamIt/internal/model"
	"gorm.io/gorm"
)

type ParticipationRepository struct {
	*Repository[model.Participation]
}

func NewParticipationRepository(db *gorm.DB) *ParticipationRepository {
	return &ParticipationRepository{
		Repository: NewRepository[model.Participation](db),
	}
}

func (r *ParticipationRepository) FindBySlamAndUser(ctx context.Context, slamID, userID string) (*model.Participation, error) {
	var slamCheck model.Slam
	if err := r.db.WithContext(ctx).Select("id").First(&slamCheck, "id = ?", slamID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("slam not found")
		}
		return nil, err
	}

	var userCheck model.User
	if err := r.db.WithContext(ctx).Select("id").First(&userCheck, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	var p model.Participation
	err := r.db.WithContext(ctx).Where("slam_id = ? AND user_id = ?", slamID, userID).First(&p).Error
	return &p, err
}

func (r *ParticipationRepository) DeleteBySlamAndUser(ctx context.Context, slamID, userID string) error {
	var participation model.Participation
	err := r.db.WithContext(ctx).
		Where("slam_id = ? AND user_id = ?", slamID, userID).
		First(&participation).Error

	if err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).Delete(&participation).Error; err != nil {
		return err
	}

	return nil
}

func (r *ParticipationRepository) FindParticipatingUsersBySlamID(ctx context.Context, slamID string, limit, offset int) ([]model.Participation, error) {
	var participations []model.Participation
	err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Preload("User").
		Preload("Stages").
		Where("slam_id = ?", slamID).
		Find(&participations).Error
	return participations, err
}

func (r *ParticipationRepository) FindParticipatedSlamsByUserID(ctx context.Context, userID string, limit, offset int) ([]model.Participation, error) {
	var participations []model.Participation
	err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Preload("Slam").
		Preload("Stages").
		Where("user_id = ?", userID).
		Find(&participations).Error
	return participations, err
}
