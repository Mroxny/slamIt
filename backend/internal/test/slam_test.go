package api_test

import (
	"net/http"
	"testing"

	"github.com/Mroxny/slamIt/internal/router"
	"github.com/Mroxny/slamIt/internal/utils"
)

func TestSlam(t *testing.T) {
	r := router.SetupTestRouter()
	slamId := utils.TestSlams[0].Id

	tests := []utils.TestParams{
		{
			Name:     "create valid slam",
			Method:   "POST",
			Url:      "/slams",
			Body:     `{"title":"Test Slam","description":"Testing slam endpoint", "public":true}`,
			Auth:     true,
			WantCode: http.StatusCreated,
		},
		{
			Name:     "create invalid slam (missing title)",
			Method:   "POST",
			Url:      "/slams",
			Body:     `{"description":"Charlie"}`,
			Auth:     true,
			WantCode: http.StatusBadRequest,
		},
		{
			Name:     "get all public slams",
			Method:   "GET",
			Url:      "/slams",
			WantCode: http.StatusOK,
		},
		{
			Name:     "get all slams with invalid pagination (should work)",
			Method:   "GET",
			Url:      "/slams?page=-1",
			WantCode: http.StatusOK,
		},
		{
			Name:     "get existing slam",
			Method:   "GET",
			Url:      "/slams/" + slamId,
			Auth:     true,
			WantCode: http.StatusOK,
		},
		{
			Name:     "get non-existing slam",
			Method:   "GET",
			Url:      "/slams/999999999",
			Auth:     true,
			WantCode: http.StatusNotFound,
		},
		{
			Name:     "update slam valid",
			Method:   "PUT",
			Url:      "/slams/" + slamId,
			Body:     `{"title":"Slam Updated","description":"Updated description", "public":false}`,
			Auth:     true,
			WantCode: http.StatusOK,
		},
		{
			Name:     "delete slam valid",
			Method:   "DELETE",
			Url:      "/slams/" + slamId,
			Auth:     true,
			WantCode: http.StatusNoContent,
		},
		{
			Name:     "delete slam invalid (deleted before)",
			Method:   "DELETE",
			Url:      "/slams/" + slamId,
			Auth:     true,
			WantCode: http.StatusNotFound,
		},
	}

	utils.RunTests(t, tests, r, "bob@example.com", "P@ssw0rd")
}
