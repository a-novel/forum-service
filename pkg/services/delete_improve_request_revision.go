package services

import (
	"context"
	goerrors "errors"
	auth "github.com/a-novel/auth-service/framework"
	"github.com/a-novel/forum-service/pkg/dao"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
)

type DeleteImproveRequestRevisionService interface {
	Delete(ctx context.Context, tokenRaw string, id uuid.UUID) error
}

func NewDeleteImproveRequestRevisionService(repository dao.ImproveRequestRepository, authClient auth.Client) DeleteImproveRequestRevisionService {
	return &deleteImproveRequestRevisionServiceImpl{
		repository: repository,
		authClient: authClient,
	}
}

type deleteImproveRequestRevisionServiceImpl struct {
	repository dao.ImproveRequestRepository
	authClient auth.Client
}

func (s *deleteImproveRequestRevisionServiceImpl) Delete(ctx context.Context, tokenRaw string, id uuid.UUID) error {
	token, err := s.authClient.IntrospectToken(ctx, tokenRaw)
	if err != nil {
		return goerrors.Join(ErrIntrospectToken, err)
	}
	if !token.OK {
		return goerrors.Join(goframework.ErrInvalidCredentials, ErrInvalidToken)
	}

	revision, err := s.repository.GetRevision(ctx, id)
	if err != nil {
		return goerrors.Join(ErrGetImproveRequestRevision, err)
	}
	if revision.UserID != token.Token.Payload.ID {
		return goerrors.Join(goframework.ErrInvalidCredentials, ErrNotTheCreator)
	}

	if err := s.repository.DeleteRevision(ctx, id); err != nil {
		return goerrors.Join(ErrDeleteImproveRequestRevision, err)
	}

	return nil
}
