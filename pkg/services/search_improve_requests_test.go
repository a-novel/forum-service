package services_test

import (
	"context"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/errors"
	"github.com/a-novel/go-framework/postgresql"
	"github.com/a-novel/go-framework/test"
	"github.com/a-novel/go-framework/types"
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
					Metadata:                 postgresql.NewMetadata(test.NumberUUID(10), baseTime, nil),
					UserID:                   test.NumberUUID(100),
					Title:                    "title",
					Content:                  "content",
					UpVotes:                  10,
					DownVotes:                5,
					SuggestionsCount:         2,
					AcceptedSuggestionsCount: 1,
					RevisionCount:            3,
				},
				{
					Metadata:                 postgresql.NewMetadata(test.NumberUUID(20), baseTime, nil),
					UserID:                   test.NumberUUID(200),
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
					ID:                       test.NumberUUID(10),
					CreatedAt:                baseTime,
					UserID:                   test.NumberUUID(100),
					Title:                    "title",
					Content:                  "content",
					UpVotes:                  10,
					DownVotes:                5,
					SuggestionsCount:         2,
					AcceptedSuggestionsCount: 1,
					RevisionCount:            3,
				},
				{
					ID:                       test.NumberUUID(20),
					CreatedAt:                baseTime,
					UserID:                   test.NumberUUID(200),
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
				UserID: types.StringUUID(test.NumberUUID(1).String()),
				Query:  "foo bar",
				Order:  models.OrderScore,
				Limit:  10,
			},
			shouldCallDAO: true,
			shouldCallDAOWithForm: dao.ImproveRequestSearchQuery{
				UserID: lo.ToPtr(test.NumberUUID(1)),
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
				UserID: types.StringUUID("invalid uuid"),
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
			expectedErr: errors.ErrInvalidEntity,
		},
		{
			name: "Error/LimitTooHigh",
			query: models.SearchImproveRequestsQuery{
				Limit: services.MaxSearchLimit + 1,
			},
			expectedErr: errors.ErrInvalidEntity,
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
