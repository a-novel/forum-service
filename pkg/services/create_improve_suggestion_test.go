package services_test

import (
	"context"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	apiclients "github.com/a-novel/go-apis/clients"
	apiclientsmocks "github.com/a-novel/go-apis/clients/mocks"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

func TestCreateImproveSuggestionService(t *testing.T) {
	data := []struct {
		name string

		suggestion *models.ImproveSuggestionForm
		tokenRaw   string
		id         uuid.UUID
		now        time.Time

		authClientResp *apiclients.UserTokenStatus
		authClientErr  error

		shouldCallPermissionsClient bool
		permissionsClientErr        error

		shouldCallGetRevision bool
		getRevisionResp       *dao.ImproveRequestRevisionModel
		getRevisionErr        error

		shouldCallCreateSuggestion bool
		createSuggestionResp       *dao.ImproveSuggestionModel
		createSuggestionErr        error

		expect    *models.ImproveSuggestion
		expectErr error
	}{
		{
			name: "Success",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			tokenRaw: "token",
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallPermissionsClient: true,
			shouldCallGetRevision:       true,
			getRevisionResp: &dao.ImproveRequestRevisionModel{
				SourceID: goframework.NumberUUID(10),
			},
			shouldCallCreateSuggestion: true,
			createSuggestionResp: &dao.ImproveSuggestionModel{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, nil),
				SourceID: goframework.NumberUUID(10),
				UserID:   goframework.NumberUUID(100),
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: goframework.NumberUUID(1),
					Title:     "title",
					Content:   "content",
				},
			},
			expect: &models.ImproveSuggestion{
				ID:        goframework.NumberUUID(1),
				CreatedAt: baseTime,
				SourceID:  goframework.NumberUUID(10),
				UserID:    goframework.NumberUUID(100),
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
		{
			name: "Error/CreateSuggestionFailure",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			tokenRaw: "token",
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallPermissionsClient: true,
			shouldCallGetRevision:       true,
			getRevisionResp: &dao.ImproveRequestRevisionModel{
				SourceID: goframework.NumberUUID(10),
			},
			shouldCallCreateSuggestion: true,
			createSuggestionErr:        fooErr,
			expectErr:                  fooErr,
		},
		{
			name: "Error/GetRevisionError",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			tokenRaw: "token",
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallPermissionsClient: true,
			shouldCallGetRevision:       true,
			getRevisionErr:              fooErr,
			expectErr:                   fooErr,
		},
		{
			name:     "Error/BadTitle",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     "title\nwith a line break",
				Content:   "content",
			},
			id:  goframework.NumberUUID(1),
			now: baseTime,
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallPermissionsClient: true,
			expectErr:                   goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/TitleTooShort",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     "t",
				Content:   "content",
			},
			id:  goframework.NumberUUID(1),
			now: baseTime,
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallPermissionsClient: true,
			expectErr:                   goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/NoTitle",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Content:   "content",
			},
			id:  goframework.NumberUUID(1),
			now: baseTime,
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallPermissionsClient: true,
			expectErr:                   goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/TitleTooLong",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     strings.Repeat("a", services.MaxTitleLength+1),
				Content:   "content",
			},
			id:  goframework.NumberUUID(1),
			now: baseTime,
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallPermissionsClient: true,
			expectErr:                   goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/ContentTooShort",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "c",
			},
			id:  goframework.NumberUUID(1),
			now: baseTime,
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallPermissionsClient: true,
			expectErr:                   goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/NoContent",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
			},
			id:  goframework.NumberUUID(1),
			now: baseTime,
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallPermissionsClient: true,
			expectErr:                   goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/ContentTooLong",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   strings.Repeat("a", services.MaxContentLength+1),
			},
			id:  goframework.NumberUUID(1),
			now: baseTime,
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallPermissionsClient: true,
			expectErr:                   goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/NotAuthenticated",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			id:             goframework.NumberUUID(1),
			now:            baseTime,
			authClientResp: &apiclients.UserTokenStatus{},
			expectErr:      goframework.ErrInvalidCredentials,
		},
		{
			name:     "Error/AuthClientFailure",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			id:            goframework.NumberUUID(1),
			now:           baseTime,
			authClientErr: fooErr,
			expectErr:     fooErr,
		},
		{
			name:     "Error/GetPermissions",
			tokenRaw: "token",
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			suggestion: &models.ImproveSuggestionForm{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			authClientResp: &apiclients.UserTokenStatus{
				OK: true,
				Token: &apiclients.UserToken{
					Payload: apiclients.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallPermissionsClient: true,
			permissionsClientErr:        fooErr,
			expectErr:                   fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveSuggestionRepository(t)
			requestsRepository := daomocks.NewImproveRequestRepository(t)
			authClient := apiclientsmocks.NewAuthClient(t)
			permissionsClient := apiclientsmocks.NewPermissionsClient(t)

			authClient.On("IntrospectToken", context.Background(), d.tokenRaw).Return(d.authClientResp, d.authClientErr)

			if d.shouldCallPermissionsClient {
				permissionsClient.
					On("HasUserScope", context.Background(), apiclients.HasUserScopeQuery{
						UserID: d.authClientResp.Token.Payload.ID,
						Scope:  apiclients.CanPostImproveSuggestion,
					}).
					Return(d.permissionsClientErr)
			}

			if d.shouldCallGetRevision {
				requestsRepository.
					On("GetRevision", context.Background(), d.suggestion.RequestID).
					Return(d.getRevisionResp, d.getRevisionErr)
			}

			if d.shouldCallCreateSuggestion {
				repository.
					On("Create", context.Background(), mock.Anything, d.authClientResp.Token.Payload.ID, d.getRevisionResp.SourceID, d.id, d.now).
					Return(d.createSuggestionResp, d.createSuggestionErr)
			}

			service := services.NewCreateImproveSuggestionService(repository, requestsRepository, authClient, permissionsClient)
			resp, err := service.Create(context.Background(), d.tokenRaw, d.suggestion, d.id, d.now)

			require.ErrorIs(t, err, d.expectErr)
			require.Equal(t, d.expect, resp)

			repository.AssertExpectations(t)
			requestsRepository.AssertExpectations(t)
			authClient.AssertExpectations(t)
			permissionsClient.AssertExpectations(t)
		})
	}
}
