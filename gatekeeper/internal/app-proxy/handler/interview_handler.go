package handler

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/mapper"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/service"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/errors"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AppInterviewHandler interface {
	GetByID(c *gin.Context)
	Create(c *gin.Context)
}

type appInterviewHandlerImpl struct {
	interviewService   service.AppInterviewService
	interviewerService service.AppInterviewerService
}

func NewAppInterviewHandler(interviewService service.AppInterviewService, interviewerService service.AppInterviewerService) AppInterviewHandler {
	return &appInterviewHandlerImpl{
		interviewService:   interviewService,
		interviewerService: interviewerService,
	}
}

func (h *appInterviewHandlerImpl) GetByID(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	interview, err := h.interviewService.FindByID(c.Request.Context(), uuid.MustParse(id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": mapper.MapInterviewModelToAppDto(interview)})
}

func (h *appInterviewHandlerImpl) Create(c *gin.Context) {
	var inputDto *dto.CreateAppInterviewRequestDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := inputDto.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var interviewerID *uuid.UUID
	if inputDto.InterviewerID != nil {
		parsedID, err := uuid.Parse(*inputDto.InterviewerID)
		if err != nil {
			c.JSON(400, gin.H{"error": errors.ErrInvalidID.Error()})
			return
		}

		interviewerID = &parsedID
	}

	interviewer, err := h.interviewerService.FindByIDOrRandom(c.Request.Context(), interviewerID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	interview, err := h.interviewService.Create(c.Request.Context(), &model.Interview{
		Status:        model.InterviewStatusPending,
		InterviewerID: &interviewer.ID,
		Interviewer:   interviewer,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": mapper.MapInterviewModelToAppDto(interview)})
}
