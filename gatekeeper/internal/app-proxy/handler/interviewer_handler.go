package handler

import (
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/mapper"
	"github.com/dev-oleksandrv/poznawca/gatekeeper/internal/app-proxy/service"
	"github.com/gin-gonic/gin"
)

type AppInterviewerHandler interface {
	GetList(c *gin.Context)
}

type appInterviewerHandlerImpl struct {
	interviewerService service.AppInterviewerService
}

func NewAppInterviewerHandler(interviewerService service.AppInterviewerService) AppInterviewerHandler {
	return &appInterviewerHandlerImpl{
		interviewerService: interviewerService,
	}
}

func (h *appInterviewerHandlerImpl) GetList(c *gin.Context) {
	interviewers, err := h.interviewerService.FindAll(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get interviewers"})
		return
	}

	c.JSON(200, gin.H{"data": mapper.MapInterviewerModelToAppDtoList(interviewers)})
}
