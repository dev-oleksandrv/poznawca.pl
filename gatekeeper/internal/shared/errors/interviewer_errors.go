package errors

import "errors"

var (
	ErrInvalidID          = errors.New("invalid ID")
	ErrNoInterviewerFound = errors.New("no interviewer found")
)
