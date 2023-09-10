package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/go-framework/errors"
	"github.com/google/uuid"
)

type VoteImproveRequestService interface {
	Vote(ctx context.Context, id, userID uuid.UUID, upVotes, downVotes int) error
}

func NewVoteImproveRequestService(repository dao.ImproveRequestRepository) VoteImproveRequestService {
	return &voteImproveRequestServiceImpl{
		repository: repository,
	}
}

type voteImproveRequestServiceImpl struct {
	repository dao.ImproveRequestRepository
}

func (s *voteImproveRequestServiceImpl) Vote(ctx context.Context, id, userID uuid.UUID, upVotes, downVotes int) error {
	request, err := s.repository.Get(ctx, id)
	if err != nil {
		return goerrors.Join(ErrGetImproveRequestRevision, err)
	}

	// User is not allowed to vote on its own post.
	if request.UserID == userID {
		return goerrors.Join(errors.ErrInvalidCredentials, ErrTheCreator)
	}

	if err := s.repository.UpdateVotes(ctx, id, upVotes, downVotes); err != nil {
		return goerrors.Join(ErrUpdateImproveRequestRevision, err)
	}

	return nil
}
