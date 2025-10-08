package handler

import (
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
)

func (s *Server) PostAuthRegister(w http.ResponseWriter, r *http.Request) {
	var req api.RegisterRequest
	if err := ValidateJSON(r.Body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := s.authService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteJSON(w, http.StatusCreated, user)
}

func (s *Server) PostAuthLogin(w http.ResponseWriter, r *http.Request) {
	var req api.LoginRequest
	if err := ValidateJSON(r.Body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := s.authService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	WriteJSON(w, http.StatusOK, response)
}
