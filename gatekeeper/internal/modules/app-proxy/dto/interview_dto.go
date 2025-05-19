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

type UpdateInterviewStatusInputDto struct {
	InterviewID uuid.UUID             `json:"interview_id" validate:"required,uuid"`
	Status      model.InterviewStatus `json:"status" validate:"required"`
}

func (d *UpdateInterviewStatusInputDto) Validate() error {
	validate := validator.New()

	return validate.Struct(d)
}

type ProcessInterviewClientMessageInputDto struct {
	InterviewID uuid.UUID `json:"interview_id" validate:"required,uuid"`
	ThreadID    string    `json:"thread_id" validate:"required"`
	Content     string    `json:"content" validate:"required"`
}

func (d *ProcessInterviewClientMessageInputDto) Validate() error {
	validate := validator.New()

	return validate.Struct(d)
}

type CreateInterviewInitialMessageInputDto struct {
	InterviewID uuid.UUID `json:"interview_id" validate:"required,uuid"`
	ContentText string    `json:"content_text" validate:"required"`
}

func (d *CreateInterviewInitialMessageInputDto) Validate() error {
	validate := validator.New()

	return validate.Struct(d)
}

type GenerateInterviewResultsInputDto struct {
	InterviewID uuid.UUID `json:"interview_id" validate:"required,uuid"`
	ThreadID    string    `json:"thread_id" validate:"required"`
}

type InterviewOutputDto struct {
	ID          uuid.UUID                 `json:"id"`
	Status      model.InterviewStatus     `json:"status"`
	Interviewer *InterviewerOutputDto     `json:"interviewer"`
	Result      *InterviewResultOutputDto `json:"result"`
	ThreadID    string                    `json:"thread_id"`
	CreatedAt   string                    `json:"created_at"`
	UpdatedAt   string                    `json:"updated_at"`
}
