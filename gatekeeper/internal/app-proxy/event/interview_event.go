package event

import "github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/dto"

type AppInterviewEventType string

const (
	AppInterviewEventErrorMessageSentType AppInterviewEventType = "error_message_sent"

	AppInterviewEventInterviewerMessagePendingType AppInterviewEventType = "interviewer_message_pending"
	AppInterviewEventInterviewerMessageSentType    AppInterviewEventType = "interviewer_message_sent"

	AppInterviewEventResultsPendingType AppInterviewEventType = "results_pending"
	AppInterviewEventResultsSentType    AppInterviewEventType = "results_sent"

	AppInterviewEventUserMessageSentType       AppInterviewEventType = "user_message_sent"
	AppInterviewEventUserCompleteInterviewType AppInterviewEventType = "user_complete_interview"

	AppInterviewEventInterviewCompletedType AppInterviewEventType = "interview_completed"
)

type AppBaseInterviewEvent struct {
	Type AppInterviewEventType `json:"type"`
}

type AppInterviewInterviewerMessagePendingEvent struct {
	AppBaseInterviewEvent
}

type AppInterviewInterviewerMessageSentEvent struct {
	AppBaseInterviewEvent
	Data dto.AppInterviewMessageDto `json:"data"`
}

type AppInterviewErrorMessageSentEvent struct {
	AppBaseInterviewEvent
	ErrorKey string `json:"error_key"`
}

type AppInterviewUserMessageSentEvent struct {
	AppBaseInterviewEvent
	Content string `json:"content"`
}

type AppInterviewResultsPendingEvent struct {
	AppBaseInterviewEvent
}

type AppInterviewResultsSentEvent struct {
	AppBaseInterviewEvent
	Data dto.AppInterviewResultDto `json:"data"`
}

type AppInterviewUserCompleteEvent struct {
	AppBaseInterviewEvent
}

type AppInterviewCompletedEvent struct {
	AppBaseInterviewEvent
}
