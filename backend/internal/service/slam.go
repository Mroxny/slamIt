package service

import (
	"context"
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/repository"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type SlamService struct {
	slamRepo *repository.SlamRepository
}

func NewSlamService(slams *repository.SlamRepository) *SlamService {
	return &SlamService{slamRepo: slams}
}

func (s *SlamService) GetAll(ctx context.Context) ([]api.Slam, error) {
	modelSlams, err := s.slamRepo.FindAllPublic(ctx)
	if err != nil {
		return nil, err
	}

	var apiSlams []api.Slam
	if err := copier.Copy(&apiSlams, &modelSlams); err != nil {
		return nil, err
	}

	return apiSlams, nil
}

func (s *SlamService) GetByID(ctx context.Context, id string) (*api.Slam, error) {
	modelSlam, err := s.slamRepo.FindPublicByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("slam not found")
		}
		return nil, err
	}

	var apiSlam api.Slam
	if err := copier.Copy(&apiSlam, &modelSlam); err != nil {
		return nil, err
	}

	return &apiSlam, nil
}

func (s *SlamService) Create(ctx context.Context, slam api.SlamRequest, userId string) (*api.Slam, error) {
	modelSlam := model.Slam{}
	copier.Copy(&modelSlam, &slam)
	modelSlam.Id = uuid.New().String()

	if err := s.slamRepo.CreateWithCreatorTx(ctx, &modelSlam, userId); err != nil {
		return nil, err
	}

	apiSlam := api.Slam{}
	copier.Copy(&apiSlam, &modelSlam)
	return &apiSlam, nil
}

func (s *SlamService) Update(ctx context.Context, id string, slam api.SlamRequest) (*api.Slam, error) {
	modelSlam := model.Slam{}
	copier.Copy(&modelSlam, &slam)
	modelSlam.Id = id

	if err := s.slamRepo.Update(ctx, &modelSlam); err != nil {
		return nil, err
	}

	apiSlam := api.Slam{}
	copier.Copy(&apiSlam, &modelSlam)
	return &apiSlam, nil
}

func (s *SlamService) Delete(ctx context.Context, id string) error {
	return s.slamRepo.Delete(ctx, id)
}
