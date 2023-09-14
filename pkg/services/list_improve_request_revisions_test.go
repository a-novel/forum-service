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

func TestListImproveRequestRevisionService(t *testing.T) {
	data := []struct {
		name string

		id uuid.UUID

		daoResp []*dao.ImproveRequestRevisionPreview
		daoErr  error

		expect    []*models.ImproveRequestRevisionPreview
		expectErr error
	}{
		{
			name: "Success",
			id:   goframework.NumberUUID(1),
			daoResp: []*dao.ImproveRequestRevisionPreview{
				{
					Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(21), baseTime.Add(time.Hour), &updateTime),
					SuggestionsCount:         10,
					AcceptedSuggestionsCount: 5,
				},
				{
					Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(22), baseTime, &updateTime),
					SuggestionsCount:         2,
					AcceptedSuggestionsCount: 1,
				},
			},
			expect: []*models.ImproveRequestRevisionPreview{
				{
					ID:                       goframework.NumberUUID(21),
					CreatedAt:                baseTime.Add(time.Hour),
					SuggestionsCount:         10,
					AcceptedSuggestionsCount: 5,
				},
				{
					ID:                       goframework.NumberUUID(22),
					CreatedAt:                baseTime,
					SuggestionsCount:         2,
					AcceptedSuggestionsCount: 1,
				},
			},
		},
		{
			name:   "Error/DAOFailure",
			id:     goframework.NumberUUID(1),
			expect: []*models.ImproveRequestRevisionPreview{},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveRequestRepository(t)

			repository.On("ListRevisions", context.Background(), d.id).Return(d.daoResp, d.daoErr)

			service := services.NewListImproveRequestRevisionsService(repository)

			resp, err := service.List(context.Background(), d.id)

			require.ErrorIs(t, err, d.expectErr)
			require.Equal(t, d.expect, resp)

			repository.AssertExpectations(t)
		})
	}
}
