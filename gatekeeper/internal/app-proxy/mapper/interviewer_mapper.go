package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
)

func MapInterviewerModelToAppDto(interviewer *model.Interviewer) *dto.AppInterviewerDto {
	return &dto.AppInterviewerDto{
		ID:                        interviewer.ID.String(),
		Name:                      interviewer.Name,
		AvatarURL:                 interviewer.AvatarURL,
		Description:               interviewer.Description,
		DescriptionTranslationKey: interviewer.DescriptionTranslationKey,
	}
}

func MapInterviewerModelToAppDtoList(interviewers []*model.Interviewer) []*dto.AppInterviewerDto {
	interviewersDtos := make([]*dto.AppInterviewerDto, len(interviewers))
	for i, interviewer := range interviewers {
		interviewersDtos[i] = MapInterviewerModelToAppDto(interviewer)
	}
	return interviewersDtos
}
