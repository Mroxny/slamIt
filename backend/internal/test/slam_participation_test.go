package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/router"
	"github.com/Mroxny/slamIt/internal/utils"
)

func TestSlamParticipation(t *testing.T) {
	r := router.SetupTestRouter()
	uId, token := utils.GetAuthToken(r, "bob@example.com", "P@ssw0rd", false)
	slamId := "2"

	tests := []struct {
		name     string
		method   string
		url      string
		auth     bool
		wantCode int
	}{
		{
			name:     "join valid user+slam",
			method:   "POST",
			url:      "/participation/users/" + uId + "/slams/" + slamId,
			auth:     true,
			wantCode: http.StatusCreated,
		},
		{
			name:     "join same slam twice (should fail)",
			method:   "POST",
			url:      "/participation/users/" + uId + "/slams/" + slamId,
			auth:     true,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "join invalid user",
			method:   "POST",
			url:      "/participation/users/xxx/slams/" + slamId,
			auth:     true,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "join invalid slam",
			method:   "POST",
			url:      "/participation/users/" + uId + "/slams/999999999",
			auth:     true,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "list slams for user",
			method:   "GET",
			url:      "/participation/users/" + uId + "/slams",
			auth:     true,
			wantCode: http.StatusOK,
		},
		{
			name:     "list users for slam",
			method:   "GET",
			url:      "/participation/slams/" + slamId + "/users",
			auth:     true,
			wantCode: http.StatusOK,
		},
		{
			name:     "leave slam successfully",
			method:   "DELETE",
			url:      "/participation/users/" + uId + "/slams/" + slamId,
			auth:     true,
			wantCode: http.StatusNoContent,
		},
		{
			name:     "leave slam again (not found)",
			method:   "DELETE",
			url:      "/participation/users/" + uId + "/slams/" + slamId,
			auth:     true,
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, api.ServerUrlDev+tt.url, nil)
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
