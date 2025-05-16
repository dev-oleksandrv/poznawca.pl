package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/pkg/date"
)

func MapInterviewModelToOutput(interview *model.Interview) *dto.InterviewOutputDto {
	output := &dto.InterviewOutputDto{
		ID:        interview.ID,
		Status:    interview.Status,
		CreatedAt: date.FormatTime(interview.CreatedAt),
		UpdatedAt: date.FormatTime(interview.UpdatedAt),
	}

	if interview.Interviewer != nil {
		output.Interviewer = MapInterviewerModelToOutput(interview.Interviewer)
	}

	return output
}
