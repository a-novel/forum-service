package services

import (
	"context"
	goerrors "errors"
	auth "github.com/a-novel/auth-service/framework"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/pkg/adapters"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"time"
)

type CreateImproveRequestService interface {
	Create(ctx context.Context, tokenRaw, title, content string, sourceID, id uuid.UUID, now time.Time) (*models.ImproveRequestPreview, error)
}

func NewCreateImproveRequestService(repository dao.ImproveRequestRepository, authClient auth.Client) CreateImproveRequestService {
	return &createImproveRequestServiceImpl{
		repository: repository,
		authClient: authClient,
	}
}

type createImproveRequestServiceImpl struct {
	repository dao.ImproveRequestRepository
	authClient auth.Client
}

func (s *createImproveRequestServiceImpl) Create(ctx context.Context, tokenRaw, title, content string, sourceID, id uuid.UUID, now time.Time) (*models.ImproveRequestPreview, error) {
	token, err := s.authClient.IntrospectToken(ctx, tokenRaw)
	if err != nil {
		return nil, goerrors.Join(ErrIntrospectToken, err)
	}
	if !token.OK {
		return nil, goerrors.Join(goframework.ErrInvalidCredentials, ErrInvalidToken)
	}

	if err := goframework.CheckMinMax(title, MinTitleLength, MaxTitleLength); err != nil {
		return nil, goerrors.Join(goframework.ErrInvalidEntity, ErrInvalidTitle, err)
	}
	if err := goframework.CheckMinMax(content, MinContentLength, MaxContentLength); err != nil {
		return nil, goerrors.Join(goframework.ErrInvalidEntity, ErrInvalidContent, err)
	}
	if err := goframework.CheckRegexp(title, titleRegexp); err != nil {
		return nil, goerrors.Join(goframework.ErrInvalidEntity, ErrInvalidTitle, err)
	}

	request, err := s.repository.Get(ctx, sourceID)
	if err != nil && !goerrors.Is(err, bunovel.ErrNotFound) {
		return nil, goerrors.Join(ErrGetImproveRequestRevision, err)
	}

	// Only the original creator can make revisions on a post.
	if request != nil && request.UserID != token.Token.Payload.ID {
		return nil, goerrors.Join(goframework.ErrInvalidCredentials, ErrNotTheCreator)
	}

	res, err := s.repository.Create(ctx, token.Token.Payload.ID, title, content, sourceID, id, now)
	if err != nil {
		return nil, goerrors.Join(ErrCreateImproveRequest, err)
	}

	return adapters.ImproveRequestPreviewToModel(res), nil
}
