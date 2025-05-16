package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/dto"
)

func MapInterviewerModelToOutput(model *model.Interviewer) *dto.InterviewerOutputDto {
	return &dto.InterviewerOutputDto{
		ID:           model.ID,
		Name:         model.Name,
		AvatarURL:    model.AvatarURL,
		EntryMessage: model.EntryMessage,
	}
}
