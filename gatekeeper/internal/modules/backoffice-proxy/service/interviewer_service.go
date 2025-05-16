package service

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/backoffice-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/backoffice-proxy/mapper"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/repository"
	"github.com/google/uuid"
)

type InterviewerService interface {
	FindAll(ctx context.Context) ([]*dto.InterviewerOutputDto, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.InterviewerOutputDto, error)
	Create(ctx context.Context, interviewer *dto.InterviewerInputDto) (*dto.InterviewerOutputDto, error)
	Update(ctx context.Context, id uuid.UUID, interviewer *dto.InterviewerInputDto) (*dto.InterviewerOutputDto, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type interviewerServiceImpl struct {
	interviewerRepository repository.InterviewerRepository
}

func NewInterviewerService(interviewerRepository repository.InterviewerRepository) InterviewerService {
	return &interviewerServiceImpl{
		interviewerRepository: interviewerRepository,
	}
}

func (s *interviewerServiceImpl) FindAll(ctx context.Context) ([]*dto.InterviewerOutputDto, error) {
	interviewers, err := s.interviewerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.MapInterviewerModelToOutputList(interviewers), nil
}

func (s *interviewerServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*dto.InterviewerOutputDto, error) {
	interviewer, err := s.interviewerRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.MapInterviewerModelToOutput(interviewer), nil
}

func (s *interviewerServiceImpl) Create(ctx context.Context, interviewer *dto.InterviewerInputDto) (*dto.InterviewerOutputDto, error) {
	interviewerModel, err := mapper.MapInterviewerInputToModel(interviewer, nil)
	if err != nil {
		return nil, err
	}

	createdInterviewer, err := s.interviewerRepository.Create(ctx, interviewerModel)
	if err != nil {
		return nil, err
	}

	return mapper.MapInterviewerModelToOutput(createdInterviewer), nil
}

func (s *interviewerServiceImpl) Update(ctx context.Context, id uuid.UUID, interviewer *dto.InterviewerInputDto) (*dto.InterviewerOutputDto, error) {
	existingInterviewer, err := s.interviewerRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	interviewerModel, err := mapper.MapInterviewerInputToModel(interviewer, existingInterviewer)
	if err != nil {
		return nil, err
	}

	updatedInterviewer, err := s.interviewerRepository.Update(ctx, interviewerModel)
	if err != nil {
		return nil, err
	}

	return mapper.MapInterviewerModelToOutput(updatedInterviewer), nil
}

func (s *interviewerServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.interviewerRepository.Delete(ctx, id)
}
