package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

var slamValidator = validator.New()

type SlamHandler struct {
	service *service.SlamService
}

func NewSlamHandler(service *service.SlamService) *SlamHandler {
	return &SlamHandler{service: service}
}

func (h *SlamHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.service.GetAll())
}

func (h *SlamHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	slam, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "slam not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(slam)
}

func (h *SlamHandler) Create(w http.ResponseWriter, r *http.Request) {
	var s model.Slam
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := slamValidator.Struct(s); err != nil {
		http.Error(w, "invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	created, err := h.service.Create(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *SlamHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var s model.Slam
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := slamValidator.Struct(s); err != nil {
		http.Error(w, "invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.service.Update(id, s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *SlamHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Delete(id); err != nil {
		http.Error(w, "slam not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
