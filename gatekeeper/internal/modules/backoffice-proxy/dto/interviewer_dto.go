package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type InterviewerInputDto struct {
	Name                               string `json:"name" validate:"required,min=5,max=30"`
	AvatarURL                          string `json:"avatar_url" validate:"required,url"`
	EntryMessage                       string `json:"entry_message" validate:"required,max=250"`
	CharacterDescription               string `json:"character_description" validate:"required,max=500"`
	CharacterDescriptionTranslationKey string `json:"character_description_translation_key" validate:"required,max=255"`
}

func (d *InterviewerInputDto) Validate() error {
	validate := validator.New()

	return validate.Struct(d)
}

type InterviewerOutputDto struct {
	ID                                 uuid.UUID `json:"id"`
	Name                               string    `json:"name"`
	AvatarURL                          string    `json:"avatar_url"`
	EntryMessage                       string    `json:"entry_message"`
	CharacterDescription               string    `json:"character_description"`
	CharacterDescriptionTranslationKey string    `json:"character_description_translation_key"`
	CreatedAt                          string    `json:"created_at"`
	UpdatedAt                          string    `json:"updated_at"`
}
