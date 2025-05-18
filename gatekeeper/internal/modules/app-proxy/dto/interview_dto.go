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

type ProcessClientMessageInputDto struct {
	InterviewID uuid.UUID `json:"interview_id" validate:"required,uuid"`
	ThreadID    string    `json:"thread_id" validate:"required"`
	Content     string    `json:"content" validate:"required"`
}

func (d *ProcessClientMessageInputDto) Validate() error {
	validate := validator.New()

	return validate.Struct(d)
}

type AIPromptInputDto struct {
	Content  string `json:"content" validate:"required"`
	Language string `json:"language" validate:"required"`
}

func (d *AIPromptInputDto) Validate() error {
	validate := validator.New()

	return validate.Struct(d)
}

type AIPromptOutputDto struct {
	Done             bool   `json:"done"`
	Content          string `json:"content"`
	Tips             string `json:"tips"`
	Translation      string `json:"translation"`
	GrammarScore     int    `json:"grammar_score"`
	GrammarFeedback  string `json:"grammar_feedback"`
	AccuracyScore    int    `json:"accuracy_score"`
	AccuracyFeedback string `json:"accuracy_feedback"`
	TotalScore       int    `json:"total_score"`
	TotalFeedback    string `json:"total_feedback"`
}

type InterviewOutputDto struct {
	ID          uuid.UUID             `json:"id"`
	Status      model.InterviewStatus `json:"status"`
	Interviewer *InterviewerOutputDto `json:"interviewer"`
	ThreadID    string                `json:"thread_id"`
	CreatedAt   string                `json:"created_at"`
	UpdatedAt   string                `json:"updated_at"`
}
