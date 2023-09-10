package services_test

import (
	"context"
	"github.com/a-novel/forum-service/pkg/dao"
	daomocks "github.com/a-novel/forum-service/pkg/dao/mocks"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-framework/errors"
	"github.com/a-novel/go-framework/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVoteImproveSuggestionService(t *testing.T) {
	data := []struct {
		name string

		id        uuid.UUID
		userID    uuid.UUID
		upVotes   int
		downVotes int

		getRevision    *dao.ImproveSuggestionModel
		getRevisionErr error

		shouldCallUpdateVotes bool
		updateVotesErr        error

		expectErr error
	}{
		{
			name:      "Success",
			id:        test.NumberUUID(1),
			userID:    test.NumberUUID(100),
			upVotes:   10,
			downVotes: 5,
			getRevision: &dao.ImproveSuggestionModel{
				UserID: test.NumberUUID(200),
			},
			shouldCallUpdateVotes: true,
		},
		{
			name:      "Error/UpdateFailure",
			id:        test.NumberUUID(1),
			userID:    test.NumberUUID(100),
			upVotes:   10,
			downVotes: 5,
			getRevision: &dao.ImproveSuggestionModel{
				UserID: test.NumberUUID(200),
			},
			shouldCallUpdateVotes: true,
			updateVotesErr:        fooErr,
			expectErr:             fooErr,
		},
		{
			name:      "Error/IsTheCreator",
			id:        test.NumberUUID(1),
			userID:    test.NumberUUID(100),
			upVotes:   10,
			downVotes: 5,
			getRevision: &dao.ImproveSuggestionModel{
				UserID: test.NumberUUID(100),
			},
			expectErr: errors.ErrInvalidCredentials,
		},
		{
			name:           "Error/GetRevisionFailure",
			id:             test.NumberUUID(1),
			userID:         test.NumberUUID(100),
			upVotes:        10,
			downVotes:      5,
			getRevisionErr: fooErr,
			expectErr:      fooErr,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			repository := daomocks.NewImproveSuggestionRepository(t)

			repository.On("Get", context.Background(), d.id).Return(d.getRevision, d.getRevisionErr)

			if d.shouldCallUpdateVotes {
				repository.On("UpdateVotes", context.Background(), d.id, d.upVotes, d.downVotes).Return(d.updateVotesErr)
			}

			service := services.NewVoteImproveSuggestionService(repository)
			err := service.Vote(context.Background(), d.id, d.userID, d.upVotes, d.downVotes)

			require.ErrorIs(t, err, d.expectErr)

			repository.AssertExpectations(t)
		})
	}
}
