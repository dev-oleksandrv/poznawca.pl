package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/backoffice-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/pkg/date"
)

func MapInterviewerInputToModel(input *dto.InterviewerInputDto, existing *model.Interviewer) (*model.Interviewer, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	if existing == nil {
		existing = &model.Interviewer{}
	}

	existing.Name = input.Name
	existing.AvatarURL = input.AvatarURL
	existing.EntryMessage = input.EntryMessage
	existing.CharacterDescription = input.CharacterDescription
	existing.CharacterDescriptionTranslationKey = input.CharacterDescriptionTranslationKey

	return existing, nil
}

func MapInterviewerModelToOutput(model *model.Interviewer) *dto.InterviewerOutputDto {
	return &dto.InterviewerOutputDto{
		ID:                                 model.ID,
		Name:                               model.Name,
		AvatarURL:                          model.AvatarURL,
		EntryMessage:                       model.EntryMessage,
		CharacterDescription:               model.CharacterDescription,
		CharacterDescriptionTranslationKey: model.CharacterDescriptionTranslationKey,
		CreatedAt:                          date.FormatTime(model.CreatedAt),
		UpdatedAt:                          date.FormatTime(model.UpdatedAt),
	}
}

func MapInterviewerModelToOutputList(interviewerList []*model.Interviewer) []*dto.InterviewerOutputDto {
	outputList := make([]*dto.InterviewerOutputDto, len(interviewerList))
	for i, interviewer := range interviewerList {
		outputList[i] = MapInterviewerModelToOutput(interviewer)
	}
	return outputList
}
