package handlers_test

import (
	"encoding/json"
	"github.com/a-novel/forum-service/pkg/handlers"
	"github.com/a-novel/forum-service/pkg/models"
	servicesmocks "github.com/a-novel/forum-service/pkg/services/mocks"
	"github.com/a-novel/go-framework/errors"
	"github.com/a-novel/go-framework/test"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestListImproveRequestRevisionsHandler(t *testing.T) {
	data := []struct {
		name string

		query string

		shouldCallService       bool
		shouldCallServiceWithID uuid.UUID
		serviceResp             []*models.ImproveRequestRevisionPreview
		serviceErr              error

		expect       interface{}
		expectStatus int
	}{
		{
			name:                    "Success",
			query:                   "?id=01010101-0101-0101-0101-010101010101",
			shouldCallService:       true,
			shouldCallServiceWithID: test.NumberUUID(1),
			serviceResp: []*models.ImproveRequestRevisionPreview{
				{
					ID:                       test.NumberUUID(1),
					CreatedAt:                baseTime,
					SuggestionsCount:         10,
					AcceptedSuggestionsCount: 5,
				},
				{
					ID:                       test.NumberUUID(2),
					CreatedAt:                baseTime,
					SuggestionsCount:         8,
					AcceptedSuggestionsCount: 4,
				},
			},
			expect: map[string]interface{}{
				"revisions": []interface{}{
					map[string]interface{}{
						"id":                       test.NumberUUID(1).String(),
						"createdAt":                baseTime.Format(time.RFC3339),
						"suggestionsCount":         float64(10),
						"acceptedSuggestionsCount": float64(5),
					},
					map[string]interface{}{
						"id":                       test.NumberUUID(2).String(),
						"createdAt":                baseTime.Format(time.RFC3339),
						"suggestionsCount":         float64(8),
						"acceptedSuggestionsCount": float64(4),
					},
				},
			},
			expectStatus: http.StatusOK,
		},
		{
			name:                    "Errors/NotFound",
			query:                   "?id=01010101-0101-0101-0101-010101010101",
			shouldCallService:       true,
			shouldCallServiceWithID: test.NumberUUID(1),
			serviceErr:              errors.ErrNotFound,
			expectStatus:            http.StatusNotFound,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewListImproveRequestRevisionsService(t)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/"+d.query, nil)

			if d.shouldCallService {
				service.
					On("List", c, d.shouldCallServiceWithID).
					Return(d.serviceResp, d.serviceErr)
			}

			handler := handlers.NewListImproveRequestRevisionsHandler(service)
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
