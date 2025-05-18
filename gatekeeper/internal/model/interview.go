package model

import (
	"github.com/google/uuid"
	"time"
)

type Interview struct {
	ID       uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ThreadID string          `gorm:"type:varchar(255);not null"`
	Status   InterviewStatus `gorm:"type:varchar(50);not null;default:'pending'"`

	InterviewerID *uuid.UUID   `gorm:"type:uuid;not null;index"`
	Interviewer   *Interviewer `gorm:"foreignKey:InterviewerID;constraint:OnDelete:SET NULL;"`

	Messages []*InterviewMessage `gorm:"foreignKey:InterviewID;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
