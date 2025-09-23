package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
)

type SlamRepository struct {
	slams  []api.Slam
	nextID int
}

func NewSlamRepository() *SlamRepository {
	return &SlamRepository{
		slams:  []api.Slam{},
		nextID: 1,
	}
}

func (r *SlamRepository) GetAll() []api.Slam {
	return r.slams
}

func (r *SlamRepository) GetByID(id int) (*api.Slam, error) {
	for _, s := range r.slams {
		if *s.Id == id {
			return &s, nil
		}
	}
	return nil, errors.New("slam not found")
}

func (r *SlamRepository) Create(s api.Slam) (api.Slam, error) {
	if s.Title == "" {
		return api.Slam{}, errors.New("title required")
	}
	s.Id = &r.nextID
	r.nextID++
	r.slams = append(r.slams, s)
	return s, nil
}

func (r *SlamRepository) Update(id int, updated api.Slam) (*api.Slam, error) {
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

func (r *SlamRepository) Delete(id int) error {
	for i, s := range r.slams {
		if *s.Id == id {
			r.slams = append(r.slams[:i], r.slams[i+1:]...)
			return nil
		}
	}
	return errors.New("slam not found")
}
