package handler

import (
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
)

func (s *Server) GetSlams(w http.ResponseWriter, r *http.Request, params api.GetSlamsParams) {
	page, pageSize := ParsePageNumAndSize(params.Page, params.PageSize)
	slams, err := s.slamService.GetAll(r.Context(), page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, slams)
}

func (s *Server) PostSlams(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var slam api.SlamRequest
	if err := ValidateJSON(r.Body, &slam); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := s.slamService.Create(r.Context(), slam, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusCreated, created)
}

func (s *Server) DeleteSlamsSlamID(w http.ResponseWriter, r *http.Request, slamID string) {
	if err := s.slamService.Delete(r.Context(), slamID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetSlamsSlamID(w http.ResponseWriter, r *http.Request, slamID string) {
	slam, err := s.slamService.GetByID(r.Context(), slamID)
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

	updated, err := s.slamService.Update(r.Context(), slamID, slam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, updated)
}

func (s *Server) GetSlamsSlamIDStages(w http.ResponseWriter, r *http.Request, slamID string, params api.GetSlamsSlamIDStagesParams) {
	page, pageSize := ParsePageNumAndSize(params.Page, params.PageSize)
	stages, err := s.stageService.GetStages(r.Context(), slamID, page, pageSize)
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

	created, err := s.stageService.CreateStage(r.Context(), slamID, stage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusCreated, created)
}
