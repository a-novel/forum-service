package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/a-novel/forum-service/pkg/handlers"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	servicesmocks "github.com/a-novel/forum-service/pkg/services/mocks"
	goframework "github.com/a-novel/go-framework"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateImproveRequestHandler(t *testing.T) {
	data := []struct {
		name string

		authorization string

		body interface{}

		shouldCallService            bool
		shouldCallServiceWithTitle   string
		shouldCallServiceWithContent string
		shouldCallServiceWithSource  uuid.UUID
		serviceResp                  *models.ImproveRequestPreview
		serviceErr                   error

		expect       interface{}
		expectStatus int
	}{
		{
			name:          "Success",
			authorization: "Bearer my-token",
			body: map[string]interface{}{
				"title":    "title",
				"content":  "content",
				"sourceID": goframework.NumberUUID(10).String(),
			},
			shouldCallService:            true,
			shouldCallServiceWithTitle:   "title",
			shouldCallServiceWithContent: "content",
			shouldCallServiceWithSource:  goframework.NumberUUID(10),
			serviceResp: &models.ImproveRequestPreview{
				ID:        goframework.NumberUUID(10),
				CreatedAt: baseTime,
				UserID:    goframework.NumberUUID(100),
				Title:     "title",
				Content:   "content",
			},
			expect: map[string]interface{}{
				"id":                       goframework.NumberUUID(10).String(),
				"createdAt":                baseTime.Format(time.RFC3339),
				"userID":                   goframework.NumberUUID(100).String(),
				"title":                    "title",
				"content":                  "content",
				"upVotes":                  float64(0),
				"downVotes":                float64(0),
				"revisionsCount":           float64(0),
				"suggestionsCount":         float64(0),
				"acceptedSuggestionsCount": float64(0),
			},
			expectStatus: http.StatusCreated,
		},
		{
			name:          "Error/ErrNotTheCreator",
			authorization: "Bearer my-token",
			body: map[string]interface{}{
				"title":    "title",
				"content":  "content",
				"sourceID": goframework.NumberUUID(10).String(),
			},
			shouldCallService:            true,
			shouldCallServiceWithTitle:   "title",
			shouldCallServiceWithContent: "content",
			shouldCallServiceWithSource:  goframework.NumberUUID(10),
			serviceErr:                   services.ErrNotTheCreator,
			expectStatus:                 http.StatusUnauthorized,
		},
		{
			name:          "Error/ErrInvalidCredentials",
			authorization: "Bearer my-token",
			body: map[string]interface{}{
				"title":    "title",
				"content":  "content",
				"sourceID": goframework.NumberUUID(10).String(),
			},
			shouldCallService:            true,
			shouldCallServiceWithTitle:   "title",
			shouldCallServiceWithContent: "content",
			shouldCallServiceWithSource:  goframework.NumberUUID(10),
			serviceErr:                   goframework.ErrInvalidCredentials,
			expectStatus:                 http.StatusForbidden,
		},
		{
			name:          "Error/ErrInvalidEntity",
			authorization: "Bearer my-token",
			body: map[string]interface{}{
				"title":    "title",
				"content":  "content",
				"sourceID": goframework.NumberUUID(10).String(),
			},
			shouldCallService:            true,
			shouldCallServiceWithTitle:   "title",
			shouldCallServiceWithContent: "content",
			shouldCallServiceWithSource:  goframework.NumberUUID(10),
			serviceErr:                   goframework.ErrInvalidEntity,
			expectStatus:                 http.StatusUnprocessableEntity,
		},
		{
			name:          "Error/BadRequest",
			authorization: "Bearer my-token",
			body: map[string]interface{}{
				"title":    "title",
				"content":  "content",
				"sourceID": "fake uuid",
			},
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewCreateImproveRequestService(t)

			mrshBody, err := json.Marshal(d.body)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/", bytes.NewReader(mrshBody))
			c.Request.Header.Set("Authorization", d.authorization)

			if d.shouldCallService {
				service.
					On(
						"Create", c,
						d.authorization,
						d.shouldCallServiceWithTitle,
						d.shouldCallServiceWithContent,
						d.shouldCallServiceWithSource,
						mock.Anything, mock.Anything,
					).
					Return(d.serviceResp, d.serviceErr)
			}

			handler := handlers.NewCreateImproveRequestHandler(service)
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
