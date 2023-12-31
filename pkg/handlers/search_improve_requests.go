package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-apis"
	goframework "github.com/a-novel/go-framework"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchImproveRequestsHandler interface {
	Handle(c *gin.Context)
}

func NewSearchImproveRequestsHandler(service services.SearchImproveRequestsService) SearchImproveRequestsHandler {
	return &searchImproveRequestsHandlerImpl{
		service: service,
	}
}

type searchImproveRequestsHandlerImpl struct {
	service services.SearchImproveRequestsService
}

func (h *searchImproveRequestsHandlerImpl) Handle(c *gin.Context) {
	query := new(models.SearchImproveRequestsQuery)
	if err := c.BindQuery(query); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	previews, total, err := h.service.Search(c, *query)
	if err != nil {
		apis.ErrorToHTTPCode(c, err, []apis.HTTPError{
			{goframework.ErrInvalidEntity, http.StatusBadRequest},
		}, false)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"res":   previews,
		"total": total,
	})
}
