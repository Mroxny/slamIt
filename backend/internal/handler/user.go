package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/utils"
)

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := s.userService.GetAll()
	json.NewEncoder(w).Encode(users)
}

func (s *Server) GetUsersId(w http.ResponseWriter, r *http.Request, id string) {
	user, err := s.userService.GetByID(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (s *Server) PutUsersId(w http.ResponseWriter, r *http.Request, id string) {
	var u api.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(u); err != nil {
		http.Error(w, "invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := s.userService.Update(id, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (s *Server) DeleteUsersId(w http.ResponseWriter, r *http.Request, id string) {
	if err := s.userService.Delete(id); err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
