package services

import (
	"context"
	goerrors "errors"
	auth "github.com/a-novel/auth-service/framework"
	"github.com/a-novel/forum-service/pkg/adapters"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"time"
)

type UpdateImproveSuggestionService interface {
	Update(ctx context.Context, tokenRaw string, suggestion *models.ImproveSuggestionForm, id uuid.UUID, now time.Time) (*models.ImproveSuggestion, error)
}

func NewUpdateImproveSuggestionService(
	repository dao.ImproveSuggestionRepository,
	requestRepository dao.ImproveRequestRepository,
	authClient auth.Client,
) UpdateImproveSuggestionService {
	return &updateImproveSuggestionServiceImpl{
		repository:        repository,
		requestRepository: requestRepository,
		authClient:        authClient,
	}
}

type updateImproveSuggestionServiceImpl struct {
	repository        dao.ImproveSuggestionRepository
	requestRepository dao.ImproveRequestRepository
	authClient        auth.Client
}

func (s *updateImproveSuggestionServiceImpl) Update(ctx context.Context, tokenRaw string, form *models.ImproveSuggestionForm, id uuid.UUID, now time.Time) (*models.ImproveSuggestion, error) {
	token, err := s.authClient.IntrospectToken(ctx, tokenRaw)
	if err != nil {
		return nil, goerrors.Join(ErrIntrospectToken, err)
	}
	if !token.OK {
		return nil, goerrors.Join(goframework.ErrInvalidCredentials, ErrInvalidToken)
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

	suggestion, err := s.repository.Get(ctx, form.RequestID)
	if err != nil {
		return nil, goerrors.Join(ErrGetImproveSuggestion, err)
	}

	if revision.SourceID != suggestion.SourceID {
		return nil, goerrors.Join(goframework.ErrInvalidEntity, ErrSwitchSource)
	}

	suggestion, err = s.repository.Update(ctx, adapters.ImproveSuggestionFormToDAO(form), id, now)
	if err != nil {
		return nil, goerrors.Join(ErrUpdateImproveSuggestion, err)
	}

	return adapters.ImproveSuggestionToModel(suggestion), nil
}
