package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/adapters"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	apiclients "github.com/a-novel/go-api-clients"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"time"
)

type CreateImproveSuggestionService interface {
	Create(ctx context.Context, tokenRaw string, suggestion *models.ImproveSuggestionForm, id uuid.UUID, now time.Time) (*models.ImproveSuggestion, error)
}

func NewCreateImproveSuggestionService(
	repository dao.ImproveSuggestionRepository,
	requestRepository dao.ImproveRequestRepository,
	authClient apiclients.AuthClient,
	permissionsClient apiclients.PermissionsClient,
) CreateImproveSuggestionService {
	return &createImproveSuggestionServiceImpl{
		repository:        repository,
		requestRepository: requestRepository,
		authClient:        authClient,
		permissionsClient: permissionsClient,
	}
}

type createImproveSuggestionServiceImpl struct {
	repository        dao.ImproveSuggestionRepository
	requestRepository dao.ImproveRequestRepository
	authClient        apiclients.AuthClient
	permissionsClient apiclients.PermissionsClient
}

func (s *createImproveSuggestionServiceImpl) Create(ctx context.Context, tokenRaw string, form *models.ImproveSuggestionForm, id uuid.UUID, now time.Time) (*models.ImproveSuggestion, error) {
	token, err := s.authClient.IntrospectToken(ctx, tokenRaw)
	if err != nil {
		return nil, goerrors.Join(ErrIntrospectToken, err)
	}
	if !token.OK {
		return nil, goerrors.Join(goframework.ErrInvalidCredentials, ErrInvalidToken)
	}

	if err := s.permissionsClient.HasUserScope(ctx, apiclients.HasUserScopeQuery{
		UserID: token.Token.Payload.ID,
		Scope:  apiclients.CanPostImproveSuggestion,
	}); err != nil {
		return nil, goerrors.Join(ErrGetScopes, err)
	}

	if err := goframework.CheckMinMax(form.Title, MinTitleLength, MaxTitleLength); err != nil {
		return nil, goerrors.Join(goframework.ErrInvalidEntity, ErrInvalidTitle, err)
	}
	if err := goframework.CheckMinMax(form.Content, MinContentLength, MaxContentLength); err != nil {
		return nil, goerrors.Join(goframework.ErrInvalidEntity, ErrInvalidContent, err)
	}
	if err := goframework.CheckRegexp(form.Title, titleRegexp); err != nil {
		return nil, goerrors.Join(goframework.ErrInvalidEntity, ErrInvalidTitle, err)
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
