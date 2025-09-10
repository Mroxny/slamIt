package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/model"
)

type SlamParticipationRepository struct {
	relations []model.SlamParticipation
}

func NewSlamParticipationRepository() *SlamParticipationRepository {
	return &SlamParticipationRepository{
		relations: []model.SlamParticipation{},
	}
}

func (r *SlamParticipationRepository) Add(userID string, slamID int) error {
	for _, rel := range r.relations {
		if rel.UserID == userID && rel.SlamID == slamID {
			return errors.New("user already joined this slam")
		}
	}
	r.relations = append(r.relations, model.SlamParticipation{
		UserID: userID,
		SlamID: slamID})
	return nil
}

func (r *SlamParticipationRepository) GetSlamsForUser(userID string) []int {
	ids := []int{}
	for _, rel := range r.relations {
		if rel.UserID == userID {
			ids = append(ids, rel.SlamID)
		}
	}
	return ids
}

func (r *SlamParticipationRepository) GetUsersForSlam(slamID int) []string {
	ids := []string{}
	for _, rel := range r.relations {
		if rel.SlamID == slamID {
			ids = append(ids, rel.UserID)
		}
	}
	return ids
}

func (r *SlamParticipationRepository) Remove(userID string, slamID int) error {
	for i, rel := range r.relations {
		if rel.UserID == userID && rel.SlamID == slamID {
			r.relations = append(r.relations[:i], r.relations[i+1:]...)
			return nil
		}
	}
	return errors.New("relation not found")
}
