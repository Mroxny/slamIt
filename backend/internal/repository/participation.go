package repository

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/model"
)

type ParticipationRepository struct {
	relations []model.Participation
}

func NewParticipationRepository() *ParticipationRepository {
	return &ParticipationRepository{
		relations: []model.Participation{},
	}
}

func (r *ParticipationRepository) Add(userId string, slamId string) error {
	for _, rel := range r.relations {
		if rel.UserId == userId && rel.SlamId == slamId {
			return errors.New("user already joined this slam")
		}
	}
	r.relations = append(r.relations, model.Participation{
		UserId: userId,
		SlamId: slamId})
	return nil
}

func (r *ParticipationRepository) GetSlamsForUser(userId string) []string {
	ids := []string{}
	for _, rel := range r.relations {
		if rel.UserId == userId {
			ids = append(ids, rel.SlamId)
		}
	}
	return ids
}

func (r *ParticipationRepository) GetUsersForSlam(slamId string) []string {
	ids := []string{}
	for _, rel := range r.relations {
		if rel.SlamId == slamId {
			ids = append(ids, rel.UserId)
		}
	}
	return ids
}

func (r *ParticipationRepository) UpdateParticipation(slamId string, userId string, p api.ParticipationUpdateRequest) (*api.Participation, error) {
	return nil, errors.New("error updating participation")
}

func (r *ParticipationRepository) Remove(userId string, slamId string) error {
	for i, rel := range r.relations {
		if rel.UserId == userId && rel.SlamId == slamId {
			r.relations = append(r.relations[:i], r.relations[i+1:]...)
			return nil
		}
	}
	return errors.New("relation not found")
}
