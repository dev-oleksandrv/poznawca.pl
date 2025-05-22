package service

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/repository"
)

type AppInterviewerService interface {
	FindAll(ctx context.Context) ([]*model.InterviewerModel, error)
}

type appInterviewerServiceImpl struct {
	interviewerRepository repository.InterviewerRepository
}

func NewAppInterviewerService(interviewerRepository repository.InterviewerRepository) AppInterviewerService {
	return &appInterviewerServiceImpl{
		interviewerRepository: interviewerRepository,
	}
}

func (s *appInterviewerServiceImpl) FindAll(ctx context.Context) ([]*model.InterviewerModel, error) {
	return s.interviewerRepository.FindAll(ctx)
}
