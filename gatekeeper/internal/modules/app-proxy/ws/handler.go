package ws

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/model"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"
)

type Handler interface {
	RunInterview(c *gin.Context)
}

type handlerImpl struct {
	upgrader         *websocket.Upgrader
	interviewService service.InterviewService
}

func NewHandler(interviewService service.InterviewService) Handler {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 256,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	return &handlerImpl{
		upgrader:         upgrader,
		interviewService: interviewService,
	}
}

func (h *handlerImpl) RunInterview(c *gin.Context) {
	rawInterviewID := c.Query("interview_id")
	if rawInterviewID == "" {
		c.JSON(400, gin.H{"error": "interview_id is required"})
		return
	}

	if err := uuid.Validate(rawInterviewID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	interviewModel, err := h.interviewService.FindByID(c.Request.Context(), uuid.MustParse(rawInterviewID))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to find interview"})
		return
	}

	if interviewModel == nil {
		c.JSON(404, gin.H{"error": "interview not found"})
		return
	}

	if interviewModel.Status != model.InterviewStatusPending {
		c.JSON(400, gin.H{"error": "interview is not in pending status"})
		return
	}

	if err := h.interviewService.UpdateStatus(c.Request.Context(), &dto.UpdateInterviewStatusInputDto{
		InterviewID: interviewModel.ID,
		Status:      model.InterviewStatusActive,
	}); err != nil {
		c.JSON(500, gin.H{"error": "failed to update interview status"})
		return
	}

	socket, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to upgrade connection"})
		return
	}

	session := NewClientSession(socket, h.interviewService, interviewModel)

	defer func() {
		slog.Info("closing connection and updating interview status", "interview_id", interviewModel.ID)
		err := h.interviewService.UpdateStatus(c.Request.Context(), &dto.UpdateInterviewStatusInputDto{
			InterviewID: interviewModel.ID,
			Status:      model.InterviewStatusAbandoned,
		})

		if err != nil {
			slog.Error("failed to update interview status", "error", err)
		}
	}()

	go session.Write()

	session.AddToSendQueue(&SystemMessagePendingEvent{
		BaseEvent: BaseEvent{Type: SystemMessagePendingEventType},
	})

	session.SaveInitialMessage()

	session.Read()
}
