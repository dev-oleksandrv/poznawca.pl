package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
)

func MapInterviewModelToAppDto(interview *model.InterviewModel) *dto.AppInterviewDto {
	output := &dto.AppInterviewDto{
		ID:     interview.ID.String(),
		Status: string(interview.Status),
	}

	if interview.Interviewer != nil {
		output.Interviewer = MapInterviewerModelToAppDto(interview.Interviewer)
	}

	return output
}
