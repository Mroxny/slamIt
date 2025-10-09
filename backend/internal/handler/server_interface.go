package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/service"
	"github.com/Mroxny/slamIt/internal/utils"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	userService  *service.UserService
	slamService  *service.SlamService
	authService  *service.AuthService
	partService  *service.ParticipationService
	stageService *service.StageService
	perfService  *service.PerformanceService
	voteService  *service.VoteService
}

func NewServer(
	u *service.UserService,
	s *service.SlamService,
	a *service.AuthService,
	p *service.ParticipationService,
	st *service.StageService,
	pe *service.PerformanceService,
	v *service.VoteService,
) *Server {
	return &Server{u, s, a, p, st, pe, v}
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		json.NewEncoder(w).Encode(v)
	}
}

func ValidateJSON(decoder io.ReadCloser, object interface{}) error {
	if err := json.NewDecoder(decoder).Decode(object); err != nil {
		return errors.New("invalid input (decode)")
	}

	if err := utils.Validate.Struct(object); err != nil {
		return errors.New("invalid input (validation)")
	}

	return nil
}

func GetUserFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(utils.JWTClaimsContextKey)
	if val == nil {
		return "", errors.New("no user ID in context")
	}

	userID, ok := val.(string)
	if !ok {
		return "", errors.New("invalid user ID type in context")
	}

	return userID, nil
}
