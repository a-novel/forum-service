package handlers_test

import (
	"encoding/json"
	"github.com/a-novel/forum-service/pkg/handlers"
	"github.com/a-novel/forum-service/pkg/models"
	servicesmocks "github.com/a-novel/forum-service/pkg/services/mocks"
	"github.com/a-novel/go-framework/test"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestListImproveSuggestionHandler(t *testing.T) {
	data := []struct {
		name string

		query string

		shouldCallService        bool
		shouldCallServiceWithIDs []uuid.UUID
		serviceResp              []*models.ImproveSuggestion
		serviceErr               error

		expect       interface{}
		expectStatus int
	}{
		{
			name:                     "Success",
			query:                    "?ids=01010101-0101-0101-0101-010101010101,02020202-0202-0202-0202-020202020202",
			shouldCallService:        true,
			shouldCallServiceWithIDs: []uuid.UUID{test.NumberUUID(1), test.NumberUUID(2)},
			serviceResp: []*models.ImproveSuggestion{
				{
					ID:        test.NumberUUID(1),
					CreatedAt: baseTime,
					UpdatedAt: lo.ToPtr(baseTime.Add(3 * time.Hour)),
					SourceID:  test.NumberUUID(10),
					UserID:    test.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					RequestID: test.NumberUUID(1),
					Title:     "title",
					Content:   "content",
				},
				{
					ID:        test.NumberUUID(2),
					CreatedAt: baseTime,
					UpdatedAt: lo.ToPtr(baseTime.Add(2 * time.Hour)),
					SourceID:  test.NumberUUID(20),
					UserID:    test.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					RequestID: test.NumberUUID(1),
					Title:     "title",
					Content:   "content",
				},
			},
			expect: map[string]interface{}{
				"previews": []interface{}{
					map[string]interface{}{
						"id":        test.NumberUUID(1).String(),
						"createdAt": baseTime.Format(time.RFC3339),
						"updatedAt": baseTime.Add(3 * time.Hour).Format(time.RFC3339),
						"sourceID":  test.NumberUUID(10).String(),
						"userID":    test.NumberUUID(200).String(),
						"requestID": test.NumberUUID(1).String(),
						"validated": true,
						"title":     "title",
						"content":   "content",
						"upVotes":   float64(16),
						"downVotes": float64(8),
					},
					map[string]interface{}{
						"id":        test.NumberUUID(2).String(),
						"createdAt": baseTime.Format(time.RFC3339),
						"updatedAt": baseTime.Add(2 * time.Hour).Format(time.RFC3339),
						"sourceID":  test.NumberUUID(20).String(),
						"userID":    test.NumberUUID(100).String(),
						"requestID": test.NumberUUID(1).String(),
						"validated": false,
						"title":     "title",
						"content":   "content",
						"upVotes":   float64(32),
						"downVotes": float64(16),
					},
				},
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewListImproveSuggestionsService(t)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/"+d.query, nil)

			if d.shouldCallService {
				service.
					On("List", c, d.shouldCallServiceWithIDs).
					Return(d.serviceResp, d.serviceErr)
			}

			handler := handlers.NewListImproveSuggestionsHandler(service)
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
