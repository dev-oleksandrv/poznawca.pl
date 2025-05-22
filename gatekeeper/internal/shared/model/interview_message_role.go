package model

import (
	"database/sql/driver"
	"errors"
)

type InterviewMessageRole string

const (
	InterviewMessageRoleUser        InterviewMessageRole = "user"
	InterviewMessageRoleInterviewer InterviewMessageRole = "interviewer"
	InterviewMessageRoleSystem      InterviewMessageRole = "system"
)

func (r *InterviewMessageRole) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("interview message role assertion to string failed")
	}

	switch strValue {
	case string(InterviewMessageRoleUser), string(InterviewMessageRoleInterviewer), string(InterviewMessageRoleSystem):
		*r = InterviewMessageRole(strValue)
		return nil
	default:
		return errors.New("invalid interview message role value")
	}
}

func (r *InterviewMessageRole) Value() (driver.Value, error) {
	switch *r {
	case InterviewMessageRoleUser, InterviewMessageRoleInterviewer, InterviewMessageRoleSystem:
		return string(*r), nil
	default:
		return nil, errors.New("invalid interview message role value")
	}
}
