package dto

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type AppInterviewDto struct {
	ID          string             `json:"id"`
	Status      string             `json:"status"`
	Interviewer *AppInterviewerDto `json:"interviewer,omitempty"`
}

type CreateAppInterviewRequestDto struct {
	InterviewerID *string `json:"interviewer_id,omitempty"`
}

func (d *CreateAppInterviewRequestDto) Validate() error {
	return validate.Struct(d)
}
