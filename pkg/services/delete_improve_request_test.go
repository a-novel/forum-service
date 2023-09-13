package services_test

import (
	"context"
	authmocks "github.com/a-novel/auth-service/framework/mocks"
	authmodels "github.com/a-novel/auth-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/errors"
	"github.com/a-novel/go-framework/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteImproveRequestService(t *testing.T) {
	data := []struct {
		name string

		token string
		id    uuid.UUID

		authClientResp *authmodels.UserTokenStatus
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
			id:    test.NumberUUID(1),
			authClientResp: &authmodels.UserTokenStatus{
				OK:    true,
				Token: &authmodels.UserToken{Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)}},
			},
			shouldCallGetRevision: true,
			getRevisionResp: &dao.ImproveRequestPreview{
				UserID: test.NumberUUID(100),
			},
			shouldCallDeleteRevision: true,
		},
		{
			name:  "Error/DeleteRevisionFailure",
			token: "tokenRaw",
			id:    test.NumberUUID(1),
			authClientResp: &authmodels.UserTokenStatus{
				OK:    true,
				Token: &authmodels.UserToken{Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)}},
			},
			shouldCallGetRevision: true,
			getRevisionResp: &dao.ImproveRequestPreview{
				UserID: test.NumberUUID(100),
			},
			shouldCallDeleteRevision: true,
			deleteRevisionErr:        fooErr,
			expectErr:                fooErr,
		},
		{
			name:  "Error/NotTheCreator",
			token: "tokenRaw",
			id:    test.NumberUUID(1),
			authClientResp: &authmodels.UserTokenStatus{
				OK:    true,
				Token: &authmodels.UserToken{Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(200)}},
			},
			shouldCallGetRevision: true,
			getRevisionResp: &dao.ImproveRequestPreview{
				UserID: test.NumberUUID(100),
			},
			expectErr: errors.ErrInvalidCredentials,
		},
		{
			name:  "Error/GetRevisionFailure",
			token: "tokenRaw",
			id:    test.NumberUUID(1),
			authClientResp: &authmodels.UserTokenStatus{
				OK:    true,
				Token: &authmodels.UserToken{Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(200)}},
			},
			shouldCallGetRevision: true,
			getRevisionErr:        fooErr,
			expectErr:             fooErr,
		},
		{
			name:           "Error/NotAuthenticated",
			token:          "tokenRaw",
			id:             test.NumberUUID(1),
			authClientResp: &authmodels.UserTokenStatus{},
			expectErr:      errors.ErrInvalidCredentials,
		},
		{
			name:          "Error/AuthClientFailure",
			token:         "tokenRaw",
			id:            test.NumberUUID(1),
			authClientErr: fooErr,
			expectErr:     fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveRequestRepository(t)
			authClient := authmocks.NewClient(t)

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
