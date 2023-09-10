package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListImproveRequestsHandler interface {
	Handle(c *gin.Context)
}

func NewListImproveRequestsHandler(service services.ListImproveRequestsService) ListImproveRequestsHandler {
	return &listImproveRequestsHandlerImpl{
		service: service,
	}
}

type listImproveRequestsHandlerImpl struct {
	service services.ListImproveRequestsService
}

func (h *listImproveRequestsHandlerImpl) Handle(c *gin.Context) {
	query := new(models.ListImproveRequestsQuery)
	if err := c.BindQuery(query); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	previews, err := h.service.List(c, query.IDs.Value())
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"previews": previews})
}
