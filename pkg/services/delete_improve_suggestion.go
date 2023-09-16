package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/dao"
	apiclients "github.com/a-novel/go-api-clients"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
)

type DeleteImproveSuggestionService interface {
	Delete(ctx context.Context, tokenRaw string, id uuid.UUID) error
}

func NewDeleteImproveSuggestionService(repository dao.ImproveSuggestionRepository, authClient apiclients.AuthClient) DeleteImproveSuggestionService {
	return &deleteImproveSuggestionServiceImpl{
		repository: repository,
		authClient: authClient,
	}
}

type deleteImproveSuggestionServiceImpl struct {
	repository dao.ImproveSuggestionRepository
	authClient apiclients.AuthClient
}

func (s *deleteImproveSuggestionServiceImpl) Delete(ctx context.Context, tokenRaw string, id uuid.UUID) error {
	token, err := s.authClient.IntrospectToken(ctx, tokenRaw)
	if err != nil {
		return goerrors.Join(ErrIntrospectToken, err)
	}
	if !token.OK {
		return goerrors.Join(goframework.ErrInvalidCredentials, ErrInvalidToken)
	}

	suggestion, err := s.repository.Get(ctx, id)
	if err != nil {
		return goerrors.Join(ErrGetImproveSuggestion, err)
	}
	if suggestion.UserID != token.Token.Payload.ID {
		return goerrors.Join(goframework.ErrInvalidCredentials, ErrNotTheCreator)
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return goerrors.Join(ErrDeleteImproveSuggestion, err)
	}

	return nil
}
