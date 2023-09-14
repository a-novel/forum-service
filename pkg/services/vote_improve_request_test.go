package services_test

import (
	"context"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/services"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVoteImproveRequestService(t *testing.T) {
	data := []struct {
		name string

		id        uuid.UUID
		userID    uuid.UUID
		upVotes   int
		downVotes int

		getRevision    *dao.ImproveRequestPreview
		getRevisionErr error

		shouldCallUpdateVotes bool
		updateVotesErr        error

		expectErr error
	}{
		{
			name:      "Success",
			id:        goframework.NumberUUID(1),
			userID:    goframework.NumberUUID(100),
			upVotes:   10,
			downVotes: 5,
			getRevision: &dao.ImproveRequestPreview{
				UserID: goframework.NumberUUID(200),
			},
			shouldCallUpdateVotes: true,
		},
		{
			name:      "Error/UpdateFailure",
			id:        goframework.NumberUUID(1),
			userID:    goframework.NumberUUID(100),
			upVotes:   10,
			downVotes: 5,
			getRevision: &dao.ImproveRequestPreview{
				UserID: goframework.NumberUUID(200),
			},
			shouldCallUpdateVotes: true,
			updateVotesErr:        fooErr,
			expectErr:             fooErr,
		},
		{
			name:      "Error/IsTheCreator",
			id:        goframework.NumberUUID(1),
			userID:    goframework.NumberUUID(100),
			upVotes:   10,
			downVotes: 5,
			getRevision: &dao.ImproveRequestPreview{
				UserID: goframework.NumberUUID(100),
			},
			expectErr: goframework.ErrInvalidCredentials,
		},
		{
			name:           "Error/GetRevisionFailure",
			id:             goframework.NumberUUID(1),
			userID:         goframework.NumberUUID(100),
			upVotes:        10,
			downVotes:      5,
			getRevisionErr: fooErr,
			expectErr:      fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveRequestRepository(t)

			repository.On("Get", context.Background(), d.id).Return(d.getRevision, d.getRevisionErr)

			if d.shouldCallUpdateVotes {
				repository.On("UpdateVotes", context.Background(), d.id, d.upVotes, d.downVotes).Return(d.updateVotesErr)
			}

			service := services.NewVoteImproveRequestService(repository)
			err := service.Vote(context.Background(), d.id, d.userID, d.upVotes, d.downVotes)

			require.ErrorIs(t, err, d.expectErr)

			repository.AssertExpectations(t)
		})
	}
}
