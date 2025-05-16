package handler

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/backoffice-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/modules/backoffice-proxy/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InterviewerHandler interface {
	GetList(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type interviewerHandlerImpl struct {
	interviewerService service.InterviewerService
}

func NewInterviewerHandler(interviewerService service.InterviewerService) InterviewerHandler {
	return &interviewerHandlerImpl{
		interviewerService: interviewerService,
	}
}

func (h *interviewerHandlerImpl) GetList(ctx *gin.Context) {
	interviewers, err := h.interviewerService.FindAll(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": interviewers})
}

func (h *interviewerHandlerImpl) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := uuid.Validate(id); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid interviewer id"})
		return
	}

	interviewer, err := h.interviewerService.FindByID(ctx, uuid.MustParse(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": interviewer})
}

func (h *interviewerHandlerImpl) Create(ctx *gin.Context) {
	var input *dto.InterviewerInputDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	interviewer, err := h.interviewerService.Create(ctx, input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"data": interviewer})
}

func (h *interviewerHandlerImpl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := uuid.Validate(id); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid interviewer id"})
		return
	}

	var input *dto.InterviewerInputDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	interviewer, err := h.interviewerService.Update(ctx, uuid.MustParse(id), input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": interviewer})
}

func (h *interviewerHandlerImpl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := uuid.Validate(id); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid interviewer id"})
		return
	}

	err := h.interviewerService.Delete(ctx, uuid.MustParse(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, nil)
}
