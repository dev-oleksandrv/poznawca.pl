package dto

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateInterviewInputDto struct {
	InterviewerID uuid.UUID `json:"interviewer_id" validate:"required,uuid"`
}

func (d *CreateInterviewInputDto) Validate() error {
	validate := validator.New()

	return validate.Struct(d)
}

type InterviewOutputDto struct {
	ID          uuid.UUID             `json:"id"`
	Status      model.InterviewStatus `json:"status"`
	Interviewer *InterviewerOutputDto `json:"interviewer"`
	CreatedAt   string                `json:"created_at"`
	UpdatedAt   string                `json:"updated_at"`
}
