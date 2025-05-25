package model

import (
	"github.com/google/uuid"
	"time"
)

type InterviewMessage struct {
	ID                     uuid.UUID            `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ContentText            string               `gorm:"type:text;not null"`
	ContentTranslationText string               `gorm:"type:text;not null;default:''"`
	TipsText               string               `gorm:"type:text;not null;default:''"`
	Role                   InterviewMessageRole `gorm:"type:varchar(50);not null"`
	Type                   InterviewMessageType `gorm:"type:varchar(50);not null;default:'default'"`
	IsLastMessage          bool                 `gorm:"type:boolean;not null;default:false"`
	InterviewID            uuid.UUID            `gorm:"type:uuid;not null;index"`
	Interview              *Interview           `gorm:"foreignKey:InterviewID;constraint:OnDelete:CASCADE;"`
	CreatedAt              time.Time            `gorm:"autoCreateTime"`
	UpdatedAt              time.Time            `gorm:"autoUpdateTime"`
}
