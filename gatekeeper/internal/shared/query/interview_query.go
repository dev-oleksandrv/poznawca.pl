package query

import "gorm.io/gorm"

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
