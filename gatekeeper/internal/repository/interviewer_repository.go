package repository

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/database"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/google/uuid"
)

type InterviewerRepository interface {
	FindAll(ctx context.Context) ([]*model.Interviewer, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.Interviewer, error)
	FindRandom(ctx context.Context) (*model.Interviewer, error)
	Create(ctx context.Context, interviewer *model.Interviewer) (*model.Interviewer, error)
	Update(ctx context.Context, interviewer *model.Interviewer) (*model.Interviewer, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type interviewerRepositoryImpl struct {
	db *database.PGQLDatabase
}

func NewInterviewerRepository(db *database.PGQLDatabase) InterviewerRepository {
	return &interviewerRepositoryImpl{
		db: db,
	}
}

func (r *interviewerRepositoryImpl) FindAll(ctx context.Context) ([]*model.Interviewer, error) {
	var interviewers []*model.Interviewer
	err := r.db.WithContext(ctx).Find(&interviewers).Error
	if err != nil {
		return nil, err
	}
	return interviewers, nil
}

func (r *interviewerRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*model.Interviewer, error) {
	var interviewer model.Interviewer
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&interviewer).Error
	if err != nil {
		return nil, err
	}
	return &interviewer, nil
}

func (r *interviewerRepositoryImpl) FindRandom(ctx context.Context) (*model.Interviewer, error) {
	var interviewer *model.Interviewer
	if err := r.db.WithContext(ctx).Order("random()").First(&interviewer).Error; err != nil {
		return nil, err
	}
	return interviewer, nil
}

func (r *interviewerRepositoryImpl) Create(ctx context.Context, interviewer *model.Interviewer) (*model.Interviewer, error) {
	err := r.db.WithContext(ctx).Create(interviewer).Error
	if err != nil {
		return nil, err
	}
	return interviewer, nil
}

func (r *interviewerRepositoryImpl) Update(ctx context.Context, interviewer *model.Interviewer) (*model.Interviewer, error) {
	err := r.db.WithContext(ctx).Save(interviewer).Error
	if err != nil {
		return nil, err
	}
	return interviewer, nil
}

func (r *interviewerRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.Interviewer{}, id).Error
}
