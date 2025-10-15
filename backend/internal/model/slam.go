package model

import "github.com/Mroxny/slamIt/internal/api"

type Slam struct {
	api.Slam `gorm:"embedded"`
	Model
}
