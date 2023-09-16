package services_test

import (
	"context"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/services"
	apiclients "github.com/a-novel/go-api-clients"
	apiclientsmocks "github.com/a-novel/go-api-clients/mocks"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteImproveSuggestionService(t *testing.T) {
	data := []struct {
		name string

		token string
		id    uuid.UUID

		authClientResp *apiclients.UserTokenStatus
		authClientErr  error

		shouldCallGet bool
		getResp       *dao.ImproveSuggestionModel
		getErr        error

		shouldCallDelete bool
		deleteErr        error

		expectErr error
	}{
		{
			name:  "Success",
			token: "tokenRaw",
			id:    goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK:    true,
				Token: &apiclients.UserToken{Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)}},
			},
			shouldCallGet: true,
			getResp: &dao.ImproveSuggestionModel{
				UserID: goframework.NumberUUID(100),
			},
			shouldCallDelete: true,
		},
		{
			name:  "Error/DeleteFailure",
			token: "tokenRaw",
			id:    goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK:    true,
				Token: &apiclients.UserToken{Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)}},
			},
			shouldCallGet: true,
			getResp: &dao.ImproveSuggestionModel{
				UserID: goframework.NumberUUID(100),
			},
			shouldCallDelete: true,
			deleteErr:        fooErr,
			expectErr:        fooErr,
		},
		{
			name:  "Error/NotTheCreator",
			token: "tokenRaw",
			id:    goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK:    true,
				Token: &apiclients.UserToken{Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(200)}},
			},
			shouldCallGet: true,
			getResp: &dao.ImproveSuggestionModel{
				UserID: goframework.NumberUUID(100),
			},
			expectErr: goframework.ErrInvalidCredentials,
		},
		{
			name:  "Error/GetFailure",
			token: "tokenRaw",
			id:    goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK:    true,
				Token: &apiclients.UserToken{Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(200)}},
			},
			shouldCallGet: true,
			getErr:        fooErr,
			expectErr:     fooErr,
		},
		{
			name:           "Error/NotAuthenticated",
			token:          "tokenRaw",
			id:             goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{},
			expectErr:      goframework.ErrInvalidCredentials,
		},
		{
			name:          "Error/AuthClientFailure",
			token:         "tokenRaw",
			id:            goframework.NumberUUID(1),
			authClientErr: fooErr,
			expectErr:     fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveSuggestionRepository(t)
			authClient := apiclientsmocks.NewAuthClient(t)

			authClient.On("IntrospectToken", context.Background(), d.token).Return(d.authClientResp, d.authClientErr)

			if d.shouldCallGet {
				repository.On("Get", context.Background(), d.id).Return(d.getResp, d.getErr)
			}

			if d.shouldCallDelete {
				repository.
					On("Delete", context.Background(), d.id).
					Return(d.deleteErr)
			}

			service := services.NewDeleteImproveSuggestionService(repository, authClient)
			err := service.Delete(context.Background(), d.token, d.id)

			require.ErrorIs(t, err, d.expectErr)

			repository.AssertExpectations(t)
			authClient.AssertExpectations(t)
		})
	}
}
