package handlers

import (
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-apis"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VoteImproveSuggestionHandler interface {
	Handle(c *gin.Context)
}

func NewVoteImproveSuggestionHandler(service services.VoteImproveSuggestionService) VoteImproveSuggestionHandler {
	return &voteImproveSuggestionHandlerImpl{
		service: service,
	}
}

type voteImproveSuggestionHandlerImpl struct {
	service services.VoteImproveSuggestionService
}

func (h *voteImproveSuggestionHandlerImpl) Handle(c *gin.Context) {
	form := new(models.UpdateImproveSuggestionVotesForm)
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
