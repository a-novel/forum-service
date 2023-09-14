package services_test

import (
	"context"
	authmocks "github.com/a-novel/auth-service/framework/mocks"
	authmodels "github.com/a-novel/auth-service/pkg/models"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

func TestCreateImproveRequestService(t *testing.T) {
	data := []struct {
		name string

		tokenRaw string
		title    string
		content  string
		sourceID uuid.UUID
		id       uuid.UUID
		now      time.Time

		authClientResp *authmodels.UserTokenStatus
		authClientErr  error

		shouldCallGet bool
		getResp       *dao.ImproveRequestPreview
		getErr        error

		shouldCallCreateRevision bool
		createRevisionResp       *dao.ImproveRequestPreview
		createRevisionErr        error

		expect    *models.ImproveRequestPreview
		expectErr error
	}{
		{
			name:     "Success",
			tokenRaw: "token",
			title:    "title",
			content:  "content",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallGet:            true,
			shouldCallCreateRevision: true,
			createRevisionResp: &dao.ImproveRequestPreview{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, nil),
				Title:    "title",
				Content:  "content",
				UserID:   goframework.NumberUUID(100),
			},
			expect: &models.ImproveRequestPreview{
				ID:        goframework.NumberUUID(10),
				CreatedAt: baseTime,
				Title:     "title",
				Content:   "content",
				UserID:    goframework.NumberUUID(100),
			},
		},
		{
			name:     "Success/NewRevision",
			tokenRaw: "token",
			title:    "title",
			content:  "content",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallGet: true,
			getResp: &dao.ImproveRequestPreview{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, nil),
				Title:    "old title",
				Content:  "old content",
				UserID:   goframework.NumberUUID(100),
			},
			shouldCallCreateRevision: true,
			createRevisionResp: &dao.ImproveRequestPreview{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, nil),
				Title:    "title",
				Content:  "content",
				UserID:   goframework.NumberUUID(100),
			},
			expect: &models.ImproveRequestPreview{
				ID:        goframework.NumberUUID(10),
				CreatedAt: baseTime,
				Title:     "title",
				Content:   "content",
				UserID:    goframework.NumberUUID(100),
			},
		},
		{
			name:     "Error/CreateRevisionFailure",
			tokenRaw: "token",
			title:    "title",
			content:  "content",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallGet: true,
			getResp: &dao.ImproveRequestPreview{
				UserID: goframework.NumberUUID(100),
			},
			shouldCallCreateRevision: true,
			createRevisionErr:        fooErr,
			expectErr:                fooErr,
		},
		{
			name:     "Error/NotTheCreator",
			tokenRaw: "token",
			title:    "title",
			content:  "content",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallGet: true,
			getResp: &dao.ImproveRequestPreview{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, nil),
				Title:    "old title",
				Content:  "old content",
				UserID:   goframework.NumberUUID(200),
			},
			expectErr: goframework.ErrInvalidCredentials,
		},
		{
			name:     "Error/PreviousRevisionsFailure",
			tokenRaw: "token",
			title:    "title",
			content:  "content",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			shouldCallGet: true,
			getErr:        fooErr,
			expectErr:     fooErr,
		},
		{
			name:     "Error/BadTitle",
			tokenRaw: "token",
			title:    "title\nwith a line break",
			content:  "content",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			expectErr: goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/TitleTooShort",
			tokenRaw: "token",
			title:    "t",
			content:  "content",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			expectErr: goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/NoTitle",
			tokenRaw: "token",
			content:  "content",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			expectErr: goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/TitleTooLong",
			tokenRaw: "token",
			title:    strings.Repeat("a", services.MaxTitleLength+1),
			content:  "content",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			expectErr: goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/ContentTooShort",
			tokenRaw: "token",
			title:    "title",
			content:  "c",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			expectErr: goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/NoContent",
			tokenRaw: "token",
			title:    "title",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			expectErr: goframework.ErrInvalidEntity,
		},
		{
			name:     "Error/ContentTooLong",
			tokenRaw: "token",
			title:    "title",
			content:  strings.Repeat("a", services.MaxContentLength+1),
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(1),
			now:      baseTime,
			authClientResp: &authmodels.UserTokenStatus{
				OK: true,
				Token: &authmodels.UserToken{
					Payload: authmodels.UserTokenPayload{ID: goframework.NumberUUID(100)},
				},
			},
			expectErr: goframework.ErrInvalidEntity,
		},
		{
			name:           "Error/NotAuthenticated",
			tokenRaw:       "token",
			title:          "title",
			content:        "content",
			sourceID:       goframework.NumberUUID(10),
			id:             goframework.NumberUUID(1),
			now:            baseTime,
			authClientResp: &authmodels.UserTokenStatus{},
			expectErr:      goframework.ErrInvalidCredentials,
		},
		{
			name:          "Error/AuthClientFailure",
			tokenRaw:      "token",
			title:         "title",
			content:       "content",
			sourceID:      goframework.NumberUUID(10),
			id:            goframework.NumberUUID(1),
			now:           baseTime,
			authClientErr: fooErr,
			expectErr:     fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveRequestRepository(t)
			authClient := authmocks.NewClient(t)

			authClient.On("IntrospectToken", context.Background(), d.tokenRaw).Return(d.authClientResp, d.authClientErr)

			if d.shouldCallGet {
				repository.On("Get", context.Background(), d.sourceID).Return(d.getResp, d.getErr)
			}

			if d.shouldCallCreateRevision {
				repository.
					On("Create", context.Background(), d.authClientResp.Token.Payload.ID, d.title, d.content, d.sourceID, d.id, d.now).
					Return(d.createRevisionResp, d.createRevisionErr)
			}

			service := services.NewCreateImproveRequestService(repository, authClient)
			res, err := service.Create(context.Background(), d.tokenRaw, d.title, d.content, d.sourceID, d.id, d.now)

			require.ErrorIs(t, err, d.expectErr)
			require.Equal(t, d.expect, res)

			repository.AssertExpectations(t)
			authClient.AssertExpectations(t)
		})
	}
}
