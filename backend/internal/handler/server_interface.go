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

const (
	defaultPage     = 1
	defaultPageSize = 10
	maxPageSize     = 100
)

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		json.NewEncoder(w).Encode(v)
	}
}

func ValidateJSON(reader io.ReadCloser, object interface{}) error {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(object); err != nil {
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

func ParsePageNumAndSize(page, pageSize *int) (int, int) {
	var resPage int
	var resSize int

	if page == nil || *page < 1 {
		resPage = defaultPage
	}
	resPage = *page

	if pageSize == nil || *pageSize < 1 {
		resSize = defaultPageSize
		return resPage, resSize
	}

	if *pageSize > maxPageSize {
		resSize = maxPageSize
		return resPage, resSize
	}
	resSize = *pageSize

	return resPage, resSize
}
