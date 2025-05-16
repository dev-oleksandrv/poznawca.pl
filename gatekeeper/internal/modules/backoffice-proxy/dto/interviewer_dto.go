package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type InterviewerInputDto struct {
	Name         string `json:"name" validate:"required,min=5,max=30"`
	AvatarURL    string `json:"avatar_url" validate:"required,url"`
	EntryMessage string `json:"entry_message" validate:"required,max=250"`
}

func (d *InterviewerInputDto) Validate() error {
	validate := validator.New()

	return validate.Struct(d)
}

type InterviewerOutputDto struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	AvatarURL    string    `json:"avatar_url"`
	EntryMessage string    `json:"entry_message"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
}
