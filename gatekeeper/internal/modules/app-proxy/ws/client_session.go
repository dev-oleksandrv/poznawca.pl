package ws

import (
	"context"
	"encoding/json"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/service"
	"github.com/gorilla/websocket"
	"log/slog"
)

type ClientSession struct {
	socket           *websocket.Conn
	sendQueue        chan interface{}
	interview        *dto.InterviewOutputDto
	interviewService service.InterviewService
}

func NewClientSession(socket *websocket.Conn, interviewService service.InterviewService, interview *dto.InterviewOutputDto) *ClientSession {
	return &ClientSession{
		socket:           socket,
		sendQueue:        make(chan interface{}, 10),
		interview:        interview,
		interviewService: interviewService,
	}
}

func (s *ClientSession) AddToSendQueue(msg interface{}) {
	s.sendQueue <- msg
}

func (s *ClientSession) Read() {
	for {
		_, msg, err := s.socket.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				slog.Info("client disconnected")
			} else {
				slog.Error("error while reading message")
			}
			return
		}

		s.handleClientMessage(msg)
	}
}

func (s *ClientSession) Write() {
	defer func() {
		s.socket.Close()
		slog.Info("Write goroutine exiting")
	}()

	for msgObj := range s.sendQueue {
		rawMsg, err := json.Marshal(msgObj)
		if err != nil {
			slog.Error("error while marshalling message", "error", err)
			continue
		}

		if err := s.socket.WriteMessage(websocket.TextMessage, rawMsg); err != nil {
			slog.Error("error while writing message", "error", err)
			return
		}
	}
}

func (s *ClientSession) handleClientMessage(rawMsg []byte) {
	var baseEvent *BaseEvent
	if err := json.Unmarshal(rawMsg, &baseEvent); err != nil {
		slog.Error("cannot unmarshal message", "msg", rawMsg)
		return
	}

	switch baseEvent.Type {
	case ClientMessageSentEventType:
		s.handleClientMessageSentEvent(rawMsg)
	}
}

func (s *ClientSession) handleClientMessageSentEvent(rawMsg []byte) {
	var event *ClientMessageSentEvent
	if err := json.Unmarshal(rawMsg, &event); err != nil {
		slog.Error("cannot unmarshal message", "msg", rawMsg)
		return
	}

	s.AddToSendQueue(&SystemMessagePendingEvent{
		BaseEvent: BaseEvent{Type: SystemMessagePendingEventType},
	})

	outputMessage, err := s.interviewService.ProcessClientMessage(context.Background(), &dto.ProcessClientMessageInputDto{
		InterviewID: s.interview.ID,
		ThreadID:    s.interview.ThreadID,
		Content:     event.Content,
	})
	if err != nil {
		slog.Error("error while processing client message", "error", err)
		return
	}

	if outputMessage == nil {
		slog.Error("empty output message")
		return
	}

	s.AddToSendQueue(&SystemMessageSentEvent{
		BaseEvent:   BaseEvent{Type: SystemMessageSentEventType},
		ContentText: outputMessage.ContentText,
		TipsText:    outputMessage.TipsText,
	})
}
