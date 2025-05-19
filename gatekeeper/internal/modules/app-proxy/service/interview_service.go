package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/config"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/mapper"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/query"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/repository"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/pkg/aiutils"
	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
	"log/slog"
)

type InterviewService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*dto.InterviewOutputDto, error)
	Create(ctx context.Context) (*dto.InterviewOutputDto, error)
	UpdateStatus(ctx context.Context, input *dto.UpdateInterviewStatusInputDto) error
	ProcessClientMessage(ctx context.Context, input *dto.ProcessClientMessageInputDto) (*dto.InterviewMessageOutputDto, error)
	CreateInitialMessage(ctx context.Context, input *dto.CreateInitialMessageInputDto) (*dto.InterviewMessageOutputDto, error)
}

type interviewServiceImpl struct {
	openaiClient               *openai.Client
	openaiConfig               *config.OpenAIConfig
	interviewRepository        repository.InterviewRepository
	interviewerRepository      repository.InterviewerRepository
	interviewMessageRepository repository.InterviewMessageRepository
}

func NewInterviewService(
	openaiClient *openai.Client,
	openaiConfig *config.OpenAIConfig,
	interviewRepository repository.InterviewRepository,
	interviewerRepository repository.InterviewerRepository,
	interviewMessageRepository repository.InterviewMessageRepository,
) InterviewService {
	return &interviewServiceImpl{
		openaiClient:               openaiClient,
		openaiConfig:               openaiConfig,
		interviewRepository:        interviewRepository,
		interviewerRepository:      interviewerRepository,
		interviewMessageRepository: interviewMessageRepository,
	}
}

func (s *interviewServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.InterviewOutputDto, error) {
	interviewModel, err := s.interviewRepository.FindByID(ctx, id, query.InterviewQueryWithInterviewer())
	if err != nil {
		return nil, err
	}

	return mapper.MapInterviewModelToOutput(interviewModel), nil
}

func (s *interviewServiceImpl) Create(ctx context.Context) (*dto.InterviewOutputDto, error) {
	interviewer, err := s.interviewerRepository.FindRandom(ctx)
	if err != nil {
		return nil, err
	}

	threadID, err := s.createAIThread(ctx, interviewer.EntryMessage, interviewer.CharacterDescription)
	if err != nil {
		return nil, err
	}

	if threadID == "" {
		return nil, errors.New("failed to create AI thread")
	}

	createdInterview, err := s.interviewRepository.Create(ctx, &model.Interview{
		InterviewerID: &interviewer.ID,
		Status:        model.InterviewStatusPending,
		ThreadID:      threadID,
	})
	if err != nil {
		return nil, err
	}

	return mapper.MapInterviewModelToOutput(createdInterview), nil
}

func (s *interviewServiceImpl) UpdateStatus(ctx context.Context, input *dto.UpdateInterviewStatusInputDto) error {
	_, err := s.interviewRepository.Update(ctx, &model.Interview{
		ID:     input.InterviewID,
		Status: input.Status,
	})

	return err
}

func (s *interviewServiceImpl) ProcessClientMessage(ctx context.Context, input *dto.ProcessClientMessageInputDto) (*dto.InterviewMessageOutputDto, error) {
	clientInterviewMessage := &model.InterviewMessage{
		ContentText: input.Content,
		InterviewID: input.InterviewID,
		Role:        model.InterviewMessageRoleUser,
	}
	if _, err := s.interviewMessageRepository.Create(ctx, clientInterviewMessage); err != nil {
		return nil, err
	}

	prompt, err := json.Marshal(&dto.AIPromptInputDto{
		Content: input.Content,
		// TODO: Add support for multiple languages
		Language: "pl",
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

	slog.Info("AI Response", "response", msgContent)

	var output *dto.AIPromptOutputDto
	if err := json.Unmarshal([]byte(msgContent), &output); err != nil {
		return nil, err
	}

	if output.Done {
		// TODO: Handle interview completion
		return nil, nil
	}

	assistantMessage := &model.InterviewMessage{
		ContentText: output.Content,
		InterviewID: input.InterviewID,
		TipsText:    output.Tips,
		Role:        model.InterviewMessageRoleInterviewer,
	}
	if _, err := s.interviewMessageRepository.Create(ctx, assistantMessage); err != nil {
		return nil, err
	}

	return mapper.MapInterviewMessageModelToOutput(assistantMessage), err
}

func (s *interviewServiceImpl) CreateInitialMessage(ctx context.Context, input *dto.CreateInitialMessageInputDto) (*dto.InterviewMessageOutputDto, error) {
	interviewMessage := &model.InterviewMessage{
		ContentText: input.ContentText,
		InterviewID: input.InterviewID,
		Role:        model.InterviewMessageRoleInterviewer,
	}
	if _, err := s.interviewMessageRepository.Create(ctx, interviewMessage); err != nil {
		return nil, err
	}

	return mapper.MapInterviewMessageModelToOutput(interviewMessage), nil
}

func (s *interviewServiceImpl) createAIThread(ctx context.Context, entryMessage, characterDescription string) (string, error) {
	aiThread, err := s.openaiClient.CreateThread(ctx, openai.ThreadRequest{
		Messages: []openai.ThreadMessage{
			{
				Role:    openai.ThreadMessageRoleAssistant,
				Content: fmt.Sprintf("I'm a real interviewer with the character: %s", characterDescription),
			},
			{
				Role:    openai.ThreadMessageRoleAssistant,
				Content: entryMessage,
			},
		},
	})
	if err != nil {
		return "", err
	}

	return aiThread.ID, nil
}
