package service

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/errors"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/query"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/repository"
	"github.com/google/uuid"
)

type AppWSInterviewService interface {
	ActivateInterview(ctx context.Context, interviewID uuid.UUID) (*model.InterviewModel, error)
	AbandonInterview(ctx context.Context, interview *model.InterviewModel) error
	CreateMessage(ctx context.Context, interviewMessage *model.InterviewMessageModel) (*model.InterviewMessageModel, error)
}

type appWSInterviewServiceImpl struct {
	interviewRepository        repository.InterviewRepository
	interviewMessageRepository repository.InterviewMessageRepository
}

func NewAppWSInterviewService(interviewRepository repository.InterviewRepository, interviewMessageRepository repository.InterviewMessageRepository) AppWSInterviewService {
	return &appWSInterviewServiceImpl{
		interviewRepository:        interviewRepository,
		interviewMessageRepository: interviewMessageRepository,
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

	return s.interviewRepository.UpdateColumn(ctx, interview.ID, "status", model.InterviewStatusAbandoned)
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
