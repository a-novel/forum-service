package services_test

import (
	"context"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-apis"
	goframework "github.com/a-novel/go-framework"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchImproveRequestsService(t *testing.T) {
	data := []struct {
		name string

		query models.SearchImproveRequestsQuery

		shouldCallDAO         bool
		shouldCallDAOWithForm dao.ImproveRequestSearchQuery
		queryResults          []*dao.ImproveRequestPreview
		queryTotal            int
		queryErr              error

		expectedResults []*models.ImproveRequestPreview
		expectedTotal   int
		expectedErr     error
	}{
		{
			name: "Success",
			query: models.SearchImproveRequestsQuery{
				Limit: 10,
			},
			shouldCallDAO: true,
			queryResults: []*dao.ImproveRequestPreview{
				{
					Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, nil),
					UserID:                   goframework.NumberUUID(100),
					Title:                    "title",
					Content:                  "content",
					UpVotes:                  10,
					DownVotes:                5,
					SuggestionsCount:         2,
					AcceptedSuggestionsCount: 1,
					RevisionCount:            3,
				},
				{
					Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(20), baseTime, nil),
					UserID:                   goframework.NumberUUID(200),
					Title:                    "title",
					Content:                  "content",
					UpVotes:                  32,
					DownVotes:                16,
					SuggestionsCount:         8,
					AcceptedSuggestionsCount: 4,
					RevisionCount:            2,
				},
			},
			queryTotal: 20,
			expectedResults: []*models.ImproveRequestPreview{
				{
					ID:                       goframework.NumberUUID(10),
					CreatedAt:                baseTime,
					UserID:                   goframework.NumberUUID(100),
					Title:                    "title",
					Content:                  "content",
					UpVotes:                  10,
					DownVotes:                5,
					SuggestionsCount:         2,
					AcceptedSuggestionsCount: 1,
					RevisionCount:            3,
				},
				{
					ID:                       goframework.NumberUUID(20),
					CreatedAt:                baseTime,
					UserID:                   goframework.NumberUUID(200),
					Title:                    "title",
					Content:                  "content",
					UpVotes:                  32,
					DownVotes:                16,
					SuggestionsCount:         8,
					AcceptedSuggestionsCount: 4,
					RevisionCount:            2,
				},
			},
			expectedTotal: 20,
		},
		{
			name: "Success/NoResults",
			query: models.SearchImproveRequestsQuery{
				Limit: 10,
			},
			shouldCallDAO:   true,
			queryTotal:      20,
			expectedResults: []*models.ImproveRequestPreview{},
			expectedTotal:   20,
		},
		{
			name: "Success/WithQuery",
			query: models.SearchImproveRequestsQuery{
				UserID: apis.StringUUID(goframework.NumberUUID(1).String()),
				Query:  "foo bar",
				Order:  models.OrderScore,
				Limit:  10,
			},
			shouldCallDAO: true,
			shouldCallDAOWithForm: dao.ImproveRequestSearchQuery{
				UserID: lo.ToPtr(goframework.NumberUUID(1)),
				Query:  "foo bar",
				Order:  &dao.ImproveRequestSearchQueryOrder{Score: true},
			},
			queryTotal:      20,
			expectedResults: []*models.ImproveRequestPreview{},
			expectedTotal:   20,
		},
		{
			name: "Success/WithQueryInvalid",
			query: models.SearchImproveRequestsQuery{
				UserID: apis.StringUUID("invalid uuid"),
				Query:  "foo bar",
				Order:  "invalid order",
				Limit:  10,
			},
			shouldCallDAO: true,
			shouldCallDAOWithForm: dao.ImproveRequestSearchQuery{
				Query: "foo bar",
			},
			queryTotal:      20,
			expectedResults: []*models.ImproveRequestPreview{},
			expectedTotal:   20,
		},
		{
			name: "Error/DAOFailure",
			query: models.SearchImproveRequestsQuery{
				Limit: 10,
			},
			shouldCallDAO: true,
			queryErr:      fooErr,
			expectedErr:   fooErr,
		},
		{
			name:        "Error/NoLimit",
			expectedErr: goframework.ErrInvalidEntity,
		},
		{
			name: "Error/LimitTooHigh",
			query: models.SearchImproveRequestsQuery{
				Limit: services.MaxSearchLimit + 1,
			},
			expectedErr: goframework.ErrInvalidEntity,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveRequestRepository(t)

			if d.shouldCallDAO {
				repository.
					On("Search", context.Background(), d.shouldCallDAOWithForm, d.query.Limit, d.query.Offset).
					Return(d.queryResults, d.queryTotal, d.queryErr)
			}

			service := services.NewSearchImproveRequestsService(repository)
			results, total, err := service.Search(context.Background(), d.query)

			require.ErrorIs(t, err, d.expectedErr)
			require.Equal(t, d.expectedResults, results)
			require.Equal(t, d.expectedTotal, total)

			repository.AssertExpectations(t)
		})
	}
}
