package repository

import (
	"context"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SlamRepository struct {
	*Repository[model.Slam]
}

func NewSlamRepository(db *gorm.DB) *SlamRepository {
	return &SlamRepository{
		Repository: NewRepository[model.Slam](db),
	}
}

func (r *SlamRepository) FindAllPublic(ctx context.Context, limit, offset int) ([]model.Slam, error) {
	var slams []model.Slam
	err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&slams, "public = true").Error
	return slams, err
}

func (r *SlamRepository) FindByID(ctx context.Context, id string) (*model.Slam, error) {
	var slam model.Slam
	err := r.db.WithContext(ctx).
		Preload("Users").
		Preload("Stages").
		First(&slam, "id = ?", id).Error
	return &slam, err
}

func (r *SlamRepository) FindPublicByID(ctx context.Context, id string) (*model.Slam, error) {
	var slam model.Slam
	err := r.db.WithContext(ctx).
		Where("public = true").
		Preload("Users").
		Preload("Stages").
		First(&slam, "id = ?", id).Error
	return &slam, err
}

// CreateWithCreatorTx creates a slam and a participation record for the creator in a single transaction.
func (r *SlamRepository) CreateWithCreatorTx(ctx context.Context, slam *model.Slam, userId string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(slam).Error; err != nil {
			return err
		}

		participation := model.Participation{
			Participation: api.Participation{
				Id:     uuid.New().String(),
				Role:   api.Creator,
				UserId: userId,
				SlamId: slam.Id,
			},
		}

		if err := tx.Create(&participation).Error; err != nil {
			return err
		}

		return nil
	})
}
