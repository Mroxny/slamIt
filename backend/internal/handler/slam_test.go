package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mroxny/slamIt/internal/router"
)

func TestSlamHandler(t *testing.T) {
	r := router.SetupTestRouter()

	tests := []struct {
		name     string
		method   string
		url      string
		body     string
		wantCode int
	}{
		{
			name:     "create valid slam",
			method:   "POST",
			url:      "/slams/",
			body:     `{"title":"Test Slam","description":"Testing slam endpoint"}`,
			wantCode: http.StatusCreated,
		},
		{
			name:     "create invalid slam (missing title)",
			method:   "POST",
			url:      "/slams/",
			body:     `{"descriptiom":"Charlie"}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "get existing slam",
			method:   "GET",
			url:      "/slams/1",
			wantCode: http.StatusOK,
		},
		{
			name:     "get non-existing slam",
			method:   "GET",
			url:      "/slams/999",
			wantCode: http.StatusNotFound,
		},
		{
			name:     "update slam valid",
			method:   "PUT",
			url:      "/slams/1",
			body:     `{"title":"Slam Updated","description":"Updated description"}`,
			wantCode: http.StatusOK,
		},
		{
			name:     "delete slam valid",
			method:   "DELETE",
			url:      "/slams/1",
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
