package service

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/errors"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/repository"
	"github.com/google/uuid"
)

type BackofficeInterviewerService interface {
	FindAll(ctx context.Context) ([]*model.InterviewerModel, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.InterviewerModel, error)
	Create(ctx context.Context, interviewer *model.InterviewerModel) (*model.InterviewerModel, error)
	Update(ctx context.Context, interviewer *model.InterviewerModel) (*model.InterviewerModel, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type backofficeInterviewerServiceImpl struct {
	interviewerRepository repository.InterviewerRepository
}

func NewBackofficeInterviewerService(repository repository.InterviewerRepository) BackofficeInterviewerService {
	return &backofficeInterviewerServiceImpl{
		interviewerRepository: repository,
	}
}

func (s *backofficeInterviewerServiceImpl) FindAll(ctx context.Context) ([]*model.InterviewerModel, error) {
	return s.interviewerRepository.FindAll(ctx)
}

func (s *backofficeInterviewerServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*model.InterviewerModel, error) {
	if id == uuid.Nil {
		return nil, errors.ErrInvalidID
	}
	return s.interviewerRepository.FindByID(ctx, id)
}

func (s *backofficeInterviewerServiceImpl) Create(ctx context.Context, interviewer *model.InterviewerModel) (*model.InterviewerModel, error) {
	return s.interviewerRepository.Create(ctx, interviewer)
}

func (s *backofficeInterviewerServiceImpl) Update(ctx context.Context, interviewer *model.InterviewerModel) (*model.InterviewerModel, error) {
	return s.interviewerRepository.Update(ctx, interviewer)
}

func (s *backofficeInterviewerServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.interviewerRepository.Delete(ctx, id)
}
