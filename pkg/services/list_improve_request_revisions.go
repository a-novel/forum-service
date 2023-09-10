package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type ListImproveRequestRevisionsService interface {
	List(ctx context.Context, id uuid.UUID) ([]*models.ImproveRequestRevisionPreview, error)
}

func NewListImproveRequestRevisionsService(repository dao.ImproveRequestRepository) ListImproveRequestRevisionsService {
	return &listImproveRequestRevisionServiceImpl{
		repository: repository,
	}
}

type listImproveRequestRevisionServiceImpl struct {
	repository dao.ImproveRequestRepository
}

func (s *listImproveRequestRevisionServiceImpl) List(ctx context.Context, id uuid.UUID) ([]*models.ImproveRequestRevisionPreview, error) {
	data, err := s.repository.ListRevisions(ctx, id)
	if err != nil {
		return nil, goerrors.Join(ErrListImproveRequestRevisions, err)
	}

	return lo.Map(data, func(item *dao.ImproveRequestRevisionPreview, _ int) *models.ImproveRequestRevisionPreview {
		return ParseImproveRequestRevisionPreview(item)
	}), nil
}
