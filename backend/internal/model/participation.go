package model

import "github.com/Mroxny/slamIt/internal/api"

type Participation struct {
	api.Participation
	UserId string `gorm:"uniqueIndex:idx_user_slam;not null"`
	SlamId string `gorm:"uniqueIndex:idx_user_slam;not null"`
}
