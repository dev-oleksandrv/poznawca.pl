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

type AppInterviewService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*model.Interview, error)
	FindAll(ctx context.Context) ([]*model.Interview, error)
	Create(ctx context.Context, interview *model.Interview) (*model.Interview, error)
	Update(ctx context.Context, interview *model.Interview) (*model.Interview, error)
	UpdateStatus(ctx context.Context, interview *model.Interview) error
}

type appInterviewServiceImpl struct {
	interviewRepository    repository.InterviewRepository
	openaiInterviewService AppOpenAIInterviewService
}

func NewAppInterviewService(interviewRepository repository.InterviewRepository, openaiInterviewService AppOpenAIInterviewService) AppInterviewService {
	return &appInterviewServiceImpl{
		interviewRepository:    interviewRepository,
		openaiInterviewService: openaiInterviewService,
	}
}

func (s *appInterviewServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (*model.Interview, error) {
	if id == uuid.Nil {
		return nil, errors.ErrInvalidID
	}

	interview, err := s.interviewRepository.FindByID(ctx, id, query.InterviewQueryWithInterviewer(), query.InterviewQueryWithMessages(), query.InterviewQueryWithResult())
	if err != nil {
		return nil, err
	}

	return interview, nil
}

func (s *appInterviewServiceImpl) FindAll(ctx context.Context) ([]*model.Interview, error) {
	interviews, err := s.interviewRepository.FindAll(
		ctx,
		query.InterviewQueryWithInterviewer(),
		query.InterviewQueryWithResult(),
		query.InterviewQueryWithStatus(model.InterviewStatusCompleted, model.InterviewStatusAbandoned),
	)
	if err != nil {
		return nil, err
	}

	return interviews, nil
}

func (s *appInterviewServiceImpl) Create(ctx context.Context, interview *model.Interview) (*model.Interview, error) {
	if interview.Status != model.InterviewStatusPending {
		return nil, errors.ErrInvalidInitialStatus
	}

	if interview.InterviewerID == nil || *interview.InterviewerID == uuid.Nil {
		return nil, errors.ErrInvalidID
	}

	if interview.Interviewer == nil {
		return nil, errors.ErrNoInterviewerAttached
	}

	threadRes, err := s.openaiInterviewService.CreateThread(ctx, &dto.AppOpenAICreateInterviewRequestDto{
		Description:  interview.Interviewer.Description,
		EntryMessage: interview.Interviewer.EntryMessage,
	})
	if err != nil {
		return nil, err
	}

	interview.ThreadID = threadRes.ThreadID

	return s.interviewRepository.Create(ctx, interview)
}

func (s *appInterviewServiceImpl) Update(ctx context.Context, interview *model.Interview) (*model.Interview, error) {
	if interview.ID == uuid.Nil {
		return nil, errors.ErrInvalidID
	}

	if interview.Status == model.InterviewStatusCompleted || interview.Status == model.InterviewStatusAbandoned {
		return nil, errors.ErrInvalidStatusToUpdate
	}

	return s.interviewRepository.Update(ctx, interview)
}

func (s *appInterviewServiceImpl) UpdateStatus(ctx context.Context, interview *model.Interview) error {
	if interview.ID == uuid.Nil {
		return errors.ErrInvalidID
	}

	if interview.Status != model.InterviewStatusCompleted && interview.Status != model.InterviewStatusAbandoned {
		return errors.ErrInvalidStatusToUpdate
	}

	return s.interviewRepository.UpdateColumn(ctx, interview.ID, "status", interview.Status)
}
