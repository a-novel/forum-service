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
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListImproveRequestsService(t *testing.T) {
	data := []struct {
		name string

		ids []uuid.UUID

		daoResp []*dao.ImproveRequestPreview
		daoErr  error

		expected    []*models.ImproveRequestPreview
		expectedErr error
	}{
		{
			name: "Success",
			ids:  []uuid.UUID{test.NumberUUID(1), test.NumberUUID(2)},
			daoResp: []*dao.ImproveRequestPreview{
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
			expected: []*models.ImproveRequestPreview{
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
		},
		{
			name:     "Success/NoResults",
			ids:      []uuid.UUID{test.NumberUUID(1), test.NumberUUID(2)},
			expected: []*models.ImproveRequestPreview{},
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
			repository := daomocks.NewImproveRequestRepository(t)
			repository.On("List", context.Background(), d.ids).Return(d.daoResp, d.daoErr)

			service := services.NewListImproveRequestsService(repository)
			resp, err := service.List(context.Background(), d.ids)

			require.ErrorIs(t, err, d.expectedErr)
			require.Equal(t, d.expected, resp)

			repository.AssertExpectations(t)
		})
	}
}
