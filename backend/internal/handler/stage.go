package handler

import (
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
)

func (s *Server) DeleteStagesStageID(w http.ResponseWriter, r *http.Request, stageID string) {
	if err := s.stageService.DeleteStage(stageID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) PutStagesStageID(w http.ResponseWriter, r *http.Request, stageID string) {
	var stage api.StageRequest
	if err := ValidateJSON(r.Body, &stage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := s.stageService.UpdateStage(stageID, stage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, updated)
}

func (s *Server) GetStagesStageIDPerformances(w http.ResponseWriter, r *http.Request, stageID string) {
	performances, err := s.perfService.GetPerformances(stageID)
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

	created, err := s.perfService.CreatePerformance(stageID, performance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusCreated, created)
}
