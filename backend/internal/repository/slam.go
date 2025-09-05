package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/model"
)

type SlamRepository struct {
	slams  []model.Slam
	nextID int
}

func NewSlamRepository() *SlamRepository {
	return &SlamRepository{
		slams:  []model.Slam{},
		nextID: 1,
	}
}

func (r *SlamRepository) GetAll() []model.Slam {
	return r.slams
}

func (r *SlamRepository) GetByID(id int) (*model.Slam, error) {
	for _, s := range r.slams {
		if s.ID == id {
			return &s, nil
		}
	}
	return nil, errors.New("slam not found")
}

func (r *SlamRepository) Create(s model.Slam) (model.Slam, error) {
	if s.Title == "" {
		return model.Slam{}, errors.New("title required")
	}
	s.ID = r.nextID
	r.nextID++
	r.slams = append(r.slams, s)
	return s, nil
}

func (r *SlamRepository) Update(id int, updated model.Slam) (*model.Slam, error) {
	for i, s := range r.slams {
		if s.ID == id {
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
		if s.ID == id {
			r.slams = append(r.slams[:i], r.slams[i+1:]...)
			return nil
		}
	}
	return errors.New("slam not found")
}
