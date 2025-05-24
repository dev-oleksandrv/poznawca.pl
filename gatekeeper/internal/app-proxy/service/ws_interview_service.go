package service

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/errors"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/query"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/repository"
	"github.com/google/uuid"
)

type AppWSInterviewService interface {
	ActivateInterview(ctx context.Context, interviewID uuid.UUID) (*model.InterviewModel, error)
	AbandonInterview(ctx context.Context, interview *model.InterviewModel) error
	CompleteInterview(ctx context.Context, interview *model.InterviewModel) error
	CreateMessage(ctx context.Context, interviewMessage *model.InterviewMessageModel) (*model.InterviewMessageModel, error)
	ProcessMessageWithOpenAI(ctx context.Context, threadID string, interviewMessage *model.InterviewMessageModel) (*model.InterviewMessageModel, error)
	GetResultsWithOpenAI(ctx context.Context, interviewID uuid.UUID, threadID string) (*model.InterviewResultModel, error)
	CheckInterviewCompleteAvailability(ctx context.Context, interview *model.InterviewModel) (bool, error)
}

type appWSInterviewServiceImpl struct {
	interviewRepository        repository.InterviewRepository
	interviewMessageRepository repository.InterviewMessageRepository
	interviewResultRepository  repository.InterviewResultRepository
	openaiInterviewService     AppOpenAIInterviewService
}

type NewAppWSInterviewServiceConfig struct {
	InterviewRepository        repository.InterviewRepository
	InterviewMessageRepository repository.InterviewMessageRepository
	InterviewResultRepository  repository.InterviewResultRepository
	OpenAIInterviewService     AppOpenAIInterviewService
}

func NewAppWSInterviewService(cfg *NewAppWSInterviewServiceConfig) AppWSInterviewService {
	return &appWSInterviewServiceImpl{
		interviewRepository:        cfg.InterviewRepository,
		interviewMessageRepository: cfg.InterviewMessageRepository,
		interviewResultRepository:  cfg.InterviewResultRepository,
		openaiInterviewService:     cfg.OpenAIInterviewService,
	}
}

func (s *appWSInterviewServiceImpl) ActivateInterview(ctx context.Context, interviewID uuid.UUID) (*model.InterviewModel, error) {
	if interviewID == uuid.Nil {
		return nil, errors.ErrInvalidID
	}

	interview, err := s.interviewRepository.FindByID(ctx, interviewID, query.InterviewQueryWithInterviewer())
	if err != nil {
		return nil, err
	}

	if interview == nil {
		return nil, errors.ErrInterviewNotFound
	}

	if interview.Status != model.InterviewStatusPending {
		return nil, errors.ErrInvalidInitialStatus
	}

	interview.Status = model.InterviewStatusActive
	if err := s.interviewRepository.UpdateColumn(ctx, interview.ID, "status", model.InterviewStatusActive); err != nil {
		return nil, err
	}

	return interview, nil
}

func (s *appWSInterviewServiceImpl) AbandonInterview(ctx context.Context, interview *model.InterviewModel) error {
	if interview == nil {
		return errors.ErrInterviewNotFound
	}

	if interview.ID == uuid.Nil {
		return errors.ErrInvalidID
	}

	if interview.Status != model.InterviewStatusActive {
		return errors.ErrInvalidStatusToUpdate
	}

	interview.Status = model.InterviewStatusAbandoned
	return s.interviewRepository.UpdateColumn(ctx, interview.ID, "status", model.InterviewStatusAbandoned)
}

func (s *appWSInterviewServiceImpl) CompleteInterview(ctx context.Context, interview *model.InterviewModel) error {
	if interview == nil {
		return errors.ErrInterviewNotFound
	}

	if interview.ID == uuid.Nil {
		return errors.ErrInvalidID
	}

	if interview.Status != model.InterviewStatusActive {
		return errors.ErrInvalidStatusToUpdate
	}

	interview.Status = model.InterviewStatusCompleted
	return s.interviewRepository.UpdateColumn(ctx, interview.ID, "status", model.InterviewStatusCompleted)
}

func (s *appWSInterviewServiceImpl) CreateMessage(ctx context.Context, message *model.InterviewMessageModel) (*model.InterviewMessageModel, error) {
	if message.InterviewID == uuid.Nil {
		return nil, errors.ErrInvalidID
	}

	if message.ContentText == "" {
		return nil, errors.ErrEmptyContentText
	}

	return s.interviewMessageRepository.Create(ctx, message)
}

func (s *appWSInterviewServiceImpl) ProcessMessageWithOpenAI(ctx context.Context, threadID string, userMessage *model.InterviewMessageModel) (*model.InterviewMessageModel, error) {
	if userMessage == nil {
		return nil, errors.ErrInvalidMessage
	}

	if userMessage.InterviewID == uuid.Nil {
		return nil, errors.ErrInvalidID
	}

	if userMessage.ContentText == "" {
		return nil, errors.ErrEmptyContentText
	}

	aiResponse, err := s.openaiInterviewService.SendUserAnswer(ctx, &dto.AppOpenAIInterviewSendUserAnswerRequestDto{
		ThreadID: threadID,
		Content:  userMessage.ContentText,
		Language: "pl",
	})
	if err != nil {
		return nil, err
	}

	assistantMessage, err := s.CreateMessage(ctx, &model.InterviewMessageModel{
		InterviewID:            userMessage.InterviewID,
		Role:                   model.InterviewMessageRoleInterviewer,
		ContentText:            aiResponse.ContentText,
		TipsText:               aiResponse.TipsText,
		ContentTranslationText: aiResponse.ContentTranslationText,
		IsLastMessage:          aiResponse.IsLastMessage,
	})
	if err != nil {
		return nil, err
	}

	return assistantMessage, nil
}

func (s *appWSInterviewServiceImpl) GetResultsWithOpenAI(ctx context.Context, interviewID uuid.UUID, threadID string) (*model.InterviewResultModel, error) {
	if threadID == "" {
		return nil, errors.ErrInvalidID
	}

	assistantResult, err := s.openaiInterviewService.GetResults(ctx, &dto.AppOpenAIInterviewGetResultsRequestDto{
		ThreadID: threadID,
	})
	if err != nil {
		return nil, err
	}

	interviewResult, err := s.interviewResultRepository.Create(ctx, &model.InterviewResultModel{
		InterviewID:      interviewID,
		GrammarScore:     assistantResult.GrammarScore,
		GrammarFeedback:  assistantResult.GrammarFeedback,
		AccuracyScore:    assistantResult.AccuracyScore,
		AccuracyFeedback: assistantResult.AccuracyFeedback,
		TotalScore:       assistantResult.TotalScore,
		TotalFeedback:    assistantResult.TotalFeedback,
	})
	if err != nil {
		return nil, err
	}

	return interviewResult, nil
}

func (s *appWSInterviewServiceImpl) CheckInterviewCompleteAvailability(ctx context.Context, interview *model.InterviewModel) (bool, error) {
	if interview.ID == uuid.Nil {
		return false, errors.ErrInvalidID
	}

	if interview.Status != model.InterviewStatusActive {
		return false, errors.ErrInvalidStatusToUpdate
	}

	role := model.InterviewMessageRoleUser
	messagesCount, err := s.interviewMessageRepository.GetCountByInterviewID(ctx, interview.ID, &role)
	if err != nil {
		return false, err
	}

	return messagesCount > 5, nil
}
