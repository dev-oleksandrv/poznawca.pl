package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/config"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/pkg/aiutils"
	"github.com/sashabaranov/go-openai"
)

type AppOpenAIInterviewService interface {
	CreateThread(ctx context.Context, reqInput *dto.AppOpenAICreateInterviewRequestDto) (*dto.AppOpenAICreateInterviewResponseDto, error)
	SendUserAnswer(ctx context.Context, reqInput *dto.AppOpenAIInterviewSendUserAnswerRequestDto) (*dto.AppOpenAIInterviewAssistantQuestionResponseDto, error)
	GetResults(ctx context.Context, reqInput *dto.AppOpenAIInterviewGetResultsRequestDto) (*dto.AppOpenAIInterviewAssistantResultsResponseDto, error)
}

type appOpenAIInterviewServiceImpl struct {
	openaiConfig *config.OpenAIConfig
	openaiClient *openai.Client
}

func NewAppOpenAIInterviewService(openaiConfig *config.OpenAIConfig, openaiClient *openai.Client) AppOpenAIInterviewService {
	return &appOpenAIInterviewServiceImpl{
		openaiConfig: openaiConfig,
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

func (s *appOpenAIInterviewServiceImpl) SendUserAnswer(ctx context.Context, reqInput *dto.AppOpenAIInterviewSendUserAnswerRequestDto) (*dto.AppOpenAIInterviewAssistantQuestionResponseDto, error) {
	prompt, err := json.Marshal(&dto.AppOpenAIInterviewAssistantUserAnswerRequestDto{
		AppOpenAIInterviewAssistantBaseRequestDto: dto.AppOpenAIInterviewAssistantBaseRequestDto{
			Type: dto.AppOpenAIInterviewAssistantRequestUserAnswerType,
		},
		Content:  reqInput.Content,
		Language: reqInput.Language,
	})
	if err != nil {
		return nil, err
	}

	msgContent, err := s.sendPrompt(ctx, reqInput.ThreadID, string(prompt), openai.ChatMessageRoleUser)
	if err != nil {
		return nil, err
	}

	output, err := s.parseQuestionResponse(msgContent)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *appOpenAIInterviewServiceImpl) GetResults(ctx context.Context, reqInput *dto.AppOpenAIInterviewGetResultsRequestDto) (*dto.AppOpenAIInterviewAssistantResultsResponseDto, error) {
	prompt, err := json.Marshal(&dto.AppOpenAIInterviewAssistantGetResultsRequestDto{
		AppOpenAIInterviewAssistantBaseRequestDto: dto.AppOpenAIInterviewAssistantBaseRequestDto{
			Type: dto.AppOpenAIInterviewAssistantRequestGetResultsType,
		},
	})
	if err != nil {
		return nil, err
	}

	msgContent, err := s.sendPrompt(ctx, reqInput.ThreadID, string(prompt), openai.ChatMessageRoleUser)
	if err != nil {
		return nil, err
	}

	output, err := s.parseResultsResponse(msgContent)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *appOpenAIInterviewServiceImpl) parseQuestionResponse(msgContent string) (*dto.AppOpenAIInterviewAssistantQuestionResponseDto, error) {
	var output *dto.AppOpenAIInterviewAssistantQuestionResponseDto
	if err := json.Unmarshal([]byte(msgContent), &output); err != nil {
		return nil, err
	}

	if output.Type != dto.AppOpenAIInterviewAssistantResponseQuestionType {
		return nil, fmt.Errorf("unexpected output type: %s", output.Type)
	}

	return output, nil
}

func (s *appOpenAIInterviewServiceImpl) parseResultsResponse(msgContent string) (*dto.AppOpenAIInterviewAssistantResultsResponseDto, error) {
	var output *dto.AppOpenAIInterviewAssistantResultsResponseDto
	if err := json.Unmarshal([]byte(msgContent), &output); err != nil {
		return nil, err
	}

	if output.Type != dto.AppOpenAIInterviewAssistantResponseResultsType {
		return nil, fmt.Errorf("unexpected output type: %s", output.Type)
	}

	return output, nil
}

func (s *appOpenAIInterviewServiceImpl) sendPrompt(ctx context.Context, threadID, prompt, role string) (string, error) {
	if _, err := s.openaiClient.CreateMessage(ctx, threadID, openai.MessageRequest{
		Role:    role,
		Content: prompt,
	}); err != nil {
		return "", err
	}

	run, err := s.openaiClient.CreateRun(ctx, threadID, openai.RunRequest{
		AssistantID: s.openaiConfig.InterviewAssistantID,
	})
	if err != nil {
		return "", err
	}

	if _, err := aiutils.PollRunStatus(ctx, s.openaiClient, threadID, run.ID); err != nil {
		return "", err
	}

	limit, order := 1, "desc"
	assistantResponse, err := s.openaiClient.ListMessage(ctx, threadID, &limit, &order, nil, nil, &run.ID)
	if err != nil {
		return "", err
	}

	msgContent := assistantResponse.Messages[0].Content[0].Text.Value
	if msgContent == "" {
		return "", fmt.Errorf("empty message content")
	}

	return msgContent, nil
}
