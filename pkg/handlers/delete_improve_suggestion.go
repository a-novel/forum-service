package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteImproveSuggestionHandler interface {
	Handle(c *gin.Context)
}

func NewDeleteImproveSuggestionHandler(service services.DeleteImproveSuggestionService) DeleteImproveSuggestionHandler {
	return &deleteImproveSuggestionHandlerImpl{
		service: service,
	}
}

type deleteImproveSuggestionHandlerImpl struct {
	service services.DeleteImproveSuggestionService
}

func (h *deleteImproveSuggestionHandlerImpl) Handle(c *gin.Context) {
	token := c.GetHeader("Authorization")

	query := new(models.DeleteImproveSuggestionQuery)
	if err := c.BindQuery(query); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := h.service.Delete(c, token, query.ID.Value())
	if err != nil {
		errors.ErrorToHTTPCode(c, err, []errors.HTTPError{
			{services.ErrNotTheCreator, http.StatusUnauthorized},
			{errors.ErrInvalidCredentials, http.StatusForbidden},
		})
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
