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
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, lo.ToPtr(baseTime.Add(3*time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
					SourceID:  goframework.NumberUUID(20),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
			},
			queryTotal: 20,
			expectedResults: []*models.ImproveSuggestion{
				{
					ID:        goframework.NumberUUID(1),
					CreatedAt: baseTime,
					UpdatedAt: lo.ToPtr(baseTime.Add(3 * time.Hour)),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					RequestID: goframework.NumberUUID(1),
					Title:     "title",
					Content:   "content",
				},
				{
					ID:        goframework.NumberUUID(2),
					CreatedAt: baseTime,
					UpdatedAt: lo.ToPtr(baseTime.Add(2 * time.Hour)),
					SourceID:  goframework.NumberUUID(20),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					RequestID: goframework.NumberUUID(1),
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
				UserID:    apis.StringUUID(goframework.NumberUUID(1).String()),
				SourceID:  apis.StringUUID(goframework.NumberUUID(2).String()),
				RequestID: apis.StringUUID(goframework.NumberUUID(3).String()),
				Validated: lo.ToPtr(true),
				Order:     models.OrderScore,
				Limit:     10,
			},
			shouldCallDAO: true,
			shouldCallDAOWithForm: dao.ImproveSuggestionSearchQuery{
				UserID:    lo.ToPtr(goframework.NumberUUID(1)),
				SourceID:  lo.ToPtr(goframework.NumberUUID(2)),
				RequestID: lo.ToPtr(goframework.NumberUUID(3)),
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
				UserID:    apis.StringUUID("invalid uuid"),
				SourceID:  apis.StringUUID("invalid uuid"),
				RequestID: apis.StringUUID("invalid uuid"),
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
			expectedErr: goframework.ErrInvalidEntity,
		},
		{
			name: "Error/LimitTooHigh",
			query: models.SearchImproveSuggestionsQuery{
				Limit: services.MaxSearchLimit + 1,
			},
			expectedErr: goframework.ErrInvalidEntity,
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
