package model

type Slam struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
}
