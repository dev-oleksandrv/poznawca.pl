package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/backoffice-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
)

func MapInterviewerModelToBackofficeDto(interviewer *model.InterviewerModel) *dto.BackofficeInterviewerDto {
	return &dto.BackofficeInterviewerDto{
		ID:                        interviewer.ID.String(),
		Name:                      interviewer.Name,
		AvatarURL:                 interviewer.AvatarURL,
		EntryMessage:              interviewer.EntryMessage,
		Description:               interviewer.Description,
		DescriptionTranslationKey: interviewer.DescriptionTranslationKey,
		CreatedAt:                 interviewer.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:                 interviewer.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func MapInterviewerModelToBackofficeDtoList(interviewers []*model.InterviewerModel) []*dto.BackofficeInterviewerDto {
	interviewerDtos := make([]*dto.BackofficeInterviewerDto, len(interviewers))
	for i, interviewer := range interviewers {
		interviewerDtos[i] = MapInterviewerModelToBackofficeDto(interviewer)
	}
	return interviewerDtos
}
