package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetImproveRequestHandler interface {
	Handle(c *gin.Context)
}

func NewGetImproveRequestHandler(service services.GetImproveRequestService) GetImproveRequestHandler {
	return &getImproveRequestHandlerImpl{
		service: service,
	}
}

type getImproveRequestHandlerImpl struct {
	service services.GetImproveRequestService
}

func (h *getImproveRequestHandlerImpl) Handle(c *gin.Context) {
	query := new(models.GetImproveRequestQuery)
	if err := c.BindQuery(query); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	request, err := h.service.Get(c, query.ID.Value())
	if err != nil {
		errors.ErrorToHTTPCode(c, err, []errors.HTTPError{
			{errors.ErrNotFound, http.StatusNotFound},
		})
		return
	}

	c.JSON(http.StatusOK, request)
}
