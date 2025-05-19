package dto

import (
	"github.com/google/uuid"
)

type InterviewerOutputDto struct {
	ID                                 uuid.UUID `json:"id"`
	Name                               string    `json:"name"`
	AvatarURL                          string    `json:"avatar_url"`
	EntryMessage                       string    `json:"entry_message"`
	CharacterDescription               string    `json:"character_description"`
	CharacterDescriptionTranslationKey string    `json:"character_description_translation_key"`
}
