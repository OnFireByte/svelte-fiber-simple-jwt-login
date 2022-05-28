package models

import (
	"time"

	"gorm.io/gorm"
)

type OwnModel struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"; psql:"index"`
}