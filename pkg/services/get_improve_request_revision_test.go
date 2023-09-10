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
			id:   test.NumberUUID(1),
			daoResp: &dao.ImproveRequestRevisionModel{
				Metadata: postgresql.NewMetadata(test.NumberUUID(22), baseTime.Add(time.Hour), &updateTime),
				SourceID: test.NumberUUID(20),
				UserID:   test.NumberUUID(201),
				Title:    "title",
				Content:  "content",
			},
			expect: &models.ImproveRequestRevision{
				ID:        test.NumberUUID(22),
				CreatedAt: baseTime.Add(time.Hour),
				SourceID:  test.NumberUUID(20),
				UserID:    test.NumberUUID(201),
				Title:     "title",
				Content:   "content",
			},
		},
		{
			name:      "Error/DAOFailure",
			id:        test.NumberUUID(1),
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
