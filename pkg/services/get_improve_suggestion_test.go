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
)

func TestGetImproveSuggestionService(t *testing.T) {
	data := []struct {
		name string

		id uuid.UUID

		daoResp *dao.ImproveSuggestionModel
		daoErr  error

		expect    *models.ImproveSuggestion
		expectErr error
	}{
		{
			name: "Success",
			id:   goframework.NumberUUID(1),
			daoResp: &dao.ImproveSuggestionModel{
				Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
				SourceID:  goframework.NumberUUID(10),
				UserID:    goframework.NumberUUID(100),
				Validated: true,
				UpVotes:   128,
				DownVotes: 64,
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: goframework.NumberUUID(1),
					Title:     "suggestion title",
					Content:   "suggestion content",
				},
			},
			expect: &models.ImproveSuggestion{
				ID:        goframework.NumberUUID(1),
				CreatedAt: baseTime,
				UpdatedAt: &updateTime,
				SourceID:  goframework.NumberUUID(10),
				UserID:    goframework.NumberUUID(100),
				Validated: true,
				UpVotes:   128,
				DownVotes: 64,
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
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
			repository := daomocks.NewImproveSuggestionRepository(t)

			repository.On("Get", context.Background(), d.id).Return(d.daoResp, d.daoErr)

			service := services.NewGetImproveSuggestionService(repository)
			result, err := service.Get(context.Background(), d.id)

			require.ErrorIs(t, err, d.expectErr)
			require.Equal(t, d.expect, result)

			repository.AssertExpectations(t)
		})
	}
}
