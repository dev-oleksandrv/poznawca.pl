package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/config"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/pkg/aiutils"
	"github.com/sashabaranov/go-openai"
)

type InterviewAIService interface {
	CreateThread(ctx context.Context, input *dto.InterviewAICreateThreadInputDto) (string, error)
	SendUserMessagePrompt(ctx context.Context, input *dto.InterviewAIUserMessageUserAnswerDto) (*dto.InterviewAIOutputQuestionDto, error)
	SendGetResultsPrompt(ctx context.Context, input *dto.InterviewAIUserMessageGetResultsDto) (*dto.InterviewAIOutputResultDto, error)
}

type interviewAiServiceImpl struct {
	openaiClient *openai.Client
	openaiConfig *config.OpenAIConfig
}

func NewInterviewAIService(openaiClient *openai.Client, openaiConfig *config.OpenAIConfig) InterviewAIService {
	return &interviewAiServiceImpl{
		openaiClient: openaiClient,
		openaiConfig: openaiConfig,
	}
}

func (s *interviewAiServiceImpl) CreateThread(ctx context.Context, input *dto.InterviewAICreateThreadInputDto) (string, error) {
	aiThread, err := s.openaiClient.CreateThread(ctx, openai.ThreadRequest{
		Messages: []openai.ThreadMessage{
			{
				Role:    openai.ThreadMessageRoleAssistant,
				Content: fmt.Sprintf("I'm a real interviewer with the character: %s", input.CharacterDescription),
			},
			{
				Role:    openai.ThreadMessageRoleAssistant,
				Content: input.EntryMessage,
			},
		},
	})
	if err != nil {
		return "", err
	}

	return aiThread.ID, nil
}

func (s *interviewAiServiceImpl) SendUserMessagePrompt(ctx context.Context, input *dto.InterviewAIUserMessageUserAnswerDto) (*dto.InterviewAIOutputQuestionDto, error) {
	prompt, err := json.Marshal(&dto.InterviewAIPromptUserAnswerDto{
		InterviewAIPromptBaseDto: dto.InterviewAIPromptBaseDto{
			Type: dto.InterviewAIPromptUserAnswerType,
		},
		Content:  input.Content,
		Language: input.Language,
	})
	if err != nil {
		return nil, err
	}

	if _, err := s.openaiClient.CreateMessage(ctx, input.ThreadID, openai.MessageRequest{
		Role:    openai.ChatMessageRoleUser,
		Content: string(prompt),
	}); err != nil {
		return nil, err
	}

	run, err := s.openaiClient.CreateRun(ctx, input.ThreadID, openai.RunRequest{
		AssistantID: s.openaiConfig.InterviewAssistantID,
	})
	if err != nil {
		return nil, err
	}

	if _, err := aiutils.PollRunStatus(ctx, s.openaiClient, input.ThreadID, run.ID); err != nil {
		return nil, err
	}

	limit, order := 1, "desc"
	assistantResponse, err := s.openaiClient.ListMessage(ctx, input.ThreadID, &limit, &order, nil, nil, &run.ID)
	if err != nil {
		return nil, err
	}

	msgContent := assistantResponse.Messages[0].Content[0].Text.Value
	if msgContent == "" {
		return nil, fmt.Errorf("empty message content")
	}

	var output *dto.InterviewAIOutputQuestionDto
	if err := json.Unmarshal([]byte(msgContent), &output); err != nil {
		return nil, err
	}

	if output.Type != dto.InterviewAIOutputQuestionType {
		return nil, fmt.Errorf("unexpected output type: %s", output.Type)
	}

	return output, nil
}

func (s *interviewAiServiceImpl) SendGetResultsPrompt(ctx context.Context, input *dto.InterviewAIUserMessageGetResultsDto) (*dto.InterviewAIOutputResultDto, error) {
	prompt, err := json.Marshal(&dto.InterviewAIPromptUserAnswerDto{
		InterviewAIPromptBaseDto: dto.InterviewAIPromptBaseDto{
			Type: dto.InterviewAIPromptGetResultsType,
		},
	})
	if err != nil {
		return nil, err
	}

	if _, err := s.openaiClient.CreateMessage(ctx, input.ThreadID, openai.MessageRequest{
		Role:    openai.ChatMessageRoleUser,
		Content: string(prompt),
	}); err != nil {
		return nil, err
	}

	run, err := s.openaiClient.CreateRun(ctx, input.ThreadID, openai.RunRequest{
		AssistantID: s.openaiConfig.InterviewAssistantID,
	})
	if err != nil {
		return nil, err
	}

	if _, err := aiutils.PollRunStatus(ctx, s.openaiClient, input.ThreadID, run.ID); err != nil {
		return nil, err
	}

	limit, order := 1, "desc"
	assistantResponse, err := s.openaiClient.ListMessage(ctx, input.ThreadID, &limit, &order, nil, nil, &run.ID)
	if err != nil {
		return nil, err
	}

	msgContent := assistantResponse.Messages[0].Content[0].Text.Value
	if msgContent == "" {
		return nil, fmt.Errorf("empty message content")
	}

	var output *dto.InterviewAIOutputResultDto
	if err := json.Unmarshal([]byte(msgContent), &output); err != nil {
		return nil, err
	}

	if output.Type != dto.InterviewAIOutputResultsType {
		return nil, fmt.Errorf("unexpected output type: %s", output.Type)
	}

	return output, nil
}
