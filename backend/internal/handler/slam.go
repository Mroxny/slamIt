package handler

import (
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
)

func (s *Server) GetSlams(w http.ResponseWriter, r *http.Request) {
	slams, err := s.slamService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, slams)
}

func (s *Server) PostSlams(w http.ResponseWriter, r *http.Request) {
	var slam api.SlamRequest
	if err := ValidateJSON(r.Body, &slam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := s.slamService.Create(slam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusCreated, created)
}

func (s *Server) DeleteSlamsSlamID(w http.ResponseWriter, r *http.Request, slamID string) {
	if err := s.slamService.Delete(slamID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetSlamsSlamID(w http.ResponseWriter, r *http.Request, slamID string) {
	slam, err := s.slamService.GetByID(slamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	WriteJSON(w, http.StatusOK, slam)
}

func (s *Server) PutSlamsSlamID(w http.ResponseWriter, r *http.Request, slamID string) {
	var slam api.SlamRequest
	if err := ValidateJSON(r.Body, &slam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := s.slamService.Update(slamID, slam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, updated)
}

func (s *Server) GetSlamsSlamIDStages(w http.ResponseWriter, r *http.Request, slamID string) {
	stages, err := s.stageService.GetStages(slamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	WriteJSON(w, http.StatusOK, stages)
}

func (s *Server) PostSlamsSlamIDStages(w http.ResponseWriter, r *http.Request, slamID string) {
	var stage api.StageRequest
	if err := ValidateJSON(r.Body, &stage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := s.stageService.CreateStage(slamID, stage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusCreated, created)
}
