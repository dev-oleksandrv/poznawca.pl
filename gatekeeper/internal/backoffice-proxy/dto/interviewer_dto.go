package dto

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type BackofficeInterviewerDto struct {
	ID                        string `json:"id"`
	Name                      string `json:"name"`
	AvatarURL                 string `json:"avatar_url"`
	EntryMessage              string `json:"entry_message"`
	Description               string `json:"description"`
	DescriptionTranslationKey string `json:"description_translation_key"`
	CreatedAt                 string `json:"created_at"`
	UpdatedAt                 string `json:"updated_at"`
}

type CreateBackofficeInterviewerRequestDto struct {
	Name                      string `json:"name" validate:"required"`
	AvatarURL                 string `json:"avatar_url" validate:"required,url"`
	EntryMessage              string `json:"entry_message" validate:"required"`
	Description               string `json:"description" validate:"required"`
	DescriptionTranslationKey string `json:"description_translation_key" validate:"required"`
}

func (d *CreateBackofficeInterviewerRequestDto) Validate() error {
	return validate.Struct(d)
}

type UpdateBackofficeInterviewerRequestDto struct {
	Name                      *string `json:"name" validate:"omitempty"`
	AvatarURL                 *string `json:"avatar_url" validate:"omitempty,url"`
	EntryMessage              *string `json:"entry_message" validate:"omitempty"`
	Description               *string `json:"description" validate:"omitempty"`
	DescriptionTranslationKey *string `json:"description_translation_key" validate:"omitempty"`
}

func (d *UpdateBackofficeInterviewerRequestDto) Validate() error {
	return validate.Struct(d)
}
