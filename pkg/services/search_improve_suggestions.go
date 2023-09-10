package services

import (
	"context"
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/a-novel/go-framework/errors"
	"github.com/a-novel/go-framework/validation"
	"github.com/samber/lo"
)

type SearchImproveSuggestionsService interface {
	Search(ctx context.Context, query models.SearchImproveSuggestionsQuery) ([]*models.ImproveSuggestion, int, error)
}

func NewSearchImproveSuggestionsService(repository dao.ImproveSuggestionRepository) SearchImproveSuggestionsService {
	return &searchImproveSuggestionsServiceImpl{
		repository: repository,
	}
}

type searchImproveSuggestionsServiceImpl struct {
	repository dao.ImproveSuggestionRepository
}

func (s *searchImproveSuggestionsServiceImpl) Search(ctx context.Context, query models.SearchImproveSuggestionsQuery) ([]*models.ImproveSuggestion, int, error) {
	if err := validation.CheckMinMax(query.Limit, 1, MaxSearchLimit); err != nil {
		return nil, 0, goerrors.Join(errors.ErrInvalidEntity, ErrInvalidSearchLimit, err)
	}

	res, total, err := s.repository.Search(ctx, ParseImproveSuggestionSearchQuery(query), query.Limit, query.Offset)
	if err != nil {
		return nil, 0, goerrors.Join(ErrSearchImproveSuggestions, err)
	}

	return lo.Map(res, func(item *dao.ImproveSuggestionModel, _ int) *models.ImproveSuggestion {
		return ParseImproveSuggestion(item)
	}), total, nil
}
