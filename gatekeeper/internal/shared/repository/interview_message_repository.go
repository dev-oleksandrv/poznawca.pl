package repository

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/infrastructure/database"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
)

type InterviewMessageRepository interface {
	Create(ctx context.Context, interviewMessage *model.InterviewMessageModel) (*model.InterviewMessageModel, error)
}

type interviewMessageRepositoryImpl struct {
	db *database.PGQLDatabase
}

func NewInterviewMessageRepository(db *database.PGQLDatabase) InterviewMessageRepository {
	return &interviewMessageRepositoryImpl{
		db: db,
	}
}

func (r *interviewMessageRepositoryImpl) Create(ctx context.Context, interviewMessage *model.InterviewMessageModel) (*model.InterviewMessageModel, error) {
	if err := r.db.WithContext(ctx).Create(interviewMessage).Error; err != nil {
		return nil, err
	}
	return interviewMessage, nil
}
