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

func TestGetImproveRequestRevisionHandler(t *testing.T) {
	data := []struct {
		name string

		query string

		shouldCallService       bool
		shouldCallServiceWithID uuid.UUID
		serviceResp             *models.ImproveRequestRevision
		serviceErr              error

		expect       interface{}
		expectStatus int
	}{
		{
			name:                    "Success",
			query:                   "?id=01010101-0101-0101-0101-010101010101",
			shouldCallService:       true,
			shouldCallServiceWithID: goframework.NumberUUID(1),
			serviceResp: &models.ImproveRequestRevision{
				ID:        goframework.NumberUUID(1),
				SourceID:  goframework.NumberUUID(10),
				CreatedAt: baseTime,
				UserID:    goframework.NumberUUID(100),
				Title:     "title",
				Content:   "content",
			},
			expect: map[string]interface{}{
				"id":        goframework.NumberUUID(1).String(),
				"createdAt": baseTime.Format(time.RFC3339),
				"sourceID":  goframework.NumberUUID(10).String(),
				"userID":    goframework.NumberUUID(100).String(),
				"title":     "title",
				"content":   "content",
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
			service := servicesmocks.NewGetImproveRequestRevisionService(t)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/"+d.query, nil)

			if d.shouldCallService {
				service.
					On("Get", c, d.shouldCallServiceWithID).
					Return(d.serviceResp, d.serviceErr)
			}

			handler := handlers.NewGetImproveRequestRevisionHandler(service)
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
