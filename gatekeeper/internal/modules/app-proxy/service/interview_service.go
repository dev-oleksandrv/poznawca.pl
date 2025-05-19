package service

import (
	"context"
	"errors"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/config"
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
	UpdateStatus(ctx context.Context, input *dto.UpdateInterviewStatusInputDto) error
	ProcessClientMessage(ctx context.Context, input *dto.ProcessInterviewClientMessageInputDto) (*dto.InterviewMessageOutputDto, bool, error)
	GenerateResults(ctx context.Context, input *dto.GenerateInterviewResultsInputDto) (*dto.InterviewResultOutputDto, error)
	CreateInitialMessage(ctx context.Context, input *dto.CreateInterviewInitialMessageInputDto) (*dto.InterviewMessageOutputDto, error)
}

type interviewServiceImpl struct {
	openaiClient               *openai.Client
	openaiConfig               *config.OpenAIConfig
	interviewAIService         InterviewAIService
	interviewRepository        repository.InterviewRepository
	interviewerRepository      repository.InterviewerRepository
	interviewMessageRepository repository.InterviewMessageRepository
	interviewResultRepository  repository.InterviewResultRepository
}

func NewInterviewService(
	openaiClient *openai.Client,
	openaiConfig *config.OpenAIConfig,
	interviewAIService InterviewAIService,
	interviewRepository repository.InterviewRepository,
	interviewerRepository repository.InterviewerRepository,
	interviewMessageRepository repository.InterviewMessageRepository,
	interviewResultRepository repository.InterviewResultRepository,
) InterviewService {
	return &interviewServiceImpl{
		openaiClient:               openaiClient,
		openaiConfig:               openaiConfig,
		interviewAIService:         interviewAIService,
		interviewRepository:        interviewRepository,
		interviewerRepository:      interviewerRepository,
		interviewMessageRepository: interviewMessageRepository,
		interviewResultRepository:  interviewResultRepository,
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

	createThreadInput := &dto.InterviewAICreateThreadInputDto{
		EntryMessage:         interviewer.EntryMessage,
		CharacterDescription: interviewer.CharacterDescription,
	}
	threadID, err := s.interviewAIService.CreateThread(ctx, createThreadInput)
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

func (s *interviewServiceImpl) ProcessClientMessage(ctx context.Context, input *dto.ProcessInterviewClientMessageInputDto) (*dto.InterviewMessageOutputDto, bool, error) {
	clientInterviewMessage := &model.InterviewMessage{
		ContentText: input.Content,
		InterviewID: input.InterviewID,
		Role:        model.InterviewMessageRoleUser,
	}
	if _, err := s.interviewMessageRepository.Create(ctx, clientInterviewMessage); err != nil {
		return nil, false, err
	}

	aiOutput, err := s.interviewAIService.SendUserMessagePrompt(ctx, &dto.InterviewAIUserMessageUserAnswerDto{
		ThreadID: input.ThreadID,
		Content:  input.Content,
		Language: "pl",
	})
	if err != nil {
		return nil, false, err
	}

	assistantMessage := &model.InterviewMessage{
		ContentText:            aiOutput.ContentText,
		InterviewID:            input.InterviewID,
		TipsText:               aiOutput.TipsText,
		ContentTranslationText: aiOutput.ContentTranslationText,
		Role:                   model.InterviewMessageRoleInterviewer,
	}
	if _, err := s.interviewMessageRepository.Create(ctx, assistantMessage); err != nil {
		return nil, false, err
	}

	return mapper.MapInterviewMessageModelToOutput(assistantMessage), aiOutput.IsLastMessage, err
}

func (s *interviewServiceImpl) GenerateResults(ctx context.Context, input *dto.GenerateInterviewResultsInputDto) (*dto.InterviewResultOutputDto, error) {
	aiOutput, err := s.interviewAIService.SendGetResultsPrompt(ctx, &dto.InterviewAIUserMessageGetResultsDto{
		ThreadID: input.ThreadID,
	})
	if err != nil {
		return nil, err
	}

	resultModel := &model.InterviewResult{
		InterviewID:      input.InterviewID,
		GrammarScore:     aiOutput.GrammarScore,
		GrammarFeedback:  aiOutput.GrammarFeedback,
		AccuracyScore:    aiOutput.AccuracyScore,
		AccuracyFeedback: aiOutput.AccuracyFeedback,
		TotalScore:       aiOutput.TotalScore,
		TotalFeedback:    aiOutput.TotalFeedback,
	}
	if _, err := s.interviewResultRepository.Create(ctx, resultModel); err != nil {
		return nil, err
	}

	return mapper.MapInterviewResultModelToOutput(resultModel), nil
}

func (s *interviewServiceImpl) CreateInitialMessage(ctx context.Context, input *dto.CreateInterviewInitialMessageInputDto) (*dto.InterviewMessageOutputDto, error) {
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
