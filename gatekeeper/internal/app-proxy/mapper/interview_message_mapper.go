package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
)

func MapInterviewMessageModelToAppDto(interviewMessage *model.InterviewMessageModel) *dto.AppInterviewMessageDto {
	return &dto.AppInterviewMessageDto{
		ID:                     interviewMessage.ID.String(),
		ContentText:            interviewMessage.ContentText,
		ContentTranslationText: interviewMessage.ContentTranslationText,
		TipsText:               interviewMessage.TipsText,
		Role:                   string(interviewMessage.Role),
		Type:                   string(interviewMessage.Type),
		CreatedAt:              interviewMessage.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
