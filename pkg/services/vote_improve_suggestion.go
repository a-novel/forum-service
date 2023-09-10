package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/go-framework/errors"
	"github.com/google/uuid"
)

type VoteImproveSuggestionService interface {
	Vote(ctx context.Context, id, userID uuid.UUID, upVotes, downVotes int) error
}

func NewVoteImproveSuggestionService(repository dao.ImproveSuggestionRepository) VoteImproveSuggestionService {
	return &voteImproveSuggestionServiceImpl{
		repository: repository,
	}
}

type voteImproveSuggestionServiceImpl struct {
	repository dao.ImproveSuggestionRepository
}

func (s *voteImproveSuggestionServiceImpl) Vote(ctx context.Context, id, userID uuid.UUID, upVotes, downVotes int) error {
	suggestion, err := s.repository.Get(ctx, id)
	if err != nil {
		return goerrors.Join(ErrGetImproveSuggestion, err)
	}

	// User is not allowed to vote on its own post.
	if suggestion.UserID == userID {
		return goerrors.Join(errors.ErrInvalidCredentials, ErrTheCreator)
	}

	if err := s.repository.UpdateVotes(ctx, id, upVotes, downVotes); err != nil {
		return goerrors.Join(ErrUpdateImproveSuggestion, err)
	}

	return nil
}
