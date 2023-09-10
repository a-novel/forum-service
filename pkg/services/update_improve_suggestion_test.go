package services_test

import (
	"context"
	authmocks "github.com/a-novel/auth-service/framework/mocks"
	authmodels "github.com/a-novel/auth-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/errors"
	"github.com/a-novel/go-framework/postgresql"
	"github.com/a-novel/go-framework/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

func TestUpdateImproveSuggestionService(t *testing.T) {
	data := []struct {
		name string

		suggestion *models.ImproveSuggestionForm
		tokenRaw   string
		id         uuid.UUID
		now        time.Time

		authClientResp *authmodels.UserTokenStatus
		authClientErr  error

		shouldCallGetRevision bool
		getRevisionResp       *dao.ImproveRequestRevisionModel
		getRevisionErr        error

		shouldCallGetSuggestion bool
		getSuggestionResp       *dao.ImproveSuggestionModel
		getSuggestionErr        error

		shouldCallUpdateSuggestion bool
		updateSuggestionResp       *dao.ImproveSuggestionModel
		updateSuggestionErr        error

		expect    *models.ImproveSuggestion
		expectErr error
	}{
		{
			name: "Success",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			tokenRaw: "token",
			id:       test.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			shouldCallGetRevision: true,
			getRevisionResp: &dao.ImproveRequestRevisionModel{
				SourceID: test.NumberUUID(10),
			},
			shouldCallGetSuggestion: true,
			getSuggestionResp: &dao.ImproveSuggestionModel{
				Metadata: postgresql.NewMetadata(test.NumberUUID(1), baseTime, nil),
				SourceID: test.NumberUUID(10),
				UserID:   test.NumberUUID(100),
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: test.NumberUUID(2),
					Title:     "old title",
					Content:   "old content",
				},
			},
			shouldCallUpdateSuggestion: true,
			updateSuggestionResp: &dao.ImproveSuggestionModel{
				Metadata: postgresql.NewMetadata(test.NumberUUID(1), baseTime, nil),
				SourceID: test.NumberUUID(10),
				UserID:   test.NumberUUID(100),
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: test.NumberUUID(1),
					Title:     "title",
					Content:   "content",
				},
			},
			expect: &models.ImproveSuggestion{
				ID:        test.NumberUUID(1),
				CreatedAt: baseTime,
				SourceID:  test.NumberUUID(10),
				UserID:    test.NumberUUID(100),
				RequestID: test.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
		{
			name: "Error/UpdateSuggestionFailure",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			tokenRaw: "token",
			id:       test.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			shouldCallGetRevision: true,
			getRevisionResp: &dao.ImproveRequestRevisionModel{
				SourceID: test.NumberUUID(10),
			},
			shouldCallGetSuggestion: true,
			getSuggestionResp: &dao.ImproveSuggestionModel{
				Metadata: postgresql.NewMetadata(test.NumberUUID(1), baseTime, nil),
				SourceID: test.NumberUUID(10),
				UserID:   test.NumberUUID(100),
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: test.NumberUUID(2),
					Title:     "old title",
					Content:   "old content",
				},
			},
			shouldCallUpdateSuggestion: true,
			updateSuggestionErr:        fooErr,
			expectErr:                  fooErr,
		},
		{
			name: "Error/RequestAndSourceMismatch",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			tokenRaw: "token",
			id:       test.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			shouldCallGetRevision: true,
			getRevisionResp: &dao.ImproveRequestRevisionModel{
				SourceID: test.NumberUUID(20),
			},
			shouldCallGetSuggestion: true,
			getSuggestionResp: &dao.ImproveSuggestionModel{
				Metadata: postgresql.NewMetadata(test.NumberUUID(1), baseTime, nil),
				SourceID: test.NumberUUID(10),
				UserID:   test.NumberUUID(100),
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: test.NumberUUID(2),
					Title:     "old title",
					Content:   "old content",
				},
			},
			expectErr: errors.ErrInvalidEntity,
		},
		{
			name: "Error/GetSuggestionError",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			tokenRaw: "token",
			id:       test.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			shouldCallGetRevision: true,
			getRevisionResp: &dao.ImproveRequestRevisionModel{
				SourceID: test.NumberUUID(10),
			},
			shouldCallGetSuggestion: true,
			getSuggestionErr:        fooErr,
			expectErr:               fooErr,
		},
		{
			name: "Error/GetRevisionError",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			tokenRaw: "token",
			id:       test.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			shouldCallGetRevision: true,
			getRevisionErr:        fooErr,
			expectErr:             fooErr,
		},
		{
			name:     "Error/BadTitle",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "title\nwith a line break",
				Content:   "content",
			},
			id:  test.NumberUUID(1),
			now: baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			expectErr: errors.ErrInvalidEntity,
		},
		{
			name:     "Error/TitleTooShort",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "t",
				Content:   "content",
			},
			id:  test.NumberUUID(1),
			now: baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			expectErr: errors.ErrInvalidEntity,
		},
		{
			name:     "Error/NoTitle",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Content:   "content",
			},
			id:  test.NumberUUID(1),
			now: baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			expectErr: errors.ErrInvalidEntity,
		},
		{
			name:     "Error/TitleTooLong",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     strings.Repeat("a", services.MaxTitleLength+1),
				Content:   "content",
			},
			id:  test.NumberUUID(1),
			now: baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			expectErr: errors.ErrInvalidEntity,
		},
		{
			name:     "Error/ContentTooShort",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "title",
				Content:   "c",
			},
			id:  test.NumberUUID(1),
			now: baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			expectErr: errors.ErrInvalidEntity,
		},
		{
			name:     "Error/NoContent",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "title",
			},
			id:  test.NumberUUID(1),
			now: baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			expectErr: errors.ErrInvalidEntity,
		},
		{
			name:     "Error/ContentTooLong",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "title",
				Content:   strings.Repeat("a", services.MaxContentLength+1),
			},
			id:  test.NumberUUID(1),
			now: baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: test.NumberUUID(100)},
				},
			},
			expectErr: errors.ErrInvalidEntity,
		},
		{
			name:     "Error/NotAuthenticated",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			id:             test.NumberUUID(1),
			now:            baseTime,
			authClientResp: &authmodels.UserTokenStatus{},
			expectErr:      errors.ErrInvalidCredentials,
		},
		{
			name:     "Error/AuthClientFailure",
			tokenRaw: "token",
			suggestion: &models.ImproveSuggestionForm{
				RequestID: test.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
			id:            test.NumberUUID(1),
			now:           baseTime,
			authClientErr: fooErr,
			expectErr:     fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveSuggestionRepository(t)
			requestsRepository := daomocks.NewImproveRequestRepository(t)
			authClient := authmocks.NewClient(t)

			authClient.On("IntrospectToken", context.Background(), d.tokenRaw).Return(d.authClientResp, d.authClientErr)

			if d.shouldCallGetRevision {
				requestsRepository.
					On("GetRevision", context.Background(), d.suggestion.RequestID).
					Return(d.getRevisionResp, d.getRevisionErr)
			}

			if d.shouldCallGetSuggestion {
				repository.
					On("Get", context.Background(), d.suggestion.RequestID).
					Return(d.getSuggestionResp, d.getSuggestionErr)
			}

			if d.shouldCallUpdateSuggestion {
				repository.
					On("Update", context.Background(), mock.Anything, d.id, d.now).
					Return(d.updateSuggestionResp, d.updateSuggestionErr)
			}

			service := services.NewUpdateImproveSuggestionService(repository, requestsRepository, authClient)
			resp, err := service.Update(context.Background(), d.tokenRaw, d.suggestion, d.id, d.now)

			require.ErrorIs(t, err, d.expectErr)
			require.Equal(t, d.expect, resp)

			repository.AssertExpectations(t)
			requestsRepository.AssertExpectations(t)
			authClient.AssertExpectations(t)
		})
	}
}
