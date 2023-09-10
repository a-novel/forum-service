package handlers

import (
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/errors"
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
		errors.ErrorToHTTPCode(c, err, []errors.HTTPError{
			{services.ErrTheCreator, http.StatusUnauthorized},
			{errors.ErrNotFound, http.StatusNotFound},
		})
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
