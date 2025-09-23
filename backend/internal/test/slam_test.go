package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/router"
	"github.com/Mroxny/slamIt/internal/utils"
)

func TestSlam(t *testing.T) {
	r := router.SetupTestRouter()
	_, token := utils.GetAuthToken(r, "bob@example.com", "P@ssw0rd", false)
	slamId := "3"

	tests := []struct {
		name     string
		method   string
		url      string
		body     string
		auth     bool
		wantCode int
	}{
		{
			name:     "create valid slam",
			method:   "POST",
			url:      "/slams",
			body:     `{"title":"Test Slam","description":"Testing slam endpoint", "public":true}`,
			auth:     true,
			wantCode: http.StatusCreated,
		},
		{
			name:     "create invalid slam (missing title)",
			method:   "POST",
			url:      "/slams",
			body:     `{"description":"Charlie"}`,
			auth:     true,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "get all public slams",
			method:   "GET",
			url:      "/slams",
			wantCode: http.StatusOK,
		},
		{
			name:     "get existing slam",
			method:   "GET",
			url:      "/slams/" + slamId,
			auth:     true,
			wantCode: http.StatusOK,
		},
		{
			name:     "get non-existing slam",
			method:   "GET",
			url:      "/slams/999999999",
			auth:     true,
			wantCode: http.StatusNotFound,
		},
		{
			name:     "update slam valid",
			method:   "PUT",
			url:      "/slams/" + slamId,
			body:     `{"title":"Slam Updated","description":"Updated description", "public":false}`,
			auth:     true,
			wantCode: http.StatusOK,
		},
		{
			name:     "delete slam valid",
			method:   "DELETE",
			url:      "/slams/" + slamId,
			auth:     true,
			wantCode: http.StatusNoContent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, api.ServerUrlDev+tt.url, strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			if tt.auth {
				req.Header.Set("Authorization", "Bearer "+token)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("%s: got %d, want %d, msg: %s", tt.name, w.Code, tt.wantCode, w.Body)
			}
		})
	}
}
