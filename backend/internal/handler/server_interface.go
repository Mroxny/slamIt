package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/service"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	userService *service.UserService
	slamService *service.SlamService
	authService *service.AuthService
	partService *service.SlamParticipationService
}

func NewServer(u *service.UserService, s *service.SlamService, a *service.AuthService, p *service.SlamParticipationService) *Server {
	return &Server{u, s, a, p}
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		_ = json.NewEncoder(w).Encode(v)
	}
}
