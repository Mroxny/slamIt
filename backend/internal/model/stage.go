package model

import "github.com/Mroxny/slamIt/internal/api"

type Stage struct {
	Model
	api.Stage `gorm:"embedded"`
}
