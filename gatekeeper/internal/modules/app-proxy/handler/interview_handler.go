package handler

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/app-proxy/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InterviewHandler interface {
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type interviewHandlerImpl struct {
	interviewService service.InterviewService
}

func NewInterviewHandler(interviewService service.InterviewService) InterviewHandler {
	return &interviewHandlerImpl{
		interviewService: interviewService,
	}
}

func (h *interviewHandlerImpl) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := uuid.Validate(id); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid interview id"})
		return
	}

	interview, err := h.interviewService.FindByID(ctx.Request.Context(), uuid.MustParse(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if interview == nil {
		ctx.JSON(404, gin.H{"error": "interview not found"})
		return
	}

	ctx.JSON(200, gin.H{"data": interview})
}

func (h *interviewHandlerImpl) Create(ctx *gin.Context) {
	interview, err := h.interviewService.Create(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": interview})
}
