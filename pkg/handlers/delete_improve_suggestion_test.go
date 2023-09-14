package handlers_test

import (
	"encoding/json"
	"github.com/a-novel/forum-service/pkg/handlers"
	"github.com/a-novel/forum-service/pkg/services"
	servicesmocks "github.com/a-novel/forum-service/pkg/services/mocks"
	goframework "github.com/a-novel/go-framework"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteImproveSuggestionHandler(t *testing.T) {
	data := []struct {
		name string

		authorization string
		query         string

		shouldCallService       bool
		shouldCallServiceWithID uuid.UUID
		serviceErr              error

		expect       interface{}
		expectStatus int
	}{
		{
			name:                    "Success",
			authorization:           "Bearer my-token",
			query:                   "?id=01010101-0101-0101-0101-010101010101",
			shouldCallService:       true,
			shouldCallServiceWithID: goframework.NumberUUID(1),
			expectStatus:            http.StatusNoContent,
		},
		{
			name:                    "Error/ErrInvalidCredentials",
			authorization:           "Bearer my-token",
			query:                   "?id=01010101-0101-0101-0101-010101010101",
			shouldCallService:       true,
			shouldCallServiceWithID: goframework.NumberUUID(1),
			serviceErr:              goframework.ErrInvalidCredentials,
			expectStatus:            http.StatusForbidden,
		},
		{
			name:                    "Error/ErrNotTheCreator",
			authorization:           "Bearer my-token",
			query:                   "?id=01010101-0101-0101-0101-010101010101",
			shouldCallService:       true,
			shouldCallServiceWithID: goframework.NumberUUID(1),
			serviceErr:              services.ErrNotTheCreator,
			expectStatus:            http.StatusUnauthorized,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewDeleteImproveSuggestionService(t)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/"+d.query, nil)
			c.Request.Header.Set("Authorization", d.authorization)

			if d.shouldCallService {
				service.
					On("Delete", c, d.authorization, d.shouldCallServiceWithID).
					Return(d.serviceErr)
			}

			handler := handlers.NewDeleteImproveSuggestionHandler(service)
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
