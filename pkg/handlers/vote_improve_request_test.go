package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/handlers"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	servicesmocks "github.com/a-novel/forum-service/pkg/services/mocks"
	goframework "github.com/a-novel/go-framework"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVoteImproveRequestHandler(t *testing.T) {
	data := []struct {
		name string

		body interface{}

		shouldCallService     bool
		shouldCallServiceWith models.UpdateImproveRequestVotesForm
		serviceErr            error

		expectStatus int
	}{
		{
			name: "Success",
			body: map[string]interface{}{
				"id":        goframework.NumberUUID(1).String(),
				"userID":    goframework.NumberUUID(2).String(),
				"upVotes":   10,
				"downVotes": 5,
			},
			shouldCallService: true,
			shouldCallServiceWith: models.UpdateImproveRequestVotesForm{
				ID:        goframework.NumberUUID(1),
				UserID:    goframework.NumberUUID(2),
				UpVotes:   10,
				DownVotes: 5,
			},
			expectStatus: http.StatusNoContent,
		},
		{
			name: "Error/ErrNotFound",
			body: map[string]interface{}{
				"id":        goframework.NumberUUID(1).String(),
				"userID":    goframework.NumberUUID(2).String(),
				"upVotes":   10,
				"downVotes": 5,
			},
			shouldCallService: true,
			shouldCallServiceWith: models.UpdateImproveRequestVotesForm{
				ID:        goframework.NumberUUID(1),
				UserID:    goframework.NumberUUID(2),
				UpVotes:   10,
				DownVotes: 5,
			},
			serviceErr:   bunovel.ErrNotFound,
			expectStatus: http.StatusNotFound,
		},
		{
			name: "Error/ErrTheCreator",
			body: map[string]interface{}{
				"id":        goframework.NumberUUID(1).String(),
				"userID":    goframework.NumberUUID(2).String(),
				"upVotes":   10,
				"downVotes": 5,
			},
			shouldCallService: true,
			shouldCallServiceWith: models.UpdateImproveRequestVotesForm{
				ID:        goframework.NumberUUID(1),
				UserID:    goframework.NumberUUID(2),
				UpVotes:   10,
				DownVotes: 5,
			},
			serviceErr:   services.ErrTheCreator,
			expectStatus: http.StatusUnauthorized,
		},
		{
			name: "Error/BadFor,",
			body: map[string]interface{}{
				"id":        "fake uuid",
				"userID":    goframework.NumberUUID(2).String(),
				"upVotes":   10,
				"downVotes": 5,
			},
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewVoteImproveRequestService(t)

			mrshBody, err := json.Marshal(d.body)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/", bytes.NewReader(mrshBody))

			if d.shouldCallService {
				service.
					On(
						"Vote", c,
						d.shouldCallServiceWith.ID,
						d.shouldCallServiceWith.UserID,
						d.shouldCallServiceWith.UpVotes,
						d.shouldCallServiceWith.DownVotes,
					).
					Return(d.serviceErr)
			}

			handler := handlers.NewVoteImproveRequestHandler(service)
			handler.Handle(c)

			require.Equal(t, d.expectStatus, w.Code, c.Errors.String())

			service.AssertExpectations(t)
		})
	}
}
