package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mroxny/slamIt/internal/router"
)

func TestUserHandler(t *testing.T) {
	r := router.SetupTestRouter()

	tests := []struct {
		name     string
		method   string
		url      string
		body     string
		wantCode int
	}{
		{
			name:     "create valid user",
			method:   "POST",
			url:      "/users/",
			body:     `{"name":"Bob","email":"bob@example.com"}`,
			wantCode: http.StatusCreated,
		},
		{
			name:     "create invalid user (missing email)",
			method:   "POST",
			url:      "/users/",
			body:     `{"name":"Charlie"}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "get existing user",
			method:   "GET",
			url:      "/users/1",
			wantCode: http.StatusOK,
		},
		{
			name:     "get non-existing user",
			method:   "GET",
			url:      "/users/999",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "update user valid",
			method:   "PUT",
			url:      "/users/1",
			body:     `{"name":"Alice Updated","email":"alice@new.com"}`,
			wantCode: http.StatusOK,
		},
		{
			name:     "delete user valid",
			method:   "DELETE",
			url:      "/users/1",
			wantCode: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("%s: got %d, want %d", tt.name, w.Code, tt.wantCode)
			}
		})
	}
}
