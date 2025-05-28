package query

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"gorm.io/gorm"
)

type InterviewQueryOption func(*gorm.DB) *gorm.DB

func InterviewQueryWithInterviewer() InterviewQueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Interviewer")
	}
}

func InterviewQueryWithMessages() InterviewQueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Messages")
	}
}

func InterviewQueryWithResult() InterviewQueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Result")
	}
}

func InterviewQueryWithStatus(status ...model.InterviewStatus) InterviewQueryOption {
	return func(db *gorm.DB) *gorm.DB {
		if len(status) == 0 {
			return db
		}
		return db.Where("status IN ?", status)
	}
}
