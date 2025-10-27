package api_test

import (
	"net/http"
	"testing"

	"github.com/Mroxny/slamIt/internal/router"
	"github.com/Mroxny/slamIt/internal/utils"
)

func TestVote(t *testing.T) {
	r := router.SetupTestRouter()
	perfId := utils.TestPerformances[0].Id

	tests := []utils.TestParams{
		{
			Name:     "get votes from valid performanceId",
			Method:   "GET",
			Url:      "/performances/" + perfId + "/votes",
			WantCode: http.StatusOK,
		},
		{
			Name:     "get votes with invalid pagination (should work)",
			Method:   "GET",
			Url:      "/performances/" + perfId + "/votes?page=-1",
			WantCode: http.StatusOK,
		},
		{
			Name:     "create valid performance",
			Method:   "POST",
			Body:     `{"deviceFingerprint":"123"}`,
			Auth:     true,
			Url:      "/performances/" + perfId + "/votes",
			WantCode: http.StatusCreated,
		},
	}

	utils.RunTests(t, tests, r, "alice@example.com", "P@ssw0rd")
}
