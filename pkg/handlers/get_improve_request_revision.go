package handlers

import (
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-apis"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetImproveRequestRevisionHandler interface {
	Handle(c *gin.Context)
}

func NewGetImproveRequestRevisionHandler(service services.GetImproveRequestRevisionService) GetImproveRequestRevisionHandler {
	return &getImproveRequestRevisionHandlerImpl{
		service: service,
	}
}

type getImproveRequestRevisionHandlerImpl struct {
	service services.GetImproveRequestRevisionService
}

func (h *getImproveRequestRevisionHandlerImpl) Handle(c *gin.Context) {
	query := new(models.GetImproveRequestRevisionQuery)
	if err := c.BindQuery(query); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	revision, err := h.service.Get(c, query.ID.Value())
	if err != nil {
		apis.ErrorToHTTPCode(c, err, []apis.HTTPError{
			{bunovel.ErrNotFound, http.StatusNotFound},
		}, false)
		return
	}

	c.JSON(http.StatusOK, revision)
}
