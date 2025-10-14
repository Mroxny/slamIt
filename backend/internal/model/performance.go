package model

import "github.com/Mroxny/slamIt/internal/api"

type Performance struct {
	Model
	api.Performance `gorm:"embedded"`
}
