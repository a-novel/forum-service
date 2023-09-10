package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteImproveRequestRevisionHandler interface {
	Handle(c *gin.Context)
}

func NewDeleteImproveRequestRevisionHandler(service services.DeleteImproveRequestRevisionService) DeleteImproveRequestRevisionHandler {
	return &deleteImproveRequestRevisionHandlerImpl{
		service: service,
	}
}

type deleteImproveRequestRevisionHandlerImpl struct {
	service services.DeleteImproveRequestRevisionService
}

func (h *deleteImproveRequestRevisionHandlerImpl) Handle(c *gin.Context) {
	token := c.GetHeader("Authorization")

	query := new(models.DeleteImproveRequestRevisionQuery)
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
