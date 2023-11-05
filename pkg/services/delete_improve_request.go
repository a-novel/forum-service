package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/dao"
	apiclients "github.com/a-novel/go-apis/clients"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
)

type DeleteImproveRequestService interface {
	Delete(ctx context.Context, tokenRaw string, id uuid.UUID) error
}

func NewDeleteImproveRequestService(repository dao.ImproveRequestRepository, authClient apiclients.AuthClient) DeleteImproveRequestService {
	return &deleteImproveRequestServiceImpl{
		repository: repository,
		authClient: authClient,
	}
}

type deleteImproveRequestServiceImpl struct {
	repository dao.ImproveRequestRepository
	authClient apiclients.AuthClient
}

func (s *deleteImproveRequestServiceImpl) Delete(ctx context.Context, tokenRaw string, id uuid.UUID) error {
	token, err := s.authClient.IntrospectToken(ctx, tokenRaw)
	if err != nil {
		return goerrors.Join(ErrIntrospectToken, err)
	}
	if !token.OK {
		return goerrors.Join(goframework.ErrInvalidCredentials, ErrInvalidToken)
	}

	request, err := s.repository.Get(ctx, id)
	if err != nil {
		return goerrors.Join(ErrGetImproveRequestRevision, err)
	}
	if request.UserID != token.Token.Payload.ID {
		return goerrors.Join(goframework.ErrInvalidCredentials, ErrNotTheCreator)
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return goerrors.Join(ErrDeleteImproveRequest, err)
	}

	return nil
}
