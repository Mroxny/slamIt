package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/service"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetAll()
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	created := h.service.Create(u)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
