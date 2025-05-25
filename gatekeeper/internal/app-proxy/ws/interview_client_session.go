package ws

import (
	"context"
	"encoding/json"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/errors"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/event"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/mapper"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/service"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"github.com/gorilla/websocket"
	"log/slog"
	"os"
)

type AppInterviewClientSession interface {
	AddToSendQueue(msg interface{})
	Write()
	Read()
	Init()
}

type appInterviewClientSessionImpl struct {
	context   context.Context
	socket    *websocket.Conn
	interview *model.Interview
	sendQueue chan interface{}
	logger    *slog.Logger
	service   service.AppWSInterviewService
}

type NewAppInterviewClientSessionConfig struct {
	Context   context.Context
	Socket    *websocket.Conn
	Interview *model.Interview
	Service   service.AppWSInterviewService
}

func NewAppInterviewClientSession(cfg *NewAppInterviewClientSessionConfig) AppInterviewClientSession {
	sendQueue := make(chan interface{})
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)).With("source", "AppInterviewClientSession")

	return &appInterviewClientSessionImpl{
		context:   cfg.Context,
		socket:    cfg.Socket,
		interview: cfg.Interview,
		sendQueue: sendQueue,
		logger:    logger,
		service:   cfg.Service,
	}
}

func (s *appInterviewClientSessionImpl) AddToSendQueue(msg interface{}) {
	s.sendQueue <- msg
}

func (s *appInterviewClientSessionImpl) Read() {
	for {
		_, msg, err := s.socket.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				s.logger.Info("client disconnected")
			} else {
				s.logger.Error("error while reading message")
			}
			return
		}

		s.handleClientMessage(msg)
	}
}

func (s *appInterviewClientSessionImpl) Write() {
	defer func() {
		s.logger.Info("Write goroutine exiting")
		err := s.socket.Close()
		if err != nil {
			s.logger.Error("error while closing socket", "error", err)
		}
	}()

	for msgObj := range s.sendQueue {
		rawMsg, err := json.Marshal(msgObj)
		if err != nil {
			s.logger.Error("error while marshalling message", "error", err)
			continue
		}

		if err := s.socket.WriteMessage(websocket.TextMessage, rawMsg); err != nil {
			s.logger.Error("error while writing message", "error", err)
			return
		}
	}
}

func (s *appInterviewClientSessionImpl) Init() {
	if s.interview.Interviewer == nil {
		s.logger.Error("interview is not assigned to an interviewer")
		return
	}

	interviewerMessage := &model.InterviewMessage{
		InterviewID: s.interview.ID,
		ContentText: s.interview.Interviewer.EntryMessage,
		Role:        model.InterviewMessageRoleInterviewer,
	}
	if _, err := s.service.CreateMessage(s.context, interviewerMessage); err != nil {
		s.logger.Error("error while creating initial message", "error", err)
		s.sendErrorMessage(errors.InterviewMessageErrorKeyFailedToInitializeInterview)
		return
	}

	s.AddToSendQueue(&event.AppInterviewInterviewerMessageSentEvent{
		AppBaseInterviewEvent: event.AppBaseInterviewEvent{Type: event.AppInterviewEventInterviewerMessageSentType},
		Data:                  *mapper.MapInterviewMessageModelToAppDto(interviewerMessage),
	})
}

func (s *appInterviewClientSessionImpl) sendErrorMessage(errorKey errors.InterviewMessageErrorKey) {
	errorMessage := &model.InterviewMessage{
		InterviewID: s.interview.ID,
		ContentText: string(errorKey),
		Role:        model.InterviewMessageRoleSystem,
		Type:        model.InterviewMessageTypeError,
	}
	if _, err := s.service.CreateMessage(s.context, errorMessage); err != nil {
		s.logger.Error("error while creating error message", "error", err)
	}

	s.AddToSendQueue(&event.AppInterviewErrorMessageSentEvent{
		AppBaseInterviewEvent: event.AppBaseInterviewEvent{Type: event.AppInterviewEventErrorMessageSentType},
		ErrorKey:              string(errorKey),
	})
}

func (s *appInterviewClientSessionImpl) handleClientMessage(rawMsg []byte) {
	var baseEvent event.AppBaseInterviewEvent
	if err := json.Unmarshal(rawMsg, &baseEvent); err != nil {
		s.logger.Error("error while unmarshalling message", "error", err)
		return
	}

	switch baseEvent.Type {
	case event.AppInterviewEventUserMessageSentType:
		s.handleUserMessageSentEvent(rawMsg)
	case event.AppInterviewEventUserCompleteInterviewType:
		s.handleUserCompleteInterviewEvent(rawMsg)
	}
}

func (s *appInterviewClientSessionImpl) handleUserMessageSentEvent(rawMsg []byte) {
	var userMessageEvent event.AppInterviewUserMessageSentEvent
	if err := json.Unmarshal(rawMsg, &userMessageEvent); err != nil {
		s.logger.Error("error while unmarshalling user message event", "error", err)
		s.sendErrorMessage(errors.InterviewMessageErrorKeyFailedToSendMessage)
		return
	}

	userMessage := &model.InterviewMessage{
		InterviewID: s.interview.ID,
		ContentText: userMessageEvent.Content,
		Role:        model.InterviewMessageRoleUser,
	}
	if _, err := s.service.CreateMessage(s.context, userMessage); err != nil {
		s.logger.Error("error while creating user message", "error", err)
		s.sendErrorMessage(errors.InterviewMessageErrorKeyFailedToSendMessage)
		return
	}

	s.AddToSendQueue(&event.AppInterviewInterviewerMessagePendingEvent{
		AppBaseInterviewEvent: event.AppBaseInterviewEvent{Type: event.AppInterviewEventInterviewerMessagePendingType},
	})

	interviewerMessage, err := s.service.ProcessMessageWithOpenAI(s.context, s.interview.ThreadID, userMessage)
	if err != nil {
		s.logger.Error("error while processing user message with OpenAI", "error", err)
		s.sendErrorMessage(errors.InterviewMessageErrorKeyFailedToProcessMessage)
		return
	}

	s.AddToSendQueue(&event.AppInterviewInterviewerMessageSentEvent{
		AppBaseInterviewEvent: event.AppBaseInterviewEvent{Type: event.AppInterviewEventInterviewerMessageSentType},
		Data:                  *mapper.MapInterviewMessageModelToAppDto(interviewerMessage),
	})

	if interviewerMessage.IsLastMessage == false {
		return
	}

	s.AddToSendQueue(&event.AppInterviewResultsPendingEvent{
		AppBaseInterviewEvent: event.AppBaseInterviewEvent{Type: event.AppInterviewEventResultsPendingType},
	})

	interviewResult, err := s.service.GetResultsWithOpenAI(s.context, s.interview.ID, s.interview.ThreadID)
	if err != nil {
		s.logger.Error("error while getting interview results", "error", err)
		s.sendErrorMessage(errors.InterviewMessageErrorKeyFailedToGetResults)
		return
	}

	s.AddToSendQueue(&event.AppInterviewResultsSentEvent{
		AppBaseInterviewEvent: event.AppBaseInterviewEvent{Type: event.AppInterviewEventResultsSentType},
		Data:                  *mapper.MapInterviewResultToAppDto(interviewResult),
	})

	if err := s.service.CompleteInterview(s.context, s.interview); err != nil {
		s.logger.Error("error while completing interview", "error", err)
		s.sendErrorMessage(errors.InterviewMessageErrorKeyFailedToCompleteInterview)
		return
	}

	s.AddToSendQueue(&event.AppInterviewCompletedEvent{
		AppBaseInterviewEvent: event.AppBaseInterviewEvent{Type: event.AppInterviewEventInterviewCompletedType},
	})

	if err := s.socket.Close(); err != nil {
		s.logger.Error("error while closing socket", "error", err)
		return
	}
}

func (s *appInterviewClientSessionImpl) handleUserCompleteInterviewEvent(rawMsg []byte) {
	s.logger.Info("handleUserCompleteInterviewEvent", "rawMsg", string(rawMsg))

	isAvailableToComplete, err := s.service.CheckInterviewCompleteAvailability(s.context, s.interview)
	if err != nil || !isAvailableToComplete {
		s.logger.Error("error while checking interview complete availability", "error", err)
		s.sendErrorMessage(errors.InterviewMessageErrorKeyFailedToCheckCompleteAvailability)
		return
	}

	s.AddToSendQueue(&event.AppInterviewResultsPendingEvent{
		AppBaseInterviewEvent: event.AppBaseInterviewEvent{Type: event.AppInterviewEventResultsPendingType},
	})

	interviewResult, err := s.service.GetResultsWithOpenAI(s.context, s.interview.ID, s.interview.ThreadID)
	if err != nil {
		s.logger.Error("error while getting interview results", "error", err)
		s.sendErrorMessage(errors.InterviewMessageErrorKeyFailedToGetResults)
		return
	}

	s.AddToSendQueue(&event.AppInterviewResultsSentEvent{
		AppBaseInterviewEvent: event.AppBaseInterviewEvent{Type: event.AppInterviewEventResultsSentType},
		Data:                  *mapper.MapInterviewResultToAppDto(interviewResult),
	})

	if err := s.service.CompleteInterview(s.context, s.interview); err != nil {
		s.logger.Error("error while completing interview", "error", err)
		s.sendErrorMessage(errors.InterviewMessageErrorKeyFailedToCompleteInterview)
		return
	}

	s.AddToSendQueue(&event.AppInterviewCompletedEvent{
		AppBaseInterviewEvent: event.AppBaseInterviewEvent{Type: event.AppInterviewEventInterviewCompletedType},
	})

	if err := s.socket.Close(); err != nil {
		s.logger.Error("error while closing socket", "error", err)
		return
	}
}
