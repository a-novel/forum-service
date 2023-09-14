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

func TestGetImproveRequestHandler(t *testing.T) {
	data := []struct {
		name string

		query string

		shouldCallService       bool
		shouldCallServiceWithID uuid.UUID
		serviceResp             *models.ImproveRequestPreview
		serviceErr              error

		expect       interface{}
		expectStatus int
	}{
		{
			name:                    "Success",
			query:                   "?id=01010101-0101-0101-0101-010101010101",
			shouldCallService:       true,
			shouldCallServiceWithID: goframework.NumberUUID(1),
			serviceResp: &models.ImproveRequestPreview{
				ID:                       goframework.NumberUUID(1),
				CreatedAt:                baseTime,
				UserID:                   goframework.NumberUUID(100),
				Title:                    "title",
				Content:                  "content",
				UpVotes:                  128,
				DownVotes:                64,
				SuggestionsCount:         12,
				AcceptedSuggestionsCount: 6,
				RevisionCount:            4,
			},
			expect: map[string]interface{}{
				"id":                       goframework.NumberUUID(1).String(),
				"createdAt":                baseTime.Format(time.RFC3339),
				"userID":                   goframework.NumberUUID(100).String(),
				"title":                    "title",
				"content":                  "content",
				"upVotes":                  float64(128),
				"downVotes":                float64(64),
				"suggestionsCount":         float64(12),
				"acceptedSuggestionsCount": float64(6),
				"revisionsCount":           float64(4),
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
			service := servicesmocks.NewGetImproveRequestService(t)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/"+d.query, nil)

			if d.shouldCallService {
				service.
					On("Get", c, d.shouldCallServiceWithID).
					Return(d.serviceResp, d.serviceErr)
			}

			handler := handlers.NewGetImproveRequestHandler(service)
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
