package handler

import (
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
)

func (s *Server) DeleteStagesStageID(w http.ResponseWriter, r *http.Request, stageID string) {
	if err := s.stageService.DeleteStage(r.Context(), stageID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetStagesStageID(w http.ResponseWriter, r *http.Request, stageID string) {
	stage, err := s.stageService.GetStage(r.Context(), stageID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	WriteJSON(w, http.StatusOK, stage)
}

func (s *Server) PutStagesStageID(w http.ResponseWriter, r *http.Request, stageID string) {
	var stage api.StageRequest
	if err := ValidateJSON(r.Body, &stage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := s.stageService.UpdateStage(r.Context(), stageID, stage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, updated)
}

func (s *Server) GetStagesStageIDPerformances(w http.ResponseWriter, r *http.Request, stageID string) {
	performances, err := s.perfService.GetPerformances(r.Context(), stageID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	WriteJSON(w, http.StatusOK, performances)
}

func (s *Server) PostStagesStageIDPerformances(w http.ResponseWriter, r *http.Request, stageID string) {
	var performance api.PerformanceRequest
	if err := ValidateJSON(r.Body, &performance); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := s.perfService.CreatePerformance(r.Context(), stageID, performance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusCreated, created)
}
