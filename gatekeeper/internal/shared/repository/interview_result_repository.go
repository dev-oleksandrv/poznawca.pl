package repository

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/infrastructure/database"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
)

type InterviewResultRepository interface {
	Create(ctx context.Context, interviewResult *model.InterviewResultModel) (*model.InterviewResultModel, error)
}

type interviewResultRepositoryImpl struct {
	db *database.PGQLDatabase
}

func NewInterviewResultRepository(db *database.PGQLDatabase) InterviewResultRepository {
	return &interviewResultRepositoryImpl{
		db: db,
	}
}

func (r *interviewResultRepositoryImpl) Create(ctx context.Context, interviewResult *model.InterviewResultModel) (*model.InterviewResultModel, error) {
	if err := r.db.WithContext(ctx).Create(interviewResult).Error; err != nil {
		return nil, err
	}
	return interviewResult, nil
}
