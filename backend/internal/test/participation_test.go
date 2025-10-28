package api_test

import (
	"net/http"
	"testing"

	"github.com/Mroxny/slamIt/internal/router"
	"github.com/Mroxny/slamIt/internal/utils"
)

func TestParticipation(t *testing.T) {
	r := router.SetupTestRouter()
	bobId := utils.TestUsers[0].Id
	aliceId := utils.TestUsers[1].Id
	johnId := utils.TestUsers[2].Id

	slamId := utils.TestSlams[2].Id

	tests := []utils.TestParams{
		{
			Name:     "join valid slam",
			Method:   "POST",
			Url:      "/participations/slams/" + slamId,
			Auth:     true,
			WantCode: http.StatusCreated,
		},
		{
			Name:     "join same slam twice (should fail)",
			Method:   "POST",
			Url:      "/participations/slams/" + slamId,
			Auth:     true,
			WantCode: http.StatusBadRequest,
		},
		{
			Name:     "join invalid slam",
			Method:   "POST",
			Url:      "/participations/slams/xxx",
			Auth:     true,
			WantCode: http.StatusBadRequest,
		},
		{
			Name:     "leave valid slam",
			Method:   "DELETE",
			Url:      "/participations/slams/" + slamId,
			Auth:     true,
			WantCode: http.StatusNoContent,
		},
		{
			Name:     "leave same slam twice (should fail)",
			Method:   "DELETE",
			Url:      "/participations/slams/" + slamId,
			Auth:     true,
			WantCode: http.StatusBadRequest,
		},
		{
			Name:       "leave same as creator (should fail)",
			Method:     "DELETE",
			UserName:   "alice@example.com",
			UserPasswd: "P@ssw0rd",
			Url:        "/participations/slams/" + slamId,
			Auth:       true,
			WantCode:   http.StatusBadRequest,
		},
		{
			Name:     "list slams for user",
			Method:   "GET",
			Url:      "/participations/users/" + bobId + "/slams",
			Auth:     true,
			WantCode: http.StatusOK,
		},
		{
			Name:     "list slams for user with invalid pagination (should work)",
			Method:   "GET",
			Url:      "/participations/users/" + bobId + "/slams?page=-1",
			Auth:     true,
			WantCode: http.StatusOK,
		},
		{
			Name:     "list users for slam",
			Method:   "GET",
			Url:      "/participations/slams/" + slamId + "/users",
			Auth:     true,
			WantCode: http.StatusOK,
		},
		{
			Name:     "list users for slam with invalid pagination (should work)",
			Method:   "GET",
			Url:      "/participations/slams/" + slamId + "/users?page=-1",
			Auth:     true,
			WantCode: http.StatusOK,
		},
		{
			Name:     "add valid performer",
			Method:   "POST",
			Url:      "/participations/slams/" + slamId + "/users",
			Body:     `{"userId":"` + bobId + `","role": "performer"}`,
			Auth:     true,
			WantCode: http.StatusCreated,
		},
		{
			Name:     "add invalid user (invalid user id)",
			Method:   "POST",
			Url:      "/participations/slams/" + slamId + "/users",
			Body:     `{"userId":"xxx","role": "performer"}`,
			Auth:     true,
			WantCode: http.StatusBadRequest,
		},
		{
			Name:     "add invalid user (invalid role)",
			Method:   "POST",
			Url:      "/participations/slams/" + slamId + "/users",
			Body:     `{"userId":"xxx","role": "INVALID"}`,
			Auth:     true,
			WantCode: http.StatusBadRequest,
		},
		{
			Name:     "add valid user to invalid slam (should fail)",
			Method:   "POST",
			Url:      "/participations/slams/xxx/users",
			Body:     `{"userId":"` + bobId + `"}`,
			Auth:     true,
			WantCode: http.StatusBadRequest,
		},
		{
			Name:     "update valid participation",
			Method:   "PUT",
			Url:      "/participations/slams/" + slamId + "/users/" + bobId,
			Body:     `{"role":"performer"}`,
			Auth:     true,
			WantCode: http.StatusOK,
		},
		{
			Name:     "update invalid participation (invalid role)",
			Method:   "PUT",
			Url:      "/participations/slams/" + slamId + "/users/" + bobId,
			Body:     `{"role":"INVALID"}`,
			Auth:     true,
			WantCode: http.StatusBadRequest,
		},
		{
			Name:     "remove user from slam successfully",
			Method:   "DELETE",
			Url:      "/participations/slams/" + slamId + "/users/" + johnId,
			Auth:     true,
			WantCode: http.StatusNoContent,
		},
		{
			Name:     "remove user from slam again",
			Method:   "DELETE",
			Url:      "/participations/slams/" + slamId + "/users/" + johnId,
			Auth:     true,
			WantCode: http.StatusBadRequest,
		},
		{
			Name:     "remove creator from slam (should fail)",
			Method:   "DELETE",
			Url:      "/participations/slams/" + slamId + "/users/" + aliceId,
			Auth:     true,
			WantCode: http.StatusBadRequest,
		},
	}

	utils.RunTests(t, tests, r, "bob@example.com", "P@ssw0rd")

}
