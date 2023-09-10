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

type CreateImproveSuggestionHandler interface {
	Handle(c *gin.Context)
}

func NewCreateImproveSuggestionHandler(service services.CreateImproveSuggestionService) CreateImproveSuggestionHandler {
	return &createImproveSuggestionHandlerImpl{
		service: service,
	}
}

type createImproveSuggestionHandlerImpl struct {
	service services.CreateImproveSuggestionService
}

func (h *createImproveSuggestionHandlerImpl) Handle(c *gin.Context) {
	token := c.GetHeader("Authorization")

	form := new(models.ImproveSuggestionForm)
	if err := c.BindJSON(form); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.service.Create(c, token, form, uuid.New(), time.Now())
	if err != nil {
		errors.ErrorToHTTPCode(c, err, []errors.HTTPError{
			{errors.ErrInvalidCredentials, http.StatusForbidden},
			{errors.ErrNotFound, http.StatusNotFound},
			{errors.ErrInvalidEntity, http.StatusUnprocessableEntity},
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}
