package ws

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
	ContentText string `json:"content_text"`
	TipsText    string `json:"tips_text"`
}

type SystemResultPendingEvent struct {
	BaseEvent
}

type SystemResultSentEvent struct {
	BaseEvent
	TotalScore int `json:"total_score"`
}
