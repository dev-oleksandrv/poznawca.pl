package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/dto"
)

func MapInterviewMessageModelToOutput(model *model.InterviewMessage) *dto.InterviewMessageOutputDto {
	return &dto.InterviewMessageOutputDto{
		ID:                     model.ID,
		ContentText:            model.ContentText,
		ContentTranslationText: model.ContentTranslationText,
		TipsText:               model.TipsText,
		Role:                   model.Role,
	}
}
