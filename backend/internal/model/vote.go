package model

import "github.com/Mroxny/slamIt/internal/api"

type Vote struct {
	api.Vote `gorm:"embedded"`
	Model
}
