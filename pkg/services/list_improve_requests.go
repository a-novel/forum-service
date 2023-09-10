package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type ListImproveRequestsService interface {
	List(ctx context.Context, ids []uuid.UUID) ([]*models.ImproveRequestPreview, error)
}

func NewListImproveRequestsService(repository dao.ImproveRequestRepository) ListImproveRequestsService {
	return &listImproveRequestsServiceImpl{
		repository: repository,
	}
}

type listImproveRequestsServiceImpl struct {
	repository dao.ImproveRequestRepository
}

func (s *listImproveRequestsServiceImpl) List(ctx context.Context, ids []uuid.UUID) ([]*models.ImproveRequestPreview, error) {
	res, err := s.repository.List(ctx, ids)
	if err != nil {
		return nil, goerrors.Join(ErrListImproveRequests, err)
	}

	return lo.Map(res, func(item *dao.ImproveRequestPreview, _ int) *models.ImproveRequestPreview {
		return ParseImproveRequestPreview(item)
	}), nil
}
