package handler

import (
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
)

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request, params api.GetUsersParams) {
	page, pageSize := ParsePageNumAndSize(params.Page, params.PageSize)
	users, err := s.userService.GetAll(r.Context(), page, pageSize)
	if err != nil {
		http.Error(w, "error parsing users", http.StatusInternalServerError)
		return
	}
	WriteJSON(w, http.StatusOK, users)
}

func (s *Server) PostUsers(w http.ResponseWriter, r *http.Request) {
	var user api.UserRequest
	if err := ValidateJSON(r.Body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := s.userService.CreateTmpUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusCreated, created)
}

func (s *Server) DeleteUsersUserID(w http.ResponseWriter, r *http.Request, userID string) {
	if err := s.userService.Delete(r.Context(), userID); err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetUsersUserID(w http.ResponseWriter, r *http.Request, userID string) {
	user, err := s.userService.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	WriteJSON(w, http.StatusOK, user)
}

func (s *Server) PutUsersUserID(w http.ResponseWriter, r *http.Request, userID string) {
	var u api.UserRequest
	if err := ValidateJSON(r.Body, &u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := s.userService.Update(r.Context(), userID, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, updated)
}
