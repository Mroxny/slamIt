package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mroxny/slamIt/internal/router"
	"github.com/Mroxny/slamIt/internal/utils"
)

func TestUserHandler(t *testing.T) {
	r := router.SetupTestRouter()
	uId, token := utils.GetAuthToken(r, "bob@example.com", "P@ssw0rd", false)

	tests := []struct {
		name     string
		method   string
		url      string
		body     string
		auth     bool
		wantCode int
	}{
		{
			name:     "get all users",
			method:   "GET",
			url:      "/users/",
			auth:     true,
			wantCode: http.StatusOK,
		},
		{
			name:     "get existing user",
			method:   "GET",
			url:      "/users/" + uId,
			auth:     true,
			wantCode: http.StatusOK,
		},
		{
			name:     "get non-existing user",
			method:   "GET",
			url:      "/users/xxx",
			auth:     true,
			wantCode: http.StatusNotFound,
		},
		{
			name:     "update user valid",
			method:   "PUT",
			url:      "/users/" + uId,
			body:     `{"name":"Alice Updated","email":"alice@new.com"}`,
			auth:     true,
			wantCode: http.StatusOK,
		},
		{
			name:     "delete user valid",
			method:   "DELETE",
			url:      "/users/" + uId,
			auth:     true,
			wantCode: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			if tt.auth {
				req.Header.Set("Authorization", "Bearer "+token)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("%s: got %d, want %d", tt.name, w.Code, tt.wantCode)
			}
		})
	}
}
