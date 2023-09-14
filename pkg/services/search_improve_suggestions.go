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
	if err := goframework.CheckMinMax(query.Limit, 1, MaxSearchLimit); err != nil {
		return nil, 0, goerrors.Join(goframework.ErrInvalidEntity, ErrInvalidSearchLimit, err)
	}

	res, total, err := s.repository.Search(ctx, adapters.ImproveSuggestionSearchQueryToDAO(query), query.Limit, query.Offset)
	if err != nil {
		return nil, 0, goerrors.Join(ErrSearchImproveSuggestions, err)
	}

	return lo.Map(res, func(item *dao.ImproveSuggestionModel, _ int) *models.ImproveSuggestion {
		return adapters.ImproveSuggestionToModel(item)
	}), total, nil
}
