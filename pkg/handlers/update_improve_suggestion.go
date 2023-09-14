package handlers

import (
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-apis"
	goframework "github.com/a-novel/go-framework"
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
		apis.ErrorToHTTPCode(c, err, []apis.HTTPError{
			{services.ErrSwitchSource, http.StatusUnauthorized},
			{goframework.ErrInvalidCredentials, http.StatusForbidden},
			{bunovel.ErrNotFound, http.StatusNotFound},
			{goframework.ErrInvalidEntity, http.StatusUnprocessableEntity},
		}, false)
		return
	}

	c.JSON(http.StatusCreated, res)
}
