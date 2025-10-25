package model

import (
	"errors"

	"github.com/Mroxny/slamIt/internal/api"
	"gorm.io/gorm"
)

type Participation struct {
	api.Participation
	// UserId string
	// SlamId string
	Model
}

func (p *Participation) BeforeDelete(tx *gorm.DB) (err error) {
	if p.Role == api.Creator {
		return errors.New("cannot remove the creator from a slam")
	}
	return nil
}
