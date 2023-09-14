package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-apis"
	goframework "github.com/a-novel/go-framework"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteImproveRequestHandler interface {
	Handle(c *gin.Context)
}

func NewDeleteImproveRequestHandler(service services.DeleteImproveRequestService) DeleteImproveRequestHandler {
	return &deleteImproveRequestHandlerImpl{
		service: service,
	}
}

type deleteImproveRequestHandlerImpl struct {
	service services.DeleteImproveRequestService
}

func (h *deleteImproveRequestHandlerImpl) Handle(c *gin.Context) {
	token := c.GetHeader("Authorization")

	query := new(models.DeleteImproveRequestQuery)
	if err := c.BindQuery(query); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := h.service.Delete(c, token, query.ID.Value())
	if err != nil {
		apis.ErrorToHTTPCode(c, err, []apis.HTTPError{
			{services.ErrNotTheCreator, http.StatusUnauthorized},
			{goframework.ErrInvalidCredentials, http.StatusForbidden},
		}, false)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
