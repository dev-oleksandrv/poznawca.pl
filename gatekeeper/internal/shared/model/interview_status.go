package model

import (
	"database/sql/driver"
	"errors"
)

type InterviewStatus string

const (
	InterviewStatusPending   InterviewStatus = "pending"
	InterviewStatusActive    InterviewStatus = "active"
	InterviewStatusCompleted InterviewStatus = "completed"
	InterviewStatusAbandoned InterviewStatus = "abandoned"
)

func (q *InterviewStatus) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("interview status assertion to string failed")
	}

	switch strValue {
	case string(InterviewStatusPending), string(InterviewStatusActive), string(InterviewStatusCompleted), string(InterviewStatusAbandoned):
		*q = InterviewStatus(strValue)
		return nil
	default:
		return errors.New("invalid interview status value")
	}
}

func (q *InterviewStatus) Value() (driver.Value, error) {
	switch *q {
	case InterviewStatusPending, InterviewStatusActive, InterviewStatusCompleted, InterviewStatusAbandoned:
		return string(*q), nil
	default:
		return nil, errors.New("invalid interview status value")
	}
}
