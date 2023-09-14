package handlers_test

import (
	"encoding/json"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/handlers"
	"github.com/a-novel/forum-service/pkg/models"
	servicesmocks "github.com/a-novel/forum-service/pkg/services/mocks"
	goframework "github.com/a-novel/go-framework"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetImproveSuggestionHandler(t *testing.T) {
	data := []struct {
		name string

		query string

		shouldCallService       bool
		shouldCallServiceWithID uuid.UUID
		serviceResp             *models.ImproveSuggestion
		serviceErr              error

		expect       interface{}
		expectStatus int
	}{
		{
			name:                    "Success",
			query:                   "?id=01010101-0101-0101-0101-010101010101",
			shouldCallService:       true,
			shouldCallServiceWithID: goframework.NumberUUID(1),
			serviceResp: &models.ImproveSuggestion{
				ID:        goframework.NumberUUID(1),
				CreatedAt: baseTime,
				UpdatedAt: &updateTime,
				SourceID:  goframework.NumberUUID(10),
				UserID:    goframework.NumberUUID(100),
				Validated: true,
				UpVotes:   128,
				DownVotes: 64,
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
			expect: map[string]interface{}{
				"id":        goframework.NumberUUID(1).String(),
				"createdAt": baseTime.Format(time.RFC3339),
				"updatedAt": updateTime.Format(time.RFC3339),
				"sourceID":  goframework.NumberUUID(10).String(),
				"userID":    goframework.NumberUUID(100).String(),
				"requestID": goframework.NumberUUID(1).String(),
				"validated": true,
				"title":     "suggestion title",
				"content":   "suggestion content",
				"upVotes":   float64(128),
				"downVotes": float64(64),
			},
			expectStatus: http.StatusOK,
		},
		{
			name:                    "Errors/NotFound",
			query:                   "?id=01010101-0101-0101-0101-010101010101",
			shouldCallService:       true,
			shouldCallServiceWithID: goframework.NumberUUID(1),
			serviceErr:              bunovel.ErrNotFound,
			expectStatus:            http.StatusNotFound,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewGetImproveSuggestionService(t)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/"+d.query, nil)

			if d.shouldCallService {
				service.
					On("Get", c, d.shouldCallServiceWithID).
					Return(d.serviceResp, d.serviceErr)
			}

			handler := handlers.NewGetImproveSuggestionHandler(service)
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
