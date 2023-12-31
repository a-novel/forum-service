package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/handlers"
	"github.com/a-novel/forum-service/pkg/models"
	servicesmocks "github.com/a-novel/forum-service/pkg/services/mocks"
	goframework "github.com/a-novel/go-framework"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateImproveSuggestionHandler(t *testing.T) {
	data := []struct {
		name string

		authorization string

		body interface{}

		shouldCallService bool
		serviceResp       *models.ImproveSuggestion
		serviceErr        error

		expect       interface{}
		expectStatus int
	}{
		{
			name:          "Success",
			authorization: "Bearer my-token",
			body: map[string]interface{}{
				"title":     "title",
				"content":   "content",
				"requestID": goframework.NumberUUID(1).String(),
			},
			shouldCallService: true,
			serviceResp: &models.ImproveSuggestion{
				ID:        goframework.NumberUUID(1),
				CreatedAt: baseTime,
				RequestID: goframework.NumberUUID(1),
				SourceID:  goframework.NumberUUID(10),
				UserID:    goframework.NumberUUID(100),
				Title:     "title",
				Content:   "content",
			},
			expect: map[string]interface{}{
				"id":        goframework.NumberUUID(1).String(),
				"createdAt": baseTime.Format(time.RFC3339),
				"updatedAt": nil,
				"requestID": goframework.NumberUUID(1).String(),
				"sourceID":  goframework.NumberUUID(10).String(),
				"userID":    goframework.NumberUUID(100).String(),
				"title":     "title",
				"content":   "content",
				"upVotes":   float64(0),
				"downVotes": float64(0),
				"validated": false,
			},
			expectStatus: http.StatusCreated,
		},
		{
			name:          "Error/ErrInvalidCredentials",
			authorization: "Bearer my-token",
			body: map[string]interface{}{
				"title":     "title",
				"content":   "content",
				"requestID": goframework.NumberUUID(1).String(),
			},
			shouldCallService: true,
			serviceErr:        goframework.ErrInvalidCredentials,
			expectStatus:      http.StatusForbidden,
		},
		{
			name:          "Error/ErrInvalidEntity",
			authorization: "Bearer my-token",
			body: map[string]interface{}{
				"title":     "title",
				"content":   "content",
				"requestID": goframework.NumberUUID(1).String(),
			},
			shouldCallService: true,
			serviceErr:        goframework.ErrInvalidEntity,
			expectStatus:      http.StatusUnprocessableEntity,
		},
		{
			name:          "Error/ErrNotFound",
			authorization: "Bearer my-token",
			body: map[string]interface{}{
				"title":     "title",
				"content":   "content",
				"requestID": goframework.NumberUUID(1).String(),
			},
			shouldCallService: true,
			serviceErr:        bunovel.ErrNotFound,
			expectStatus:      http.StatusNotFound,
		},
		{
			name:          "Error/BadRequest",
			authorization: "Bearer my-token",
			body: map[string]interface{}{
				"title":     "title",
				"content":   "content",
				"requestID": "fake uuid",
			},
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewCreateImproveSuggestionService(t)

			mrshBody, err := json.Marshal(d.body)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/", bytes.NewReader(mrshBody))
			c.Request.Header.Set("Authorization", d.authorization)

			if d.shouldCallService {
				service.
					On("Create", c, d.authorization, mock.Anything, mock.Anything, mock.Anything).
					Return(d.serviceResp, d.serviceErr)
			}

			handler := handlers.NewCreateImproveSuggestionHandler(service)
			handler.Handle(c)

			require.Equal(t, d.expectStatus, w.Code, c.Errors.String())
			if d.expect != nil {
				var body interface{}
				require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
				require.Equal(t, d.expect, body)
			}

			service.AssertExpectations(t)
		})
	}
}
