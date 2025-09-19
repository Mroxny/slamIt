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

func (r *SlamParticipationRepository) Add(userId string, slamId int) error {
	for _, rel := range r.relations {
		if rel.UserId == userId && rel.SlamId == slamId {
			return errors.New("user already joined this slam")
		}
	}
	r.relations = append(r.relations, model.SlamParticipation{
		UserId: userId,
		SlamId: slamId})
	return nil
}

func (r *SlamParticipationRepository) GetSlamsForUser(userId string) []int {
	ids := []int{}
	for _, rel := range r.relations {
		if rel.UserId == userId {
			ids = append(ids, rel.SlamId)
		}
	}
	return ids
}

func (r *SlamParticipationRepository) GetUsersForSlam(slamId int) []string {
	ids := []string{}
	for _, rel := range r.relations {
		if rel.SlamId == slamId {
			ids = append(ids, rel.UserId)
		}
	}
	return ids
}

func (r *SlamParticipationRepository) Remove(userId string, slamId int) error {
	for i, rel := range r.relations {
		if rel.UserId == userId && rel.SlamId == slamId {
			r.relations = append(r.relations[:i], r.relations[i+1:]...)
			return nil
		}
	}
	return errors.New("relation not found")
}
