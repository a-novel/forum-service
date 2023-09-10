package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/a-novel/forum-service/pkg/handlers"
	"github.com/a-novel/forum-service/pkg/services"
	servicesmocks "github.com/a-novel/forum-service/pkg/services/mocks"
	"github.com/a-novel/go-framework/errors"
	"github.com/a-novel/go-framework/test"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateImproveSuggestionHandler(t *testing.T) {
	data := []struct {
		name string

		authorization string
		body          interface{}

		shouldCallService              bool
		shouldCallServiceWithValidated bool
		shouldCallServiceWithID        uuid.UUID
		serviceErr                     error

		expect       interface{}
		expectStatus int
	}{
		{
			name: "Success",
			body: map[string]interface{}{
				"validated": true,
				"id":        test.NumberUUID(1).String(),
			},
			shouldCallService:              true,
			shouldCallServiceWithValidated: true,
			shouldCallServiceWithID:        test.NumberUUID(1),
			expectStatus:                   http.StatusNoContent,
		},
		{
			name: "Error/ErrNotTheCreator",
			body: map[string]interface{}{
				"validated": true,
				"id":        test.NumberUUID(1).String(),
			},
			shouldCallService:              true,
			shouldCallServiceWithValidated: true,
			shouldCallServiceWithID:        test.NumberUUID(1),
			serviceErr:                     services.ErrNotTheCreator,
			expectStatus:                   http.StatusUnauthorized,
		},
		{
			name: "Error/ErrInvalidCredentials",
			body: map[string]interface{}{
				"validated": true,
				"id":        test.NumberUUID(1).String(),
			},
			shouldCallService:              true,
			shouldCallServiceWithValidated: true,
			shouldCallServiceWithID:        test.NumberUUID(1),
			serviceErr:                     errors.ErrInvalidCredentials,
			expectStatus:                   http.StatusForbidden,
		},
		{
			name: "Error/BadForm",
			body: map[string]interface{}{
				"validated": true,
				"id":        "fake uuid",
			},
			expectStatus: http.StatusBadRequest,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			service := servicesmocks.NewValidateImproveSuggestionService(t)

			mrshBody, err := json.Marshal(d.body)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", "/", bytes.NewReader(mrshBody))
			c.Request.Header.Set("Authorization", d.authorization)

			if d.shouldCallService {
				service.
					On("Validate", c, d.authorization, d.shouldCallServiceWithValidated, d.shouldCallServiceWithID).
					Return(d.serviceErr)
			}

			handler := handlers.NewValidateImproveSuggestionHandler(service)
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
