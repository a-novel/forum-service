package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-apis"
	goframework "github.com/a-novel/go-framework"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchImproveSuggestionsHandler interface {
	Handle(c *gin.Context)
}

func NewSearchImproveSuggestionsHandler(service services.SearchImproveSuggestionsService) SearchImproveSuggestionsHandler {
	return &searchImproveSuggestionsHandlerImpl{
		service: service,
	}
}

type searchImproveSuggestionsHandlerImpl struct {
	service services.SearchImproveSuggestionsService
}

func (h *searchImproveSuggestionsHandlerImpl) Handle(c *gin.Context) {
	query := new(models.SearchImproveSuggestionsQuery)
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
