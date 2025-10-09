package handler

import (
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
)

func (s *Server) DeleteParticipationsSlamsSlamID(w http.ResponseWriter, r *http.Request, slamID string) {
	userID, err := GetUserFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.partService.Leave(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) PostParticipationsSlamsSlamID(w http.ResponseWriter, r *http.Request, slamID string) {
	userID, err := GetUserFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.partService.Join(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) GetParticipationsSlamsSlamIDUsers(w http.ResponseWriter, r *http.Request, slamID string) {
	users, err := s.partService.GetUsersForSlam(slamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, users)
}

func (s *Server) PostParticipationsSlamsSlamIDUsers(w http.ResponseWriter, r *http.Request, slamID string) {
	userID, err := GetUserFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.partService.Join(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) DeleteParticipationsSlamsSlamIDUsersUserID(w http.ResponseWriter, r *http.Request, slamID string, userID string) {
	if err := s.partService.Leave(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) PutParticipationsSlamsSlamIDUsersUserID(w http.ResponseWriter, r *http.Request, slamID string, userID string) {
	var participation api.ParticipationUpdateRequest
	if err := ValidateJSON(r.Body, &participation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := s.partService.UpdateParticipation(slamID, userID, participation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, updated)
}

func (s *Server) GetParticipationsUsersUserIDSlams(w http.ResponseWriter, r *http.Request, userID string) {
	slams, err := s.partService.GetSlamsForUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, slams)
}
