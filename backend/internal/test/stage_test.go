package api_test

import (
	"net/http"
	"testing"

	"github.com/Mroxny/slamIt/internal/router"
	"github.com/Mroxny/slamIt/internal/utils"
)

func TestStage(t *testing.T) {
	r := router.SetupTestRouter()
	slamId := utils.TestSlams[0].Id
	stageId := utils.TestStages[0].Id

	tests := []utils.TestParams{
		{
			Name:     "get stages from valid slamId",
			Method:   "GET",
			Url:      "/slams/" + slamId + "/stages",
			WantCode: http.StatusOK,
		},
		{
			Name:     "get stages with invalid pagination (should work)",
			Method:   "GET",
			Url:      "/slams/" + slamId + "/stages?page=-1",
			WantCode: http.StatusOK,
		},
		{
			Name:     "get stages from invalid slamId",
			Method:   "GET",
			Url:      "/slams/xxx/stages",
			WantCode: http.StatusNotFound,
		},
		{
			Name:     "create valid stage",
			Method:   "POST",
			Body:     `{"stageType":"simple"}`,
			Auth:     true,
			Url:      "/slams/" + slamId + "/stages",
			WantCode: http.StatusCreated,
		},
		{
			Name:     "create invalid stage (invalid stage type)",
			Method:   "POST",
			Body:     `{"stageType":"INVALID"}`,
			Auth:     true,
			Url:      "/slams/xxx/stages",
			WantCode: http.StatusBadRequest,
		},
		{
			Name:     "get valid stage",
			Method:   "GET",
			Url:      "/stages/" + stageId,
			WantCode: http.StatusOK,
		},
		{
			Name:     "get invalid stage",
			Method:   "GET",
			Url:      "/stages/xxx",
			WantCode: http.StatusNotFound,
		},
		{
			Name:     "update valid stage",
			Method:   "PUT",
			Body:     `{"stageType":"battle","round":1,"details":"Updated stage"}`,
			Auth:     true,
			Url:      "/stages/" + stageId,
			WantCode: http.StatusOK,
		},
		{
			Name:     "update invalid stage (wrong round type)",
			Method:   "PUT",
			Body:     `{"stageType":"battle","round":"test","details":"Updated stage"}`,
			Auth:     true,
			Url:      "/stages/" + stageId,
			WantCode: http.StatusBadRequest,
		},
		{
			Name:     "delete valid stage",
			Method:   "DELETE",
			Auth:     true,
			Url:      "/stages/" + stageId,
			WantCode: http.StatusNoContent,
		},
		{
			Name:     "delete invalid stage (deleted before)",
			Method:   "DELETE",
			Auth:     true,
			Url:      "/stages/" + stageId,
			WantCode: http.StatusNotFound,
		},
	}

	utils.RunTests(t, tests, r, "bob@example.com", "P@ssw0rd")
}
