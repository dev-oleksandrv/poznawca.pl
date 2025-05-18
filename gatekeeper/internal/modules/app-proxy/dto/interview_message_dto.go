package dto

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/google/uuid"
)

type InterviewMessageOutputDto struct {
	ID                     uuid.UUID                  `json:"id"`
	ContentText            string                     `json:"content_text"`
	ContentTranslationText string                     `json:"content_translation_text"`
	TipsText               string                     `json:"tips_text"`
	Role                   model.InterviewMessageRole `json:"role"`
}
