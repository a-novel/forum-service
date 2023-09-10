package services_test

import (
	"context"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/postgresql"
	"github.com/a-novel/go-framework/test"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestListImproveSuggestionsService(t *testing.T) {
	data := []struct {
		name string

		ids []uuid.UUID

		daoResp []*dao.ImproveSuggestionModel
		daoErr  error

		expected    []*models.ImproveSuggestion
		expectedErr error
	}{
		{
			name: "Success",
			ids:  []uuid.UUID{test.NumberUUID(1), test.NumberUUID(2)},
			daoResp: []*dao.ImproveSuggestionModel{
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
			expected: []*models.ImproveSuggestion{
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
		},
		{
			name:     "Success/NoResults",
			ids:      []uuid.UUID{test.NumberUUID(1), test.NumberUUID(2)},
			expected: []*models.ImproveSuggestion{},
		},
		{
			name:        "Error/DAOFailure",
			ids:         []uuid.UUID{test.NumberUUID(1), test.NumberUUID(2)},
			daoErr:      fooErr,
			expectedErr: fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveSuggestionRepository(t)
			repository.On("List", context.Background(), d.ids).Return(d.daoResp, d.daoErr)

			service := services.NewListImproveSuggestionsService(repository)
			resp, err := service.List(context.Background(), d.ids)

			require.ErrorIs(t, err, d.expectedErr)
			require.Equal(t, d.expected, resp)

			repository.AssertExpectations(t)
		})
	}
}
