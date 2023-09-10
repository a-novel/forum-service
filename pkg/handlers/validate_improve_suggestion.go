package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ValidateImproveSuggestionHandler interface {
	Handle(c *gin.Context)
}

func NewValidateImproveSuggestionHandler(service services.ValidateImproveSuggestionService) ValidateImproveSuggestionHandler {
	return &validateImproveSuggestionHandlerImpl{
		service: service,
	}
}

type validateImproveSuggestionHandlerImpl struct {
	service services.ValidateImproveSuggestionService
}

func (h *validateImproveSuggestionHandlerImpl) Handle(c *gin.Context) {
	token := c.GetHeader("Authorization")

	form := new(models.ValidateImproveSuggestionForm)
	if err := c.BindJSON(form); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.service.Validate(c, token, form.Validated, form.ID); err != nil {
		errors.ErrorToHTTPCode(c, err, []errors.HTTPError{
			{services.ErrNotTheCreator, http.StatusUnauthorized},
			{errors.ErrInvalidCredentials, http.StatusForbidden},
		})
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
