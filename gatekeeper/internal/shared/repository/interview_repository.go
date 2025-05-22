package repository

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/infrastructure/database"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/query"
	"github.com/google/uuid"
)

type InterviewRepository interface {
	FindByID(ctx context.Context, id uuid.UUID, opts ...query.InterviewQueryOption) (*model.InterviewModel, error)
	Create(ctx context.Context, interview *model.InterviewModel) (*model.InterviewModel, error)
}

type interviewRepositoryImpl struct {
	db *database.PGQLDatabase
}

func NewInterviewRepository(db *database.PGQLDatabase) InterviewRepository {
	return &interviewRepositoryImpl{
		db: db,
	}
}

func (r *interviewRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID, opts ...query.InterviewQueryOption) (*model.InterviewModel, error) {
	var interview *model.InterviewModel
	tx := r.db.WithContext(ctx).Begin()

	for _, opt := range opts {
		tx = opt(tx)
	}

	if err := tx.First(&interview, "id = ?", id).Commit().Error; err != nil {
		tx.Rollback()

		return nil, err
	}

	return interview, nil
}

func (r *interviewRepositoryImpl) Create(ctx context.Context, interview *model.InterviewModel) (*model.InterviewModel, error) {
	if err := r.db.WithContext(ctx).Create(interview).Error; err != nil {
		return nil, err
	}
	return interview, nil
}
