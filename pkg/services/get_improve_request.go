package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/google/uuid"
)

type GetImproveRequestService interface {
	Get(ctx context.Context, id uuid.UUID) (*models.ImproveRequestPreview, error)
}

func NewGetImproveRequestService(repository dao.ImproveRequestRepository) GetImproveRequestService {
	return &getImproveRequestServiceImpl{
		repository: repository,
	}
}

type getImproveRequestServiceImpl struct {
	repository dao.ImproveRequestRepository
}

func (s *getImproveRequestServiceImpl) Get(ctx context.Context, id uuid.UUID) (*models.ImproveRequestPreview, error) {
	data, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, goerrors.Join(ErrGetImproveRequest, err)
	}

	return ParseImproveRequestPreview(data), nil
}
