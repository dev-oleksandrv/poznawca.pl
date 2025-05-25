package service

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/errors"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/repository"
	"github.com/google/uuid"
)

type AppInterviewerService interface {
	FindAll(ctx context.Context) ([]*model.Interviewer, error)
	FindRandom(ctx context.Context) (*model.Interviewer, error)
	FindByIDOrRandom(ctx context.Context, id *uuid.UUID) (*model.Interviewer, error)
}

type appInterviewerServiceImpl struct {
	interviewerRepository repository.InterviewerRepository
}

func NewAppInterviewerService(interviewerRepository repository.InterviewerRepository) AppInterviewerService {
	return &appInterviewerServiceImpl{
		interviewerRepository: interviewerRepository,
	}
}

func (s *appInterviewerServiceImpl) FindAll(ctx context.Context) ([]*model.Interviewer, error) {
	return s.interviewerRepository.FindAll(ctx)
}

func (s *appInterviewerServiceImpl) FindRandom(ctx context.Context) (*model.Interviewer, error) {
	return s.interviewerRepository.FindRandom(ctx)
}

func (s *appInterviewerServiceImpl) FindByIDOrRandom(ctx context.Context, id *uuid.UUID) (*model.Interviewer, error) {
	if id != nil && *id == uuid.Nil {
		return nil, errors.ErrInvalidID
	}

	var interviewer *model.Interviewer
	if id == nil {
		randomInterviewer, err := s.interviewerRepository.FindRandom(ctx)
		if err != nil {
			return nil, err
		}
		if randomInterviewer == nil {
			return nil, errors.ErrNoInterviewerFound
		}
		interviewer = randomInterviewer
	} else {
		dbInterviewer, err := s.interviewerRepository.FindByID(ctx, *id)
		if err != nil {
			return nil, err
		}
		if dbInterviewer == nil {
			return nil, errors.ErrNoInterviewerFound
		}
		interviewer = dbInterviewer
	}

	return interviewer, nil
}
