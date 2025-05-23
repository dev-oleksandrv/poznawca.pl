package event

import "github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/dto"

type AppInterviewEventType string

const (
	AppInterviewEventErrorMessageSentType AppInterviewEventType = "error_message_sent"

	AppInterviewEventInterviewerMessagePendingType AppInterviewEventType = "interviewer_message_pending"
	AppInterviewEventInterviewerMessageSentType    AppInterviewEventType = "interviewer_message_sent"

	AppInterviewEventUserMessageSentType AppInterviewEventType = "user_message_sent"
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
