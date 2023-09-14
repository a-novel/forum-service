package services

import (
	"context"
	goerrors "errors"
	auth "github.com/a-novel/auth-service/framework"
	"github.com/a-novel/forum-service/pkg/dao"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
)

type ValidateImproveSuggestionService interface {
	Validate(ctx context.Context, tokenRaw string, validated bool, id uuid.UUID) error
}

func NewValidateImproveSuggestionService(
	repository dao.ImproveSuggestionRepository,
	requestRepository dao.ImproveRequestRepository,
	authClient auth.Client,
) ValidateImproveSuggestionService {
	return &validateImproveSuggestionServiceImpl{
		repository:        repository,
		requestRepository: requestRepository,
		authClient:        authClient,
	}
}

type validateImproveSuggestionServiceImpl struct {
	repository        dao.ImproveSuggestionRepository
	requestRepository dao.ImproveRequestRepository
	authClient        auth.Client
}

func (s *validateImproveSuggestionServiceImpl) Validate(ctx context.Context, tokenRaw string, validated bool, id uuid.UUID) error {
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

	request, err := s.requestRepository.GetRevision(ctx, suggestion.RequestID)
	if err != nil {
		return goerrors.Join(ErrGetImproveRequestRevision, err)
	}

	if request.UserID != token.Token.Payload.ID {
		return goerrors.Join(goframework.ErrInvalidCredentials, ErrNotTheCreator)
	}

	if _, err := s.repository.Validate(ctx, validated, id); err != nil {
		return goerrors.Join(ErrValidateImproveSuggestion, err)
	}

	return nil
}
