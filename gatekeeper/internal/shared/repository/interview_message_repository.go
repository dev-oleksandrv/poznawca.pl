package repository

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/infrastructure/database"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"github.com/google/uuid"
)

type InterviewMessageRepository interface {
	Create(ctx context.Context, interviewMessage *model.InterviewMessageModel) (*model.InterviewMessageModel, error)
	GetCountByInterviewID(ctx context.Context, interviewID uuid.UUID, role *model.InterviewMessageRole) (int64, error)
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

func (r *interviewMessageRepositoryImpl) GetCountByInterviewID(ctx context.Context, interviewID uuid.UUID, role *model.InterviewMessageRole) (int64, error) {
	var count int64
	dbQuery := r.db.WithContext(ctx).Model(&model.InterviewMessageModel{}).Where("interview_id = ?", interviewID)

	if role != nil {
		dbQuery = dbQuery.Where("role = ?", *role)
	}

	if err := dbQuery.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
