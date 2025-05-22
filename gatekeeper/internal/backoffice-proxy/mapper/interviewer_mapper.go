package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/backoffice-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
)

func MapInterviewerModelToDto(interviewer *model.InterviewerModel) *dto.InterviewerDto {
	return &dto.InterviewerDto{
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

func MapInterviewerModelToDtoList(interviewers []*model.InterviewerModel) []*dto.InterviewerDto {
	interviewerDtos := make([]*dto.InterviewerDto, len(interviewers))
	for i, interviewer := range interviewers {
		interviewerDtos[i] = MapInterviewerModelToDto(interviewer)
	}
	return interviewerDtos
}
