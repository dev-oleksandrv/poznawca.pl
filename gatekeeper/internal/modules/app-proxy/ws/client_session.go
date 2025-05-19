package ws

import (
	"context"
	"encoding/json"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/service"
	"github.com/gorilla/websocket"
	"log/slog"
)

type ClientSession interface {
	SaveInitialMessage()
	AddToSendQueue(msg interface{})
	Read()
	Write()
}

type clientSessionImpl struct {
	socket           *websocket.Conn
	sendQueue        chan interface{}
	interview        *dto.InterviewOutputDto
	interviewService service.InterviewService
}

func NewClientSession(socket *websocket.Conn, interviewService service.InterviewService, interview *dto.InterviewOutputDto) ClientSession {
	return &clientSessionImpl{
		socket:           socket,
		sendQueue:        make(chan interface{}, 10),
		interview:        interview,
		interviewService: interviewService,
	}
}

func (s *clientSessionImpl) SaveInitialMessage() {
	message, err := s.interviewService.CreateInitialMessage(context.Background(), &dto.CreateInterviewInitialMessageInputDto{
		InterviewID: s.interview.ID,
		ContentText: s.interview.Interviewer.EntryMessage,
	})
	if err != nil {
		slog.Error("error while creating initial message", "error", err)
		return
	}

	event := &SystemMessageSentEvent{
		BaseEvent: BaseEvent{Type: SystemMessageSentEventType},
		Details:   *message,
	}

	s.AddToSendQueue(event)
}

func (s *clientSessionImpl) AddToSendQueue(msg interface{}) {
	s.sendQueue <- msg
}

func (s *clientSessionImpl) Read() {
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

func (s *clientSessionImpl) Write() {
	defer func() {
		slog.Info("Write goroutine exiting")
		err := s.socket.Close()
		if err != nil {
			slog.Error("error while closing socket", "error", err)
		}
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

func (s *clientSessionImpl) handleClientMessage(rawMsg []byte) {
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

func (s *clientSessionImpl) handleClientMessageSentEvent(rawMsg []byte) {
	var event *ClientMessageSentEvent
	if err := json.Unmarshal(rawMsg, &event); err != nil {
		slog.Error("cannot unmarshal message", "msg", rawMsg)
		return
	}

	s.AddToSendQueue(&SystemMessagePendingEvent{
		BaseEvent: BaseEvent{Type: SystemMessagePendingEventType},
	})

	outputMessage, isLastMessage, err := s.interviewService.ProcessClientMessage(context.Background(), &dto.ProcessInterviewClientMessageInputDto{
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
		BaseEvent: BaseEvent{Type: SystemMessageSentEventType},
		Details:   *outputMessage,
	})

	if isLastMessage == false {
		return
	}

	s.AddToSendQueue(&SystemResultPendingEvent{
		BaseEvent: BaseEvent{Type: SystemResultPendingEventType},
	})

	outputResult, err := s.interviewService.GenerateResults(context.Background(), &dto.GenerateInterviewResultsInputDto{
		ThreadID:    s.interview.ThreadID,
		InterviewID: s.interview.ID,
	})
	if err != nil {
		slog.Error("error while generating results", "error", err)
		return
	}

	if outputResult == nil {
		slog.Error("empty output result")
		return
	}

	s.AddToSendQueue(&SystemResultSentEvent{
		BaseEvent:        BaseEvent{Type: SystemResultSentEventType},
		TotalScore:       outputResult.TotalScore,
		TotalFeedback:    outputResult.TotalFeedback,
		GrammarScore:     outputResult.GrammarScore,
		GrammarFeedback:  outputResult.GrammarFeedback,
		AccuracyScore:    outputResult.AccuracyScore,
		AccuracyFeedback: outputResult.AccuracyFeedback,
	})

	s.interviewService.UpdateStatus(context.Background(), &dto.UpdateInterviewStatusInputDto{
		InterviewID: s.interview.ID,
		Status:      model.InterviewStatusCompleted,
	})
}
