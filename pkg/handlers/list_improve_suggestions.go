package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListImproveSuggestionsHandler interface {
	Handle(c *gin.Context)
}

func NewListImproveSuggestionsHandler(service services.ListImproveSuggestionsService) ListImproveSuggestionsHandler {
	return &listImproveSuggestionsHandlerImpl{
		service: service,
	}
}

type listImproveSuggestionsHandlerImpl struct {
	service services.ListImproveSuggestionsService
}

func (h *listImproveSuggestionsHandlerImpl) Handle(c *gin.Context) {
	query := new(models.ListImproveSuggestionQuery)
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
