package handler

import (
	"net/http"
)

func (s *Server) PostParticipationsUsersUserIDSlamsSlamID(w http.ResponseWriter, r *http.Request, userID string, slamID int) {

	if err := s.partService.Join(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) DeleteParticipationsUsersUserIDSlamsSlamID(w http.ResponseWriter, r *http.Request, userID string, slamID int) {
	if err := s.partService.Leave(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetParticipationsUsersUserIDSlams(w http.ResponseWriter, r *http.Request, userID string) {
	slams, err := s.partService.GetSlamsForUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, slams)
}

func (s *Server) GetParticipationsSlamsSlamIDUsers(w http.ResponseWriter, r *http.Request, slamID int) {
	users, err := s.partService.GetUsersForSlam(slamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, users)
}
