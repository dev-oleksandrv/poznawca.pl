package repository

import "github.com/dev-oleksandrv/poznawca/gatekeeper/internal/database"

type InterviewCacheRepository interface{}

type interviewCacheRepositoryImpl struct {
	db *database.RedisDatabase
}

func NewInterviewCacheRepository(db *database.RedisDatabase) InterviewCacheRepository {
	return &interviewCacheRepositoryImpl{db}
}
