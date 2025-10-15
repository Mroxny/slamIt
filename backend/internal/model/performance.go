package model

import "github.com/Mroxny/slamIt/internal/api"

type Performance struct {
	api.Performance `gorm:"embedded"`
	Votes           []Vote
	Model
}
