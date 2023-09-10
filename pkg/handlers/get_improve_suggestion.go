package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetImproveSuggestionHandler interface {
	Handle(c *gin.Context)
}

func NewGetImproveSuggestionHandler(service services.GetImproveSuggestionService) GetImproveSuggestionHandler {
	return &getImproveSuggestionHandlerImpl{
		service: service,
	}
}

type getImproveSuggestionHandlerImpl struct {
	service services.GetImproveSuggestionService
}

func (h *getImproveSuggestionHandlerImpl) Handle(c *gin.Context) {
	query := new(models.GetImproveSuggestionQuery)
	if err := c.BindQuery(query); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	suggestion, err := h.service.Get(c, query.ID.Value())
	if err != nil {
		errors.ErrorToHTTPCode(c, err, []errors.HTTPError{
			{errors.ErrNotFound, http.StatusNotFound},
		})
		return
	}

	c.JSON(http.StatusOK, suggestion)
}
