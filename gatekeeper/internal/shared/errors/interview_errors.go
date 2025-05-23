package errors

import "errors"

var (
	ErrInvalidInitialStatus  = errors.New("invalid initial status")
	ErrInvalidStatusToUpdate = errors.New("invalid status to update")
	ErrInterviewNotFound     = errors.New("interview not found")
	ErrEmptyContentText      = errors.New("content text cannot be empty")
	ErrNoInterviewerAttached = errors.New("no interviewer attached")
)
