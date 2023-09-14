package handlers

import (
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-apis"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VoteImproveRequestHandler interface {
	Handle(c *gin.Context)
}

func NewVoteImproveRequestHandler(service services.VoteImproveRequestService) VoteImproveRequestHandler {
	return &voteImproveRequestHandlerImpl{
		service: service,
	}
}

type voteImproveRequestHandlerImpl struct {
	service services.VoteImproveRequestService
}

func (h *voteImproveRequestHandlerImpl) Handle(c *gin.Context) {
	form := new(models.UpdateImproveRequestVotesForm)
	if err := c.BindJSON(form); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.service.Vote(c, form.ID, form.UserID, form.UpVotes, form.DownVotes); err != nil {
		apis.ErrorToHTTPCode(c, err, []apis.HTTPError{
			{services.ErrTheCreator, http.StatusUnauthorized},
			{bunovel.ErrNotFound, http.StatusNotFound},
		}, false)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
