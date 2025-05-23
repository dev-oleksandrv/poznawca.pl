package service

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/errors"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/query"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/repository"
	"github.com/google/uuid"
)

type AppInterviewService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*model.InterviewModel, error)
	Create(ctx context.Context, interview *model.InterviewModel) (*model.InterviewModel, error)
	Update(ctx context.Context, interview *model.InterviewModel) (*model.InterviewModel, error)
	UpdateStatus(ctx context.Context, interview *model.InterviewModel) error
}

type appInterviewServiceImpl struct {
	interviewRepository repository.InterviewRepository
}

func NewAppInterviewService(interviewRepository repository.InterviewRepository) AppInterviewService {
	return &appInterviewServiceImpl{
		interviewRepository: interviewRepository,
	}
}

func (s *appInterviewServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*model.InterviewModel, error) {
	if id == uuid.Nil {
		return nil, errors.ErrInvalidID
	}

	interview, err := s.interviewRepository.FindByID(ctx, id, query.InterviewQueryWithInterviewer())
	if err != nil {
		return nil, err
	}

	return interview, nil
}

func (s *appInterviewServiceImpl) Create(ctx context.Context, interview *model.InterviewModel) (*model.InterviewModel, error) {
	if interview.Status != model.InterviewStatusPending {
		return nil, errors.ErrInvalidInitialStatus
	}

	if interview.InterviewerID == nil || *interview.InterviewerID == uuid.Nil {
		return nil, errors.ErrInvalidID
	}

	return s.interviewRepository.Create(ctx, interview)
}

func (s *appInterviewServiceImpl) Update(ctx context.Context, interview *model.InterviewModel) (*model.InterviewModel, error) {
	if interview.ID == uuid.Nil {
		return nil, errors.ErrInvalidID
	}

	if interview.Status == model.InterviewStatusCompleted || interview.Status == model.InterviewStatusAbandoned {
		return nil, errors.ErrInvalidStatusToUpdate
	}

	return s.interviewRepository.Update(ctx, interview)
}

func (s *appInterviewServiceImpl) UpdateStatus(ctx context.Context, interview *model.InterviewModel) error {
	if interview.ID == uuid.Nil {
		return errors.ErrInvalidID
	}

	if interview.Status != model.InterviewStatusCompleted && interview.Status != model.InterviewStatusAbandoned {
		return errors.ErrInvalidStatusToUpdate
	}

	return s.interviewRepository.UpdateColumn(ctx, interview.ID, "status", interview.Status)
}
