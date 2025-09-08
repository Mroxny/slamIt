package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mroxny/slamIt/internal/router"
)

func TestSlamParticipationHandler(t *testing.T) {
	r := router.SetupTestRouter()

	tests := []struct {
		name     string
		method   string
		url      string
		wantCode int
	}{
		{
			name:     "join valid user+slam",
			method:   "POST",
			url:      "/participation/users/1/slams/1",
			wantCode: http.StatusCreated,
		},
		{
			name:     "join same slam twice (should fail)",
			method:   "POST",
			url:      "/participation/users/1/slams/1",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "join invalid user",
			method:   "POST",
			url:      "/participation/users/999/slams/1",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "join invalid slam",
			method:   "POST",
			url:      "/participation/users/1/slams/999",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "list slams for user",
			method:   "GET",
			url:      "/participation/users/1/slams",
			wantCode: http.StatusOK,
		},
		{
			name:     "list users for slam",
			method:   "GET",
			url:      "/participation/slams/1/users",
			wantCode: http.StatusOK,
		},
		{
			name:     "leave slam successfully",
			method:   "DELETE",
			url:      "/participation/users/1/slams/1",
			wantCode: http.StatusNoContent,
		},
		{
			name:     "leave slam again (not found)",
			method:   "DELETE",
			url:      "/participation/users/1/slams/1",
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tt.wantCode {
				t.Errorf("%s: got %d, want %d", tt.name, w.Code, tt.wantCode)
			}
		})
	}
}
