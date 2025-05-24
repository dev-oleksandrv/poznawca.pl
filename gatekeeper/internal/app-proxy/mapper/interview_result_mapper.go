package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
)

func MapInterviewResultToAppDto(interviewResult *model.InterviewResultModel) *dto.AppInterviewResultDto {
	return &dto.AppInterviewResultDto{
		ID:               interviewResult.ID.String(),
		GrammarScore:     interviewResult.GrammarScore,
		AccuracyScore:    interviewResult.AccuracyScore,
		TotalScore:       interviewResult.TotalScore,
		GrammarFeedback:  interviewResult.GrammarFeedback,
		AccuracyFeedback: interviewResult.AccuracyFeedback,
		TotalFeedback:    interviewResult.TotalFeedback,
	}
}
