package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/Mroxny/slamIt/internal/api"
)

func GetAuthToken(r http.Handler, email, password string, createUser bool) (string, string) {
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
