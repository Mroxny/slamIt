package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/service"
	"github.com/Mroxny/slamIt/internal/utils"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetAll godoc
//
//	@Summary		Get all users
//	@Description	Retrieve a list of all registered users
//	@Tags			users
//	@Produce		json
//	@Success		200	{array}	model.User
//	@Router			/users/ [get]
//	@Security		BearerAuth
func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetAll()
	json.NewEncoder(w).Encode(users)
}

// GetByID godoc
//
//	@Summary		Get a user by ID
//	@Description	Retrieve a single user by their ID
//	@Tags			users
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	model.User
//	@Failure		404	{string}	string	"user not found"
//	@Router			/users/{id} [get]
//	@Security		BearerAuth
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Update godoc
//
//	@Summary		Update a user
//	@Description	Update an existing user by their ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string		true	"User ID"
//	@Param			user	body		model.User	true	"Updated user data"
//	@Success		200		{object}	model.User
//	@Failure		400		{string}	string	"invalid input"
//	@Failure		404		{string}	string	"user not found"
//	@Router			/users/{id} [put]
//	@Security		BearerAuth
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(u); err != nil {
		http.Error(w, "invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.service.Update(id, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

// Delete godoc
//
//	@Summary		Delete a user
//	@Description	Remove a user by their ID
//	@Tags			users
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		204	{string}	string	"no content"
//	@Failure		404	{string}	string	"user not found"
//	@Router			/users/{id} [delete]
//	@Security		BearerAuth
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.service.Delete(id); err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
