package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type ListImproveSuggestionsService interface {
	List(ctx context.Context, ids []uuid.UUID) ([]*models.ImproveSuggestion, error)
}

func NewListImproveSuggestionsService(repository dao.ImproveSuggestionRepository) ListImproveSuggestionsService {
	return &listImproveSuggestionsServiceImpl{
		repository: repository,
	}
}

type listImproveSuggestionsServiceImpl struct {
	repository dao.ImproveSuggestionRepository
}

func (s *listImproveSuggestionsServiceImpl) List(ctx context.Context, ids []uuid.UUID) ([]*models.ImproveSuggestion, error) {
	res, err := s.repository.List(ctx, ids)
	if err != nil {
		return nil, goerrors.Join(ErrListImproveSuggestions, err)
	}

	return lo.Map(res, func(item *dao.ImproveSuggestionModel, _ int) *models.ImproveSuggestion {
		return ParseImproveSuggestion(item)
	}), nil
}
