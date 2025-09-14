package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

func GetAuthToken(r http.Handler, email, password string, createUser bool) (string, string) {

	if createUser {
		registerBody := fmt.Sprintf(`{"name":"%s","email":"%s","password":"%s"}`, "TestUser", email, password)
		req := httptest.NewRequest("POST", "/auth/register", strings.NewReader(registerBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusCreated && w.Code != http.StatusConflict {
			panic("failed to register user in test")
		}
	}

	loginBody := fmt.Sprintf(`{"email":"%s","password":"%s"}`, email, password)
	req := httptest.NewRequest("POST", "/auth/login", strings.NewReader(loginBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		fmt.Println(w.Body)
		panic("failed to login user in test")
	}

	var resp struct {
		ID    string `json:"id"`
		Token string `json:"token"`
	}
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		panic("failed to parse login response in test")
	}

	return resp.ID, resp.Token
}
