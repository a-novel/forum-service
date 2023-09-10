package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type UpdateImproveSuggestionHandler interface {
	Handle(c *gin.Context)
}

func NewUpdateImproveSuggestionHandler(service services.UpdateImproveSuggestionService) UpdateImproveSuggestionHandler {
	return &updateImproveSuggestionHandlerImpl{
		service: service,
	}
}

type updateImproveSuggestionHandlerImpl struct {
	service services.UpdateImproveSuggestionService
}

func (h *updateImproveSuggestionHandlerImpl) Handle(c *gin.Context) {
	token := c.GetHeader("Authorization")

	form := new(models.ImproveSuggestionForm)
	if err := c.BindJSON(form); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.service.Update(c, token, form, uuid.New(), time.Now())
	if err != nil {
		errors.ErrorToHTTPCode(c, err, []errors.HTTPError{
			{services.ErrSwitchSource, http.StatusUnauthorized},
			{errors.ErrInvalidCredentials, http.StatusForbidden},
			{errors.ErrNotFound, http.StatusNotFound},
			{errors.ErrInvalidEntity, http.StatusUnprocessableEntity},
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}
