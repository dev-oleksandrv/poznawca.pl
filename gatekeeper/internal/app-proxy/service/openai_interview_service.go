package service

import (
	"context"
	"fmt"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/dto"
	"github.com/sashabaranov/go-openai"
)

type AppOpenAIInterviewService interface {
	CreateThread(ctx context.Context, reqInput *dto.AppOpenAICreateInterviewRequestDto) (*dto.AppOpenAICreateInterviewResponseDto, error)
}

type appOpenAIInterviewServiceImpl struct {
	openaiClient *openai.Client
}

func NewAppOpenAIInterviewService(openaiClient *openai.Client) AppOpenAIInterviewService {
	return &appOpenAIInterviewServiceImpl{
		openaiClient: openaiClient,
	}
}

func (s *appOpenAIInterviewServiceImpl) CreateThread(ctx context.Context, reqInput *dto.AppOpenAICreateInterviewRequestDto) (*dto.AppOpenAICreateInterviewResponseDto, error) {
	aiThread, err := s.openaiClient.CreateThread(ctx, openai.ThreadRequest{
		Messages: []openai.ThreadMessage{
			{
				Role:    openai.ThreadMessageRoleAssistant,
				Content: fmt.Sprintf("I'm a real interviewer with the character: %s", reqInput.Description),
			},
			{
				Role:    openai.ThreadMessageRoleAssistant,
				Content: reqInput.EntryMessage,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &dto.AppOpenAICreateInterviewResponseDto{
		ThreadID: aiThread.ID,
	}, nil
}
