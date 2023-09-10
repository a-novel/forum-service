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
	"time"
)

func TestSearchImproveSuggestionsService(t *testing.T) {
	data := []struct {
		name string

		query models.SearchImproveSuggestionsQuery

		shouldCallDAO         bool
		shouldCallDAOWithForm dao.ImproveSuggestionSearchQuery
		queryResults          []*dao.ImproveSuggestionModel
		queryTotal            int
		queryErr              error

		expectedResults []*models.ImproveSuggestion
		expectedTotal   int
		expectedErr     error
	}{
		{
			name: "Success",
			query: models.SearchImproveSuggestionsQuery{
				Limit: 10,
			},
			shouldCallDAO: true,
			queryResults: []*dao.ImproveSuggestionModel{
				{
					Metadata:  postgresql.NewMetadata(test.NumberUUID(1), baseTime, lo.ToPtr(baseTime.Add(3*time.Hour))),
					SourceID:  test.NumberUUID(10),
					UserID:    test.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: test.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  postgresql.NewMetadata(test.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
					SourceID:  test.NumberUUID(20),
					UserID:    test.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: test.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
			},
			queryTotal: 20,
			expectedResults: []*models.ImproveSuggestion{
				{
					ID:        test.NumberUUID(1),
					CreatedAt: baseTime,
					UpdatedAt: lo.ToPtr(baseTime.Add(3 * time.Hour)),
					SourceID:  test.NumberUUID(10),
					UserID:    test.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					RequestID: test.NumberUUID(1),
					Title:     "title",
					Content:   "content",
				},
				{
					ID:        test.NumberUUID(2),
					CreatedAt: baseTime,
					UpdatedAt: lo.ToPtr(baseTime.Add(2 * time.Hour)),
					SourceID:  test.NumberUUID(20),
					UserID:    test.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					RequestID: test.NumberUUID(1),
					Title:     "title",
					Content:   "content",
				},
			},
			expectedTotal: 20,
		},
		{
			name: "Success/NoResults",
			query: models.SearchImproveSuggestionsQuery{
				Limit: 10,
			},
			shouldCallDAO:   true,
			queryTotal:      20,
			expectedResults: []*models.ImproveSuggestion{},
			expectedTotal:   20,
		},
		{
			name: "Success/WithQuery",
			query: models.SearchImproveSuggestionsQuery{
				UserID:    types.StringUUID(test.NumberUUID(1).String()),
				SourceID:  types.StringUUID(test.NumberUUID(2).String()),
				RequestID: types.StringUUID(test.NumberUUID(3).String()),
				Validated: lo.ToPtr(true),
				Order:     models.OrderScore,
				Limit:     10,
			},
			shouldCallDAO: true,
			shouldCallDAOWithForm: dao.ImproveSuggestionSearchQuery{
				UserID:    lo.ToPtr(test.NumberUUID(1)),
				SourceID:  lo.ToPtr(test.NumberUUID(2)),
				RequestID: lo.ToPtr(test.NumberUUID(3)),
				Validated: lo.ToPtr(true),
				Order:     &dao.ImproveSuggestionSearchQueryOrder{Score: true},
			},
			queryTotal:      20,
			expectedResults: []*models.ImproveSuggestion{},
			expectedTotal:   20,
		},
		{
			name: "Success/WithQueryInvalid",
			query: models.SearchImproveSuggestionsQuery{
				UserID:    types.StringUUID("invalid uuid"),
				SourceID:  types.StringUUID("invalid uuid"),
				RequestID: types.StringUUID("invalid uuid"),
				Order:     "invalid order",
				Limit:     10,
			},
			shouldCallDAO:   true,
			queryTotal:      20,
			expectedResults: []*models.ImproveSuggestion{},
			expectedTotal:   20,
		},
		{
			name: "Error/DAOFailure",
			query: models.SearchImproveSuggestionsQuery{
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
			query: models.SearchImproveSuggestionsQuery{
				Limit: services.MaxSearchLimit + 1,
			},
			expectedErr: errors.ErrInvalidEntity,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveSuggestionRepository(t)

			if d.shouldCallDAO {
				repository.
					On("Search", context.Background(), d.shouldCallDAOWithForm, d.query.Limit, d.query.Offset).
					Return(d.queryResults, d.queryTotal, d.queryErr)
			}

			service := services.NewSearchImproveSuggestionsService(repository)
			results, total, err := service.Search(context.Background(), d.query)

			require.ErrorIs(t, err, d.expectedErr)
			require.Equal(t, d.expectedResults, results)
			require.Equal(t, d.expectedTotal, total)

			repository.AssertExpectations(t)
		})
	}
}
