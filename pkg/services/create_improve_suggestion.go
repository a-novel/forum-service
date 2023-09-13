package services

import (
	"context"
	goerrors "errors"
	auth "github.com/a-novel/auth-service/framework"
	"github.com/a-novel/forum-service/pkg/adapters"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/go-framework/errors"
	"github.com/a-novel/go-framework/validation"
	"github.com/google/uuid"
	"time"
)

type CreateImproveSuggestionService interface {
	Create(ctx context.Context, tokenRaw string, suggestion *models.ImproveSuggestionForm, id uuid.UUID, now time.Time) (*models.ImproveSuggestion, error)
}

func NewCreateImproveSuggestionService(
	repository dao.ImproveSuggestionRepository,
	requestRepository dao.ImproveRequestRepository,
	authClient auth.Client,
) CreateImproveSuggestionService {
	return &createImproveSuggestionServiceImpl{
		repository:        repository,
		requestRepository: requestRepository,
		authClient:        authClient,
	}
}

type createImproveSuggestionServiceImpl struct {
	repository        dao.ImproveSuggestionRepository
	requestRepository dao.ImproveRequestRepository
	authClient        auth.Client
}

func (s *createImproveSuggestionServiceImpl) Create(ctx context.Context, tokenRaw string, form *models.ImproveSuggestionForm, id uuid.UUID, now time.Time) (*models.ImproveSuggestion, error) {
	token, err := s.authClient.IntrospectToken(ctx, tokenRaw)
	if err != nil {
		return nil, goerrors.Join(ErrIntrospectToken, err)
	}
	if !token.OK {
		return nil, goerrors.Join(errors.ErrInvalidCredentials, ErrInvalidToken)
	}

	if err := validation.CheckMinMax(form.Title, MinTitleLength, MaxTitleLength); err != nil {
		return nil, goerrors.Join(errors.ErrInvalidEntity, ErrInvalidTitle, err)
	}
	if err := validation.CheckMinMax(form.Content, MinContentLength, MaxContentLength); err != nil {
		return nil, goerrors.Join(errors.ErrInvalidEntity, ErrInvalidContent, err)
	}
	if err := validation.CheckRegexp(form.Title, titleRegexp); err != nil {
		return nil, goerrors.Join(errors.ErrInvalidEntity, ErrInvalidTitle, err)
	}

	revision, err := s.requestRepository.GetRevision(ctx, form.RequestID)
	if err != nil {
		return nil, goerrors.Join(ErrGetImproveRequestRevision, err)
	}

	suggestion, err := s.repository.Create(ctx, adapters.ImproveSuggestionFormToDAO(form), token.Token.Payload.ID, revision.SourceID, id, now)
	if err != nil {
		return nil, goerrors.Join(ErrCreateImproveSuggestion, err)
	}

	return adapters.ImproveSuggestionToModel(suggestion), nil
}
