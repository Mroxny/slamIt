package handler

import (
	"encoding/json"
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
	slams, _ := s.partService.GetSlamsForUser(userID)
	json.NewEncoder(w).Encode(slams)
}

func (s *Server) GetParticipationSlamsSlamIDUsers(w http.ResponseWriter, r *http.Request, slamID int) {
	users, _ := s.partService.GetUsersForSlam(slamID)
	json.NewEncoder(w).Encode(users)
}
