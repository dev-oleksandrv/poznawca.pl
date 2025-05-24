package handler

import (
	"context"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/service"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/ws"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"
)

type AppWSInterviewHandler interface {
	RunInterview(c *gin.Context)
}

type appWSInterviewHandlerImpl struct {
	upgrader *websocket.Upgrader
	service  service.AppWSInterviewService
}

func NewAppWSInterviewHandler(service service.AppWSInterviewService) AppWSInterviewHandler {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 256,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	return &appWSInterviewHandlerImpl{
		upgrader: upgrader,
		service:  service,
	}
}

func (h *appWSInterviewHandlerImpl) RunInterview(c *gin.Context) {
	rawInterviewID := c.Query("interview_id")
	if rawInterviewID == "" {
		c.JSON(400, gin.H{"error": "interview_id is required"})
		return
	}

	if err := uuid.Validate(rawInterviewID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	interviewModel, err := h.service.ActivateInterview(context.Background(), uuid.MustParse(rawInterviewID))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to activate interview"})
		return
	}

	if interviewModel.Interviewer == nil {
		c.JSON(500, gin.H{"error": "interview is not assigned to an interviewer"})
		return
	}

	socket, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to upgrade connection"})
		return
	}

	clientSession := ws.NewAppInterviewClientSession(&ws.NewAppInterviewClientSessionConfig{
		Context:   c.Request.Context(),
		Socket:    socket,
		Interview: interviewModel,
		Service:   h.service,
	})

	defer func() {
		slog.Info("closing connection and updating interview status", "interview_id", interviewModel.ID)
		if err := h.service.AbandonInterview(c.Request.Context(), interviewModel); err != nil {
			slog.Error("failed to update interview status", "error", err)
		}
	}()

	go clientSession.Write()

	clientSession.Init()

	clientSession.Read()
}
