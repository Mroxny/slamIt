package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Mroxny/slamIt/internal/model"
	"github.com/Mroxny/slamIt/internal/service"
	"github.com/Mroxny/slamIt/internal/utils"
	"github.com/go-chi/chi/v5"
)

type SlamHandler struct {
	service *service.SlamService
}

func NewSlamHandler(service *service.SlamService) *SlamHandler {
	return &SlamHandler{service: service}
}

// GetAll godoc
//
//	@Summary		List slams
//	@Description	Get all public slams
//	@Tags			slams
//	@Produce		json
//	@Success		200	{array}	model.Slam
//	@Router			/slams/ [get]
func (h *SlamHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.service.GetAll())
}

// GetByID godoc
//
//	@Summary		Get a slam by ID
//	@Description	Retrieve a single slam by its ID
//	@Tags			slams
//	@Produce		json
//	@Param			id	path		int	true	"Slam ID"
//	@Success		200	{object}	model.Slam
//	@Failure		404	{string}	string	"slam not found"
//	@Router			/slams/{id} [get]
//	@Security		BearerAuth
func (h *SlamHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	slam, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "slam not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(slam)
}

// Create godoc
//
//	@Summary		Create a slam
//	@Description	Create a new slam with the provided input
//	@Tags			slams
//	@Accept			json
//	@Produce		json
//	@Param			slam	body		model.Slam	true	"Slam data"
//	@Success		201		{object}	model.Slam
//	@Failure		400		{string}	string	"invalid input"
//	@Router			/slams/ [post]
//	@Security		BearerAuth
func (h *SlamHandler) Create(w http.ResponseWriter, r *http.Request) {
	var s model.Slam
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(s); err != nil {
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

// Update godoc
//
//	@Summary		Update a slam
//	@Description	Update an existing slam by its ID
//	@Tags			slams
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"Slam ID"
//	@Param			slam	body		model.Slam	true	"Updated slam data"
//	@Success		200		{object}	model.Slam
//	@Failure		400		{string}	string	"invalid input"
//	@Failure		404		{string}	string	"slam not found"
//	@Router			/slams/{id} [put]
//	@Security		BearerAuth
func (h *SlamHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var s model.Slam
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(s); err != nil {
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

// Delete godoc
//
//	@Summary		Delete a slam
//	@Description	Remove a slam by its ID
//	@Tags			slams
//	@Produce		json
//	@Param			id	path		int		true	"Slam ID"
//	@Success		204	{string}	string	"no content"
//	@Failure		404	{string}	string	"slam not found"
//	@Router			/slams/{id} [delete]
//	@Security		BearerAuth
func (h *SlamHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Delete(id); err != nil {
		http.Error(w, "slam not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
