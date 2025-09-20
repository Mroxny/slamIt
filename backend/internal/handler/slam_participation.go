package handler

import (
	"net/http"
)

func (s *Server) PostParticipationUsersUserIDSlamsSlamID(w http.ResponseWriter, r *http.Request, userID string, slamID int) {

	if err := s.partService.Join(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) DeleteParticipationUsersUserIDSlamsSlamID(w http.ResponseWriter, r *http.Request, userID string, slamID int) {
	if err := s.partService.Leave(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetParticipationUsersUserIDSlams(w http.ResponseWriter, r *http.Request, userID string) {
	slams, err := s.partService.GetSlamsForUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, slams)
}

func (s *Server) GetParticipationSlamsSlamIDUsers(w http.ResponseWriter, r *http.Request, slamID int) {
	users, err := s.partService.GetUsersForSlam(slamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, users)
}
