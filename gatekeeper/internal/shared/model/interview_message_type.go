package model

import (
	"database/sql/driver"
	"errors"
)

type InterviewMessageType string

const (
	InterviewMessageTypeError   InterviewMessageType = "error"
	InterviewMessageTypeDefault InterviewMessageType = "default"
)

func (r *InterviewMessageType) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("interview message type assertion to string failed")
	}

	switch strValue {
	case string(InterviewMessageTypeError), string(InterviewMessageTypeDefault):
		*r = InterviewMessageType(strValue)
		return nil
	default:
		return errors.New("invalid interview message type value")
	}
}

func (r *InterviewMessageType) Value() (driver.Value, error) {
	switch *r {
	case InterviewMessageTypeError, InterviewMessageTypeDefault:
		return string(*r), nil
	default:
		return nil, errors.New("invalid interview message type value")
	}
}
