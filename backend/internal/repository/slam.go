package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/jinzhu/copier"
)

type SlamRepository struct {
	slams []api.Slam
}

func NewSlamRepository() *SlamRepository {
	return &SlamRepository{
		slams: []api.Slam{},
	}
}

func (r *SlamRepository) GetAll() ([]api.Slam, error) {
	return r.slams, nil
}

func (r *SlamRepository) GetByID(id string) (*api.Slam, error) {
	for _, s := range r.slams {
		if *s.Id == id {
			return &s, nil
		}
	}
	return nil, errors.New("slam not found")
}

func (r *SlamRepository) Create(s api.SlamRequest) (api.Slam, error) {
	if s.Title == "" {
		return api.Slam{}, errors.New("title required")
	}
	modelSlam := api.Slam{}
	copier.Copy(&modelSlam, &s)

	r.slams = append(r.slams, modelSlam)
	return modelSlam, nil
}

func (r *SlamRepository) Update(id string, updated api.SlamRequest) (*api.Slam, error) {
	for i, s := range r.slams {
		if *s.Id == id {
			if updated.Title == "" {
				return nil, errors.New("title required")
			}
			r.slams[i].Title = updated.Title
			r.slams[i].Description = updated.Description
			return &r.slams[i], nil
		}
	}
	return nil, errors.New("slam not found")
}

func (r *SlamRepository) Delete(id string) error {
	for i, s := range r.slams {
		if *s.Id == id {
			r.slams = append(r.slams[:i], r.slams[i+1:]...)
			return nil
		}
	}
	return errors.New("slam not found")
}
