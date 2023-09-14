package handlers_test

import (
	"encoding/json"
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

func TestListImproveRequestsHandler(t *testing.T) {
	data := []struct {
		name string

		query string

		shouldCallService        bool
		shouldCallServiceWithIDs []uuid.UUID
		serviceResp              []*models.ImproveRequestPreview
		serviceErr               error

		expect       interface{}
		expectStatus int
	}{
		{
			name:                     "Success",
			query:                    "?ids=01010101-0101-0101-0101-010101010101,02020202-0202-0202-0202-020202020202",
			shouldCallService:        true,
			shouldCallServiceWithIDs: []uuid.UUID{goframework.NumberUUID(1), goframework.NumberUUID(2)},
			serviceResp: []*models.ImproveRequestPreview{
				{
					ID:                       goframework.NumberUUID(22),
					CreatedAt:                baseTime.Add(time.Hour),
					UserID:                   goframework.NumberUUID(201),
					Title:                    "title",
					Content:                  "content",
					UpVotes:                  128,
					DownVotes:                64,
					SuggestionsCount:         2,
					AcceptedSuggestionsCount: 1,
					RevisionCount:            3,
				},
				{
					ID:                       goframework.NumberUUID(2),
					CreatedAt:                baseTime,
					UserID:                   goframework.NumberUUID(200),
					Title:                    "title-2",
					Content:                  "content-2",
					UpVotes:                  256,
					DownVotes:                512,
					SuggestionsCount:         3,
					AcceptedSuggestionsCount: 2,
					RevisionCount:            2,
				},
			},
			expect: map[string]interface{}{
				"previews": []interface{}{
					map[string]interface{}{
						"id":                       goframework.NumberUUID(22).String(),
						"createdAt":                baseTime.Add(time.Hour).Format(time.RFC3339),
						"userID":                   goframework.NumberUUID(201).String(),
						"title":                    "title",
						"content":                  "content",
						"upVotes":                  float64(128),
						"downVotes":                float64(64),
						"suggestionsCount":         float64(2),
						"acceptedSuggestionsCount": float64(1),
						"revisionsCount":           float64(3),
					},
					map[string]interface{}{
						"id":                       goframework.NumberUUID(2).String(),
						"createdAt":                baseTime.Format(time.RFC3339),
						"userID":                   goframework.NumberUUID(200).String(),
						"title":                    "title-2",
						"content":                  "content-2",
						"upVotes":                  float64(256),
						"downVotes":                float64(512),
						"suggestionsCount":         float64(3),
						"acceptedSuggestionsCount": float64(2),
						"revisionsCount":           float64(2),
					},
				},
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewListImproveRequestsService(t)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/"+d.query, nil)

			if d.shouldCallService {
				service.
					On("List", c, d.shouldCallServiceWithIDs).
					Return(d.serviceResp, d.serviceErr)
			}

			handler := handlers.NewListImproveRequestsHandler(service)
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
