package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
)

func MapInterviewModelToAppDto(interview *model.Interview) *dto.AppInterviewDto {
	output := &dto.AppInterviewDto{
		ID:       interview.ID.String(),
		Status:   string(interview.Status),
		Messages: make([]*dto.AppInterviewMessageDto, 0),
	}

	if interview.Interviewer != nil {
		output.Interviewer = MapInterviewerModelToAppDto(interview.Interviewer)
	}

	if interview.Result != nil {
		output.Result = MapInterviewResultToAppDto(interview.Result)
	}

	if interview.Messages != nil && len(interview.Messages) > 0 {
		output.Messages = make([]*dto.AppInterviewMessageDto, len(interview.Messages))
		for i, msg := range interview.Messages {
			output.Messages[i] = MapInterviewMessageModelToAppDto(msg)
		}
	}

	return output
}

func MapInterviewModelToAppDtoList(interviews []*model.Interview) []*dto.AppInterviewDto {
	output := make([]*dto.AppInterviewDto, len(interviews))
	for i, interview := range interviews {
		output[i] = MapInterviewModelToAppDto(interview)
	}
	return output
}
