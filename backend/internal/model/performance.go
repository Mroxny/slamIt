package model

import "github.com/Mroxny/slamIt/internal/api"

type Performance struct {
	api.Performance `gorm:"embedded"`
	// ParticipationId string
	// StageId         string
	Votes []Vote
	Model
}
