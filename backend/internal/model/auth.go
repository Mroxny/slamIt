package model

type RegisterRequest struct {
	Name     string `json:"name" example:"Bob"`
	Email    string `json:"email" example:"bob@example.com"`
	Password string `json:"password" example:"P@ssw0rd"`
}

type LoginRequest struct {
	Email    string `json:"email" example:"bob@example.com"`
	Password string `json:"password" example:"P@ssw0rd"`
}

type LoginResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
