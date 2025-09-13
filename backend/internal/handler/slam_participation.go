package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mroxny/slamIt/internal/service"
	"github.com/go-chi/chi/v5"
)

type SlamParticipationHandler struct {
	service *service.SlamParticipationService
}

func NewSlamParticipationHandler(service *service.SlamParticipationService) *SlamParticipationHandler {
	return &SlamParticipationHandler{service: service}
}

// Join godoc
// @Summary      Join a slam
// @Description  Add a user as a participant in a slam
// @Tags         participation
// @Produce      json
// @Param        userID  path  string  true  "User ID"
// @Param        slamID  path  int     true  "Slam ID"
// @Success      201  {string}  string "created"
// @Failure      400  {string}  string "bad request"
// @Router       /participation/users/{userID}/slams/{slamID} [post]
// @Security     BearerAuth
func (h *SlamParticipationHandler) Join(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	slamID, _ := strconv.Atoi(chi.URLParam(r, "slamID"))

	if err := h.service.Join(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// Leave godoc
// @Summary      Leave a slam
// @Description  Remove a user from a slam’s participants
// @Tags         participation
// @Produce      json
// @Param        userID  path  string  true  "User ID"
// @Param        slamID  path  int     true  "Slam ID"
// @Success      204  {string}  string "no content"
// @Failure      400  {string}  string "bad request"
// @Router       /participation/users/{userID}/slams/{slamID} [delete]
// @Security     BearerAuth
func (h *SlamParticipationHandler) Leave(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	slamID, _ := strconv.Atoi(chi.URLParam(r, "slamID"))

	if err := h.service.Leave(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetSlamsForUser godoc
// @Summary      List slams for a user
// @Description  Get all slams a user is participating in
// @Tags         participation
// @Produce      json
// @Param        userID  path  string  true  "User ID"
// @Success      200  {array}   model.Slam
// @Router       /participation/users/{userID}/slams [get]
// @Security     BearerAuth
func (h *SlamParticipationHandler) GetSlamsForUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	slams, _ := h.service.GetSlamsForUser(userID)
	json.NewEncoder(w).Encode(slams)
}

// GetUsersForSlam godoc
// @Summary      List users in a slam
// @Description  Get all users participating in a given slam
// @Tags         participation
// @Produce      json
// @Param        slamID  path  int  true  "Slam ID"
// @Success      200  {array}   model.User
// @Router       /participation/slams/{slamID}/users [get]
// @Security     BearerAuth
func (h *SlamParticipationHandler) GetUsersForSlam(w http.ResponseWriter, r *http.Request) {
	slamID, _ := strconv.Atoi(chi.URLParam(r, "slamID"))
	users, _ := h.service.GetUsersForSlam(slamID)
	json.NewEncoder(w).Encode(users)
}
