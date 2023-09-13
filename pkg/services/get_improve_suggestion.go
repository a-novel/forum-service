package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/adapters"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/google/uuid"
)

type GetImproveSuggestionService interface {
	Get(ctx context.Context, id uuid.UUID) (*models.ImproveSuggestion, error)
}

func NewGetImproveSuggestionService(service dao.ImproveSuggestionRepository) GetImproveSuggestionService {
	return &getImproveSuggestionServiceImpl{
		service: service,
	}
}

type getImproveSuggestionServiceImpl struct {
	service dao.ImproveSuggestionRepository
}

func (s *getImproveSuggestionServiceImpl) Get(ctx context.Context, id uuid.UUID) (*models.ImproveSuggestion, error) {
	suggestion, err := s.service.Get(ctx, id)
	if err != nil {
		return nil, goerrors.Join(ErrGetImproveSuggestion, err)
	}

	return adapters.ImproveSuggestionToModel(suggestion), nil
}
