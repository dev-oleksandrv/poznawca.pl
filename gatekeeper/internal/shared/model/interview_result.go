package model

import (
	"github.com/google/uuid"
	"time"
)

type InterviewResult struct {
	ID               uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	GrammarScore     int        `gorm:"type:int;not null"`
	AccuracyScore    int        `gorm:"type:int;not null"`
	TotalScore       int        `gorm:"type:int;not null"`
	GrammarFeedback  string     `gorm:"type:text;not null"`
	AccuracyFeedback string     `gorm:"type:text;not null"`
	TotalFeedback    string     `gorm:"type:text;not null"`
	InterviewID      uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex"`
	Interview        *Interview `gorm:"foreignKey:InterviewID"`
	CreatedAt        time.Time  `gorm:"autoCreateTime"`
	UpdatedAt        time.Time  `gorm:"autoUpdateTime"`
}
