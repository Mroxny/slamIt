package handler

import (
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
)

func (s *Server) DeletePerformancesPerformanceID(w http.ResponseWriter, r *http.Request, performanceID string) {
	if err := s.perfService.DeletePerformance(r.Context(), performanceID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) GetPerformancesPerformanceID(w http.ResponseWriter, r *http.Request, performanceID string) {
	perf, err := s.perfService.GetPerformance(r.Context(), performanceID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	WriteJSON(w, http.StatusOK, perf)
}

func (s *Server) PutPerformancesPerformanceID(w http.ResponseWriter, r *http.Request, performanceID string) {
	var performance api.PerformanceUpdateRequest
	if err := ValidateJSON(r.Body, &performance); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := s.perfService.UpdatePerformance(r.Context(), performanceID, performance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusOK, updated)
}

func (s *Server) GetPerformancesPerformanceIDVotes(w http.ResponseWriter, r *http.Request, performanceID string, params api.GetPerformancesPerformanceIDVotesParams) {
	page, pageSize := ParsePageNumAndSize(params.Page, params.PageSize)
	votes, err := s.voteService.GetVotes(r.Context(), performanceID, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	WriteJSON(w, http.StatusOK, votes)
}

func (s *Server) PostPerformancesPerformanceIDVotes(w http.ResponseWriter, r *http.Request, performanceID string) {
	var vote api.VoteRequest
	if err := ValidateJSON(r.Body, &vote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := s.voteService.CreateVote(r.Context(), performanceID, vote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	WriteJSON(w, http.StatusCreated, created)
}
