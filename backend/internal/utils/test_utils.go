package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mroxny/slamIt/internal/api"
	"github.com/Mroxny/slamIt/internal/model"
	"github.com/go-chi/chi/v5"
)

func getAuthToken(r http.Handler, email, password string, createUser bool) (string, string) {
	if createUser {
		registerBody := api.RegisterRequest{
			Name:     "TestUser",
			Email:    email,
			Password: password,
		}
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(registerBody); err != nil {
			panic("failed to encode register body: " + err.Error())
		}

		req := httptest.NewRequest("POST", api.ServerUrlDev+"/auth/register", buf)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusCreated && w.Code != http.StatusConflict {
			panic(fmt.Sprintf("failed to register user in test: %d", w.Code))
		}
	}

	loginBody := api.LoginRequest{
		Email:    email,
		Password: password,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(loginBody); err != nil {
		panic("failed to encode login body: " + err.Error())
	}

	req := httptest.NewRequest("POST", api.ServerUrlDev+"/auth/login", buf)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		fmt.Println("login failed, response body:", w.Body.String())
		panic("failed to login user in test")
	}

	var resp api.LoginResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		panic("failed to parse login response in test: " + err.Error())
	}

	return *resp.UserId, *resp.Token
}

func RunTests(t *testing.T, tests []TestParams, r *chi.Mux, defaultUserName string, defaultUserPasswd string) {
	if defaultUserName == "" {
		defaultUserName = "bob@example.com"
	}
	if defaultUserPasswd == "" {
		defaultUserPasswd = "P@ssw0rd"
	}
	_, token := getAuthToken(r, defaultUserName, defaultUserPasswd, false)

	for _, tt := range tests {
		var tmpToken string
		if tt.UserName != "" && tt.UserPasswd != "" {
			_, tmpToken = getAuthToken(r, tt.UserName, tt.UserPasswd, false)
		} else {
			tmpToken = token
		}
		t.Run(tt.Name, func(t *testing.T) {
			req := httptest.NewRequest(tt.Method, api.ServerUrlDev+tt.Url, strings.NewReader(tt.Body))
			req.Header.Set("Content-Type", "application/json")
			if tt.Auth {
				req.Header.Set("Authorization", "Bearer "+tmpToken)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != tt.WantCode {
				t.Errorf("%s: got %d, want %d, msg: %s", tt.Name, w.Code, tt.WantCode, w.Body)
			}
		})
	}
}

type TestParams struct {
	Name       string
	Method     string
	Url        string
	Body       string
	Auth       bool
	WantCode   int
	UserName   string
	UserPasswd string
}

var slamTitle = "Poetry Night"
var slamDescription = "Evening of poems"
var stageDetails = "TEST STAGE"

var TestSlams = []model.Slam{
	{
		Slam: api.Slam{
			Id:          "1b338aa8-74a1-43e9-8034-94f144e77c3a",
			Title:       slamTitle,
			Description: &slamDescription,
			Public:      true,
		},
	},
	{
		Slam: api.Slam{
			Id:          "85bf4f72-3cd2-46df-8d37-016442f150f7",
			Title:       slamTitle + " 2",
			Description: &slamDescription,
			Public:      false,
		},
	},
	{
		Slam: api.Slam{
			Id:          "6340ffb4-eb8f-4452-a1dd-19d1d4806286",
			Title:       slamTitle + " 3",
			Description: &slamDescription,
			Public:      true,
		},
	},
}
var TestUsers = []model.User{
	{
		User: api.User{
			Id:    "24222e63-d545-4fdb-9f74-8782e17fe9d1",
			Email: "bob@example.com",
			Name:  "Bob",
		},
		PasswdHash: "$2a$10$pUWBo2E1AWC8uBEYFHEzyujAwIZIrmGH1eot6NSous/2fz7IYI/BW",
	},
	{
		User: api.User{
			Id:    "d5106479-33d7-4fe4-b1c1-eb75360adfa5",
			Email: "alice@example.com",
			Name:  "Alice",
		},
		PasswdHash: "$2a$10$pUWBo2E1AWC8uBEYFHEzyujAwIZIrmGH1eot6NSous/2fz7IYI/BW",
	},
	{
		User: api.User{
			Id:    "36fddebd-d325-4c81-b367-e0dfb1606c03",
			Email: "john@example.com",
			Name:  "John",
		},
		PasswdHash: "$2a$10$pUWBo2E1AWC8uBEYFHEzyujAwIZIrmGH1eot6NSous/2fz7IYI/BW",
	},
}
var TestParticipations = []model.Participation{
	{
		Participation: api.Participation{
			Id:     "194b5238-5a88-4a59-96f6-f81ba39dad57",
			Role:   api.Performer,
			UserId: TestUsers[1].Id,
			SlamId: TestSlams[0].Id,
		},
	},
	{
		Participation: api.Participation{
			Id:     "ff9bae50-196e-4a22-8280-7fee604af3ab",
			Role:   api.Performer,
			UserId: TestUsers[2].Id,
			SlamId: TestSlams[0].Id,
		},
	},
	{
		Participation: api.Participation{
			Id:     "ef6bc88c-986f-4dd3-887c-aa7c5ad26165",
			Role:   api.Performer,
			UserId: TestUsers[2].Id,
			SlamId: TestSlams[2].Id,
		},
	},
}
var TestStages = []model.Stage{
	{
		Stage: api.Stage{
			Id:        "1b6e7682-b374-4115-a8fb-420e8f7c626a",
			SlamId:    TestSlams[0].Id,
			StageType: api.Simple,
			Details:   &stageDetails,
		},
	},
}
var TestPerformances = []model.Performance{
	{
		Performance: api.Performance{
			Id:              "c0af34c4-0ec3-41f7-ac75-e8e097efb8a8",
			ParticipationId: TestParticipations[0].Id,
			StageId:         TestStages[0].Id,
		},
	},
}
