package ws

import "github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/dto"

type EventType string

const (
	ClientMessageSentEventType  EventType = "client_message_sent"
	ClientEndInterviewEventType EventType = "client_end_interview"

	SystemMessagePendingEventType EventType = "system_message_pending"
	SystemMessageSentEventType    EventType = "system_message_sent"
	SystemResultPendingEventType  EventType = "system_result_pending"
	SystemResultSentEventType     EventType = "system_result_sent"
)

type BaseEvent struct {
	Type EventType `json:"type"`
}

type ClientMessageSentEvent struct {
	BaseEvent
	Content string `json:"content"`
}

type ClientEndInterviewEvent struct {
	BaseEvent
}

type SystemMessagePendingEvent struct {
	BaseEvent
}

type SystemMessageSentEvent struct {
	BaseEvent
	Details dto.InterviewMessageOutputDto `json:"details"`
}

type SystemResultPendingEvent struct {
	BaseEvent
}

type SystemResultSentEvent struct {
	BaseEvent
	TotalScore       int    `json:"total_score"`
	TotalFeedback    string `json:"total_feedback"`
	GrammarScore     int    `json:"grammar_score"`
	GrammarFeedback  string `json:"grammar_feedback"`
	AccuracyScore    int    `json:"accuracy_score"`
	AccuracyFeedback string `json:"accuracy_feedback"`
}
