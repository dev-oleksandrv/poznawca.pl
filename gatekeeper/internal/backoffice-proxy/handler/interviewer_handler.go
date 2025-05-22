package handler

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/backoffice-proxy/dto"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/backoffice-proxy/mapper"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/backoffice-proxy/service"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/errors"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/shared/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InterviewerHandler interface {
	GetList(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type interviewerHandler struct {
	interviewerService service.InterviewerService
}

func NewInterviewerHandler(interviewerService service.InterviewerService) InterviewerHandler {
	return &interviewerHandler{
		interviewerService: interviewerService,
	}
}

func (h *interviewerHandler) GetList(c *gin.Context) {
	interviewers, err := h.interviewerService.FindAll(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get interviewers"})
		return
	}

	c.JSON(200, gin.H{"data": mapper.MapInterviewerModelToDtoList(interviewers)})
}

func (h *interviewerHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(400, gin.H{"error": errors.ErrInvalidID})
		return
	}

	interviewer, err := h.interviewerService.FindByID(c, uuid.MustParse(id))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get interviewer"})
		return
	}

	if interviewer == nil {
		c.JSON(404, gin.H{"error": "interviewer not found"})
		return
	}

	c.JSON(200, gin.H{"data": mapper.MapInterviewerModelToDto(interviewer)})
}

func (h *interviewerHandler) Create(c *gin.Context) {
	var inputDto *dto.CreateInterviewerRequestDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	if err := inputDto.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	interviewer := &model.InterviewerModel{
		Name:                      inputDto.Name,
		AvatarURL:                 inputDto.AvatarURL,
		EntryMessage:              inputDto.EntryMessage,
		Description:               inputDto.Description,
		DescriptionTranslationKey: inputDto.DescriptionTranslationKey,
	}

	createdInterviewer, err := h.interviewerService.Create(c, interviewer)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create interviewer"})
		return
	}

	c.JSON(201, gin.H{"data": mapper.MapInterviewerModelToDto(createdInterviewer)})
}

func (h *interviewerHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(400, gin.H{"error": errors.ErrInvalidID})
		return
	}

	var inputDto *dto.UpdateInterviewerRequestDto
	if err := c.ShouldBindJSON(&inputDto); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	if err := inputDto.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	interviewer, err := h.interviewerService.FindByID(c, uuid.MustParse(id))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get interviewer"})
		return
	}

	if interviewer == nil {
		c.JSON(404, gin.H{"error": "interviewer not found"})
		return
	}

	if inputDto.Name != nil {
		interviewer.Name = *inputDto.Name
	}
	if inputDto.AvatarURL != nil {
		interviewer.AvatarURL = *inputDto.AvatarURL
	}
	if inputDto.EntryMessage != nil {
		interviewer.EntryMessage = *inputDto.EntryMessage
	}
	if inputDto.Description != nil {
		interviewer.Description = *inputDto.Description
	}
	if inputDto.DescriptionTranslationKey != nil {
		interviewer.DescriptionTranslationKey = *inputDto.DescriptionTranslationKey
	}

	updatedInterviewer, err := h.interviewerService.Update(c, interviewer)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to update interviewer"})
		return
	}

	c.JSON(200, gin.H{"data": mapper.MapInterviewerModelToDto(updatedInterviewer)})
}

func (h *interviewerHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		c.JSON(400, gin.H{"error": errors.ErrInvalidID})
		return
	}

	err := h.interviewerService.Delete(c, uuid.MustParse(id))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to delete interviewer"})
		return
	}

	c.JSON(204, nil)
}
