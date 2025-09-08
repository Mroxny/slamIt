package repository

import "errors"

type SlamParticipationRepository struct {
	relations []struct {
		UserID int
		SlamID int
	}
}

func NewSlamParticipationRepository() *SlamParticipationRepository {
	return &SlamParticipationRepository{
		relations: []struct {
			UserID int
			SlamID int
		}{},
	}
}

// User joins a slam
func (r *SlamParticipationRepository) Add(userID, slamID int) error {
	// prevent duplicates
	for _, rel := range r.relations {
		if rel.UserID == userID && rel.SlamID == slamID {
			return errors.New("user already joined this slam")
		}
	}
	r.relations = append(r.relations, struct {
		UserID int
		SlamID int
	}{UserID: userID, SlamID: slamID})
	return nil
}

// List all slams for a user
func (r *SlamParticipationRepository) GetSlamsForUser(userID int) []int {
	ids := []int{}
	for _, rel := range r.relations {
		if rel.UserID == userID {
			ids = append(ids, rel.SlamID)
		}
	}
	return ids
}

// List all users for a slam
func (r *SlamParticipationRepository) GetUsersForSlam(slamID int) []int {
	ids := []int{}
	for _, rel := range r.relations {
		if rel.SlamID == slamID {
			ids = append(ids, rel.UserID)
		}
	}
	return ids
}

// Remove participation
func (r *SlamParticipationRepository) Remove(userID, slamID int) error {
	for i, rel := range r.relations {
		if rel.UserID == userID && rel.SlamID == slamID {
			r.relations = append(r.relations[:i], r.relations[i+1:]...)
			return nil
		}
	}
	return errors.New("relation not found")
}
