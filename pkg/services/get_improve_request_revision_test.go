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

func TestGetImproveRequestRevisionService(t *testing.T) {
	data := []struct {
		name string

		id uuid.UUID

		daoResp *dao.ImproveRequestRevisionModel
		daoErr  error

		expect    *models.ImproveRequestRevision
		expectErr error
	}{
		{
			name: "Success",
			id:   goframework.NumberUUID(1),
			daoResp: &dao.ImproveRequestRevisionModel{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(22), baseTime.Add(time.Hour), &updateTime),
				SourceID: goframework.NumberUUID(20),
				UserID:   goframework.NumberUUID(201),
				Title:    "title",
				Content:  "content",
			},
			expect: &models.ImproveRequestRevision{
				ID:        goframework.NumberUUID(22),
				CreatedAt: baseTime.Add(time.Hour),
				SourceID:  goframework.NumberUUID(20),
				UserID:    goframework.NumberUUID(201),
				Title:     "title",
				Content:   "content",
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

			repository.On("GetRevision", context.Background(), d.id).Return(d.daoResp, d.daoErr)

			service := services.NewGetImproveRequestRevisionService(repository)

			resp, err := service.Get(context.Background(), d.id)

			require.ErrorIs(t, err, d.expectErr)
			require.Equal(t, d.expect, resp)

			repository.AssertExpectations(t)
		})
	}
}
