package handlers

import (
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-apis"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListImproveRequestRevisionsHandler interface {
	Handle(c *gin.Context)
}

func NewListImproveRequestRevisionsHandler(service services.ListImproveRequestRevisionsService) ListImproveRequestRevisionsHandler {
	return &listImproveRequestRevisionsHandlerImpl{
		service: service,
	}
}

type listImproveRequestRevisionsHandlerImpl struct {
	service services.ListImproveRequestRevisionsService
}

func (h *listImproveRequestRevisionsHandlerImpl) Handle(c *gin.Context) {
	query := new(models.ListImproveRequestRevisionsQuery)
	if err := c.BindQuery(query); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	revisions, err := h.service.List(c, query.ID.Value())
	if err != nil {
		apis.ErrorToHTTPCode(c, err, []apis.HTTPError{
			{bunovel.ErrNotFound, http.StatusNotFound},
		}, false)
		return
	}

	c.JSON(http.StatusOK, gin.H{"revisions": revisions})
}
