package model

import "github.com/Mroxny/slamIt/internal/api"

type Stage struct {
	api.Stage `gorm:"embedded"`
	Model
}
