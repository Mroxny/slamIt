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

func (h *SlamParticipationHandler) Join(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	slamID, _ := strconv.Atoi(chi.URLParam(r, "slamID"))

	if err := h.service.Join(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *SlamParticipationHandler) Leave(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	slamID, _ := strconv.Atoi(chi.URLParam(r, "slamID"))

	if err := h.service.Leave(userID, slamID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *SlamParticipationHandler) GetSlamsForUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	slams, _ := h.service.GetSlamsForUser(userID)
	json.NewEncoder(w).Encode(slams)
}

func (h *SlamParticipationHandler) GetUsersForSlam(w http.ResponseWriter, r *http.Request) {
	slamID, _ := strconv.Atoi(chi.URLParam(r, "slamID"))
	users, _ := h.service.GetUsersForSlam(slamID)
	json.NewEncoder(w).Encode(users)
}
