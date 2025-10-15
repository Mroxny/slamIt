package model

import "github.com/Mroxny/slamIt/internal/api"

type User struct {
	api.User   `gorm:"embedded"`
	PasswdHash string `gorm:"not null"`
	Model
}
