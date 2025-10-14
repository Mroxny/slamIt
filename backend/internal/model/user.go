package model

import "github.com/Mroxny/slamIt/internal/api"

type User struct {
	Model
	api.User   `gorm:"embedded"`
	PasswdHash string `gorm:"not null"`
}
