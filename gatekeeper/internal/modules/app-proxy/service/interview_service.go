package service

import (
	"context"
	"errors"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/mapper"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/query"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/repository"
	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

type InterviewService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*dto.InterviewOutputDto, error)
	Create(ctx context.Context) (*dto.InterviewOutputDto, error)
}

type interviewServiceImpl struct {
	openaiClient          *openai.Client
	interviewRepository   repository.InterviewRepository
	interviewerRepository repository.InterviewerRepository
}

func NewInterviewService(
	openaiClient *openai.Client,
	interviewRepository repository.InterviewRepository,
	interviewerRepository repository.InterviewerRepository,
) InterviewService {
	return &interviewServiceImpl{
		openaiClient:          openaiClient,
		interviewRepository:   interviewRepository,
		interviewerRepository: interviewerRepository,
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

	threadID, err := s.createAIThread(ctx, interviewer.EntryMessage)
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

func (s *interviewServiceImpl) createAIThread(ctx context.Context, entryMessage string) (string, error) {
	aiThread, err := s.openaiClient.CreateThread(ctx, openai.ThreadRequest{
		Messages: []openai.ThreadMessage{
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
