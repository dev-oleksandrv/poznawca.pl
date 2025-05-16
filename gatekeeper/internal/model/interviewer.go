package model

import (
	"github.com/google/uuid"
	"time"
)

type Interviewer struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string    `gorm:"type:varchar(255);not null"`
	AvatarURL    string    `gorm:"type:varchar(255);not null"`
	EntryMessage string    `gorm:"type:text;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
