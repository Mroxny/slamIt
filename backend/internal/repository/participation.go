package repository

import (
	"context"

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

func (r *ParticipationRepository) FindParticipatingUsersBySlamID(ctx context.Context, slamID string) ([]model.Participation, error) {
	var participations []model.Participation
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Stages").
		Where("slam_id = ?", slamID).
		Find(&participations).Error
	return participations, err
}

func (r *ParticipationRepository) FindParticipatedSlamsByUserID(ctx context.Context, userID string) ([]model.Participation, error) {
	var participations []model.Participation
	err := r.db.WithContext(ctx).
		Preload("Slam").
		Preload("Stages").
		Where("user_id = ?", userID).
		Find(&participations).Error
	return participations, err
}
