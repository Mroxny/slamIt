package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/utils"
)

func (s *Server) GetSlams(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.slamService.GetAll())
}

func (s *Server) GetSlamsId(w http.ResponseWriter, r *http.Request, id int) {
	slam, err := s.slamService.GetByID(id)
	if err != nil {
		http.Error(w, "slam not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(slam)
}

func (s *Server) PostSlams(w http.ResponseWriter, r *http.Request) {
	var slam api.Slam
	if err := json.NewDecoder(r.Body).Decode(&slam); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(slam); err != nil {
		http.Error(w, "invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	created, err := s.slamService.Create(slam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (s *Server) PutSlamsId(w http.ResponseWriter, r *http.Request, id int) {
	var slam api.Slam
	if err := json.NewDecoder(r.Body).Decode(&slam); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(slam); err != nil {
		http.Error(w, "invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := s.slamService.Update(id, slam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (s *Server) DeleteSlamsId(w http.ResponseWriter, r *http.Request, id int) {
	if err := s.slamService.Delete(id); err != nil {
		http.Error(w, "slam not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
