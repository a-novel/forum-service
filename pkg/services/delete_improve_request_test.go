package services_test

import (
	"context"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/services"
	apiclients "github.com/a-novel/go-apis/clients"
	apiclientsmocks "github.com/a-novel/go-apis/clients/mocks"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteImproveRequestService(t *testing.T) {
	data := []struct {
		name string

		token string
		id    uuid.UUID

		authClientResp *apiclients.UserTokenStatus
		authClientErr  error

		shouldCallGetRevision bool
		getRevisionResp       *dao.ImproveRequestPreview
		getRevisionErr        error

		shouldCallDeleteRevision bool
		deleteRevisionErr        error

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
			shouldCallGetRevision: true,
			getRevisionResp: &dao.ImproveRequestPreview{
				UserID: goframework.NumberUUID(100),
			},
			shouldCallDeleteRevision: true,
		},
		{
			name:  "Error/DeleteRevisionFailure",
			token: "tokenRaw",
			id:    goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK:    true,
				Token: &apiclients.UserToken{Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)}},
			},
			shouldCallGetRevision: true,
			getRevisionResp: &dao.ImproveRequestPreview{
				UserID: goframework.NumberUUID(100),
			},
			shouldCallDeleteRevision: true,
			deleteRevisionErr:        fooErr,
			expectErr:                fooErr,
		},
		{
			name:  "Error/NotTheCreator",
			token: "tokenRaw",
			id:    goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK:    true,
				Token: &apiclients.UserToken{Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(200)}},
			},
			shouldCallGetRevision: true,
			getRevisionResp: &dao.ImproveRequestPreview{
				UserID: goframework.NumberUUID(100),
			},
			expectErr: goframework.ErrInvalidCredentials,
		},
		{
			name:  "Error/GetRevisionFailure",
			token: "tokenRaw",
			id:    goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK:    true,
				Token: &apiclients.UserToken{Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(200)}},
			},
			shouldCallGetRevision: true,
			getRevisionErr:        fooErr,
			expectErr:             fooErr,
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
			repository := daomocks.NewImproveRequestRepository(t)
			authClient := apiclientsmocks.NewAuthClient(t)

			authClient.On("IntrospectToken", context.Background(), d.token).Return(d.authClientResp, d.authClientErr)

			if d.shouldCallGetRevision {
				repository.On("Get", context.Background(), d.id).Return(d.getRevisionResp, d.getRevisionErr)
			}

			if d.shouldCallDeleteRevision {
				repository.
					On("Delete", context.Background(), d.id).
					Return(d.deleteRevisionErr)
			}

			service := services.NewDeleteImproveRequestService(repository, authClient)
			err := service.Delete(context.Background(), d.token, d.id)

			require.ErrorIs(t, err, d.expectErr)

			repository.AssertExpectations(t)
			authClient.AssertExpectations(t)
		})
	}
}
