package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/adapters"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	goframework "github.com/a-novel/go-framework"
	"github.com/samber/lo"
)

type SearchImproveRequestsService interface {
	Search(ctx context.Context, query models.SearchImproveRequestsQuery) ([]*models.ImproveRequestPreview, int, error)
}

func NewSearchImproveRequestsService(repository dao.ImproveRequestRepository) SearchImproveRequestsService {
	return &searchImproveRequestsServiceImpl{
		repository: repository,
	}
}

type searchImproveRequestsServiceImpl struct {
	repository dao.ImproveRequestRepository
}

func (s *searchImproveRequestsServiceImpl) Search(ctx context.Context, query models.SearchImproveRequestsQuery) ([]*models.ImproveRequestPreview, int, error) {
	if err := goframework.CheckMinMax(query.Limit, 1, MaxSearchLimit); err != nil {
		return nil, 0, goerrors.Join(goframework.ErrInvalidEntity, ErrInvalidSearchLimit, err)
	}

	res, total, err := s.repository.Search(ctx, adapters.ImproveRequestSearchQueryToDAO(query), query.Limit, query.Offset)
	if err != nil {
		return nil, 0, goerrors.Join(ErrSearchImproveRequests, err)
	}

	return lo.Map(res, func(item *dao.ImproveRequestPreview, _ int) *models.ImproveRequestPreview {
		return adapters.ImproveRequestPreviewToModel(item)
	}), total, nil
}
