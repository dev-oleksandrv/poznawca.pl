package mapper

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/dto"
)

func MapInterviewResultModelToOutput(interviewResult *model.InterviewResult) *dto.InterviewResultOutputDto {
	return &dto.InterviewResultOutputDto{
		GrammarScore:     interviewResult.GrammarScore,
		AccuracyScore:    interviewResult.AccuracyScore,
		TotalScore:       interviewResult.TotalScore,
		GrammarFeedback:  interviewResult.GrammarFeedback,
		AccuracyFeedback: interviewResult.AccuracyFeedback,
		TotalFeedback:    interviewResult.TotalFeedback,
	}
}
