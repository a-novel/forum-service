package services_test

import (
	"context"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/forum-service/pkg/services"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetImproveRequestService(t *testing.T) {
	data := []struct {
		name string

		id uuid.UUID

		daoResp *dao.ImproveRequestPreview
		daoErr  error

		expect    *models.ImproveRequestPreview
		expectErr error
	}{
		{
			name: "Success",
			id:   goframework.NumberUUID(1),
			daoResp: &dao.ImproveRequestPreview{
				Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(21), baseTime.Add(time.Hour), &updateTime),
				UserID:                   goframework.NumberUUID(201),
				Title:                    "title",
				Content:                  "content",
				UpVotes:                  256,
				DownVotes:                128,
				RevisionCount:            10,
				SuggestionsCount:         20,
				AcceptedSuggestionsCount: 5,
			},
			expect: &models.ImproveRequestPreview{
				ID:                       goframework.NumberUUID(21),
				CreatedAt:                baseTime.Add(time.Hour),
				UserID:                   goframework.NumberUUID(201),
				Title:                    "title",
				Content:                  "content",
				UpVotes:                  256,
				DownVotes:                128,
				RevisionCount:            10,
				SuggestionsCount:         20,
				AcceptedSuggestionsCount: 5,
			},
		},
		{
			name:      "Error/DAOFailure",
			id:        goframework.NumberUUID(1),
			daoErr:    fooErr,
			expectErr: fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveRequestRepository(t)

			repository.On("Get", context.Background(), d.id).Return(d.daoResp, d.daoErr)

			service := services.NewGetImproveRequestService(repository)

			resp, err := service.Get(context.Background(), d.id)

			require.ErrorIs(t, err, d.expectErr)
			require.Equal(t, d.expect, resp)

			repository.AssertExpectations(t)
		})
	}
}
