package api_test

import (
	"net/http"
	"testing"

	"github.com/Mroxny/slamIt/internal/router"
	"github.com/Mroxny/slamIt/internal/utils"
)

func TestPerformance(t *testing.T) {
	r := router.SetupTestRouter()
	stageId := utils.TestStages[0].Id
	perfId := utils.TestPerformances[0].Id
	partId1 := utils.TestParticipations[0].Id
	partId2 := utils.TestParticipations[1].Id

	tests := []utils.TestParams{
		{
			Name:     "get performances from valid stageId",
			Method:   "GET",
			Url:      "/stages/" + stageId + "/performances",
			WantCode: http.StatusOK,
		},
		{
			Name:     "get performances from invalid stageId",
			Method:   "GET",
			Url:      "/stages/xxx/performances",
			WantCode: http.StatusNotFound,
		},
		{
			Name:     "get performances with invalid pagination (should work)",
			Method:   "GET",
			Url:      "/stages/" + stageId + "/performances?page=-1",
			WantCode: http.StatusOK,
		},
		{
			Name:     "create valid performance",
			Method:   "POST",
			Body:     `{"participationId":"` + partId2 + `"}`,
			Auth:     true,
			Url:      "/stages/" + stageId + "/performances",
			WantCode: http.StatusCreated,
		},
		{
			Name:       "create invalid performance (already participating)",
			Method:     "POST",
			Body:       `{"participationId":"` + partId2 + `"}`,
			Auth:       true,
			UserName:   "bob@example.com",
			UserPasswd: "P@ssw0rd",
			Url:        "/stages/" + stageId + "/performances",
			WantCode:   http.StatusBadRequest,
		},
		{
			Name:       "update valid perfromance order",
			Method:     "PUT",
			Body:       `["` + partId2 + `","` + partId1 + `"]`,
			Auth:       true,
			UserName:   "bob@example.com",
			UserPasswd: "P@ssw0rd",
			Url:        "/stages/" + stageId + "/performances",
			WantCode:   http.StatusOK,
		},
		{
			Name:     "get valid performance",
			Method:   "GET",
			Url:      "/performances/" + perfId,
			WantCode: http.StatusOK,
		},
		{
			Name:     "get invalid performance",
			Method:   "GET",
			Url:      "/performances/xxx",
			WantCode: http.StatusNotFound,
		},
		{
			Name:     "update valid performance",
			Method:   "PUT",
			Body:     `{"details":"TEST PERFORMANCE UPDATED"}`,
			Auth:     true,
			Url:      "/performances/" + perfId,
			WantCode: http.StatusOK,
		},
		{
			Name:     "update invalid performance",
			Method:   "PUT",
			Body:     `{"participationId":"` + partId2 + `"}`,
			Auth:     true,
			Url:      "/performances/" + perfId,
			WantCode: http.StatusBadRequest,
		},
		{
			Name:     "delete valid performance",
			Method:   "DELETE",
			Auth:     true,
			Url:      "/performances/" + perfId,
			WantCode: http.StatusNoContent,
		},
		{
			Name:     "delete invalid performance (deleted before)",
			Method:   "DELETE",
			Auth:     true,
			Url:      "/performances/" + perfId,
			WantCode: http.StatusNotFound,
		},
	}

	utils.RunTests(t, tests, r, "alice@example.com", "P@ssw0rd")
}
