package repository

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/database"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
)

type InterviewMessageRepository interface {
	Create(ctx context.Context, interviewMessage *model.InterviewMessage) (*model.InterviewMessage, error)
}

type interviewMessageRepositoryImpl struct {
	db *database.PGQLDatabase
}

func NewInterviewMessageRepository(db *database.PGQLDatabase) InterviewMessageRepository {
	return &interviewMessageRepositoryImpl{
		db: db,
	}
}

func (r *interviewMessageRepositoryImpl) Create(ctx context.Context, interviewMessage *model.InterviewMessage) (*model.InterviewMessage, error) {
	if err := r.db.WithContext(ctx).Create(interviewMessage).Error; err != nil {
		return nil, err
	}
	return interviewMessage, nil
}
