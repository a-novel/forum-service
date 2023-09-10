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

type CreateImproveRequestHandler interface {
	Handle(c *gin.Context)
}

func NewCreateImproveRequestHandler(service services.CreateImproveRequestService) CreateImproveRequestHandler {
	return &createImproveRequestHandlerImpl{
		service: service,
	}
}

type createImproveRequestHandlerImpl struct {
	service services.CreateImproveRequestService
}

func (h *createImproveRequestHandlerImpl) Handle(c *gin.Context) {
	token := c.GetHeader("Authorization")

	form := new(models.CreateImproveRequestForm)
	if err := c.BindJSON(form); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.service.Create(c, token, form.Title, form.Content, form.SourceID, uuid.New(), time.Now())
	if err != nil {
		errors.ErrorToHTTPCode(c, err, []errors.HTTPError{
			{services.ErrNotTheCreator, http.StatusUnauthorized},
			{errors.ErrInvalidCredentials, http.StatusForbidden},
			{errors.ErrInvalidEntity, http.StatusUnprocessableEntity},
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}
