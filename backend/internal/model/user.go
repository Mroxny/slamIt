package model

type User struct {
	Id           string `json:"id"`
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	PasswordHash string `json:"-"`
}
