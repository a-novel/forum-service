package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/adapters"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/google/uuid"
)

type GetImproveRequestRevisionService interface {
	Get(ctx context.Context, id uuid.UUID) (*models.ImproveRequestRevision, error)
}

func NewGetImproveRequestRevisionService(repository dao.ImproveRequestRepository) GetImproveRequestRevisionService {
	return &getImproveRequestRevisionServiceImpl{
		repository: repository,
	}
}

type getImproveRequestRevisionServiceImpl struct {
	repository dao.ImproveRequestRepository
}

func (s *getImproveRequestRevisionServiceImpl) Get(ctx context.Context, id uuid.UUID) (*models.ImproveRequestRevision, error) {
	data, err := s.repository.GetRevision(ctx, id)
	if err != nil {
		return nil, goerrors.Join(ErrGetImproveRequestRevision, err)
	}

	return adapters.ImproveRequestRevisionToModel(data), nil
}
