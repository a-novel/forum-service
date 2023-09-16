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

func TestValidateImproveSuggestionService(t *testing.T) {
	data := []struct {
		name string

		tokenRaw  string
		validated bool
		id        uuid.UUID

		authClientResp *apiclients.UserTokenStatus
		authClientErr  error

		shouldCallGetSuggestion bool
		getSuggestionResp       *dao.ImproveSuggestionModel
		getSuggestionErr        error

		shouldCallGetRequest bool
		getRequestResp       *dao.ImproveRequestRevisionModel
		getRequestErr        error

		shouldCallValidateSuggestion bool
		validateSuggestionErr        error

		expectErr error
	}{
		{
			name:      "Success",
			tokenRaw:  "token",
			validated: true,
			id:        goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallGetSuggestion: true,
			getSuggestionResp: &dao.ImproveSuggestionModel{
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{RequestID: goframework.NumberUUID(10)},
			},
			shouldCallGetRequest: true,
			getRequestResp: &dao.ImproveRequestRevisionModel{
				UserID: goframework.NumberUUID(100),
			},
			shouldCallValidateSuggestion: true,
		},
		{
			name:      "Error/ValidateFailure",
			tokenRaw:  "token",
			validated: true,
			id:        goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallGetSuggestion: true,
			getSuggestionResp: &dao.ImproveSuggestionModel{
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{RequestID: goframework.NumberUUID(10)},
			},
			shouldCallGetRequest: true,
			getRequestResp: &dao.ImproveRequestRevisionModel{
				UserID: goframework.NumberUUID(100),
			},
			shouldCallValidateSuggestion: true,
			validateSuggestionErr:        fooErr,
			expectErr:                    fooErr,
		},
		{
			name:      "Error/NotTheRequestOwner",
			tokenRaw:  "token",
			validated: true,
			id:        goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallGetSuggestion: true,
			getSuggestionResp: &dao.ImproveSuggestionModel{
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{RequestID: goframework.NumberUUID(10)},
			},
			shouldCallGetRequest: true,
			getRequestResp: &dao.ImproveRequestRevisionModel{
				UserID: goframework.NumberUUID(200),
			},
			expectErr: goframework.ErrInvalidCredentials,
		},
		{
			name:      "Error/GetRequestFailure",
			tokenRaw:  "token",
			validated: true,
			id:        goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallGetSuggestion: true,
			getSuggestionResp: &dao.ImproveSuggestionModel{
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{RequestID: goframework.NumberUUID(10)},
			},
			shouldCallGetRequest: true,
			getRequestErr:        fooErr,
			expectErr:            fooErr,
		},
		{
			name:      "Error/GetSuggestionFailure",
			tokenRaw:  "token",
			validated: true,
			id:        goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallGetSuggestion: true,
			getSuggestionErr:        fooErr,
			expectErr:               fooErr,
		},
		{
			name:           "Error/NotAuthenticated",
			tokenRaw:       "token",
			validated:      true,
			id:             goframework.NumberUUID(1),
			authClientResp: &apiclients.UserTokenStatus{},
			expectErr:      goframework.ErrInvalidCredentials,
		},
		{
			name:          "Error/IntrospectTokenFailure",
			tokenRaw:      "token",
			validated:     true,
			id:            goframework.NumberUUID(1),
			authClientErr: fooErr,
			expectErr:     fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveSuggestionRepository(t)
			requestRepository := daomocks.NewImproveRequestRepository(t)
			authClient := apiclientsmocks.NewAuthClient(t)

			authClient.On("IntrospectToken", context.Background(), d.tokenRaw).Return(d.authClientResp, d.authClientErr)

			if d.shouldCallGetSuggestion {
				repository.On("Get", context.Background(), d.id).Return(d.getSuggestionResp, d.getSuggestionErr)
			}

			if d.shouldCallGetRequest {
				requestRepository.
					On("GetRevision", context.Background(), d.getSuggestionResp.RequestID).
					Return(d.getRequestResp, d.getRequestErr)
			}

			if d.shouldCallValidateSuggestion {
				repository.
					On("Validate", context.Background(), d.validated, d.id).
					Return(nil, d.validateSuggestionErr)
			}

			service := services.NewValidateImproveSuggestionService(repository, requestRepository, authClient)
			err := service.Validate(context.Background(), d.tokenRaw, d.validated, d.id)

			require.ErrorIs(t, err, d.expectErr)

			repository.AssertExpectations(t)
			requestRepository.AssertExpectations(t)
			authClient.AssertExpectations(t)
		})
	}
}
