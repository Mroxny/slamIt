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

	if err := s.partService.RemoveUserFromSlam(r.Context(), userID, slamID); err != nil {
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

	part, err := s.partService.AddUserToSlam(r.Context(), userID, slamID, api.Performer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusCreated, part)
}

func (s *Server) GetParticipationsSlamsSlamIDUsers(w http.ResponseWriter, r *http.Request, slamID string, params api.GetParticipationsSlamsSlamIDUsersParams) {
	page, pageSize := ParsePageNumAndSize(params.Page, params.PageSize)
	users, err := s.partService.GetUsersForSlam(r.Context(), slamID, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, users)
}

func (s *Server) PostParticipationsSlamsSlamIDUsers(w http.ResponseWriter, r *http.Request, slamID string) {
	var req api.ParticipationRequest
	if err := ValidateJSON(r.Body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Role == nil {
		role := api.Performer
		req.Role = &role
	}

	part, err := s.partService.AddUserToSlam(r.Context(), req.UserId, slamID, *req.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusCreated, part)
}

func (s *Server) DeleteParticipationsSlamsSlamIDUsersUserID(w http.ResponseWriter, r *http.Request, slamID string, userID string) {
	if err := s.partService.RemoveUserFromSlam(r.Context(), userID, slamID); err != nil {
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

	updated, err := s.partService.UpdateParticipation(r.Context(), slamID, userID, participation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, updated)
}

func (s *Server) GetParticipationsUsersUserIDSlams(w http.ResponseWriter, r *http.Request, userID string, params api.GetParticipationsUsersUserIDSlamsParams) {
	page, pageSize := ParsePageNumAndSize(params.Page, params.PageSize)
	slams, err := s.partService.GetSlamsForUser(r.Context(), userID, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, slams)
}
