package api_test

import (
	"net/http"
	"testing"

	"github.com/Mroxny/slamIt/internal/router"
	"github.com/Mroxny/slamIt/internal/utils"
)

func TestUser(t *testing.T) {
	r := router.SetupTestRouter()
	uId := utils.TestUsers[0].Id

	tests := []utils.TestParams{
		{
			Name:     "get all users",
			Method:   "GET",
			Url:      "/users",
			Auth:     true,
			WantCode: http.StatusOK,
		},
		{
			Name:     "get all users with invalid pagination (should work)",
			Method:   "GET",
			Url:      "/users?page=-1",
			Auth:     true,
			WantCode: http.StatusOK,
		},
		{
			Name:     "create temporary user",
			Method:   "POST",
			Url:      "/users",
			Body:     `{"name":"TemporaryBob"}`,
			Auth:     true,
			WantCode: http.StatusCreated,
		},
		{
			Name:     "get existing user",
			Method:   "GET",
			Url:      "/users/" + uId,
			Auth:     true,
			WantCode: http.StatusOK,
		},
		{
			Name:     "get non-existing user",
			Method:   "GET",
			Url:      "/users/xxx",
			Auth:     true,
			WantCode: http.StatusNotFound,
		},
		{
			Name:     "update user valid",
			Method:   "PUT",
			Url:      "/users/" + uId,
			Body:     `{"name":"Alice Updated","email":"alice@new.com"}`,
			Auth:     true,
			WantCode: http.StatusOK,
		},
		{
			Name:     "delete valid user",
			Method:   "DELETE",
			Url:      "/users/" + uId,
			Auth:     true,
			WantCode: http.StatusNoContent,
		},
		{
			Name:     "delete invalid user (deleted before)",
			Method:   "DELETE",
			Url:      "/users/" + uId,
			Auth:     true,
			WantCode: http.StatusNotFound,
		},
	}

	utils.RunTests(t, tests, r, "bob@example.com", "P@ssw0rd")
}
