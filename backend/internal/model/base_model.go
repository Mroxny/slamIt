package model

import "gorm.io/gorm"

type Model struct {
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
