package adapters

import (
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

func ImproveSuggestionSearchQueryToDAO(src models.SearchImproveSuggestionsQuery) dao.ImproveSuggestionSearchQuery {
	output := dao.ImproveSuggestionSearchQuery{
		Validated: src.Validated,
	}

	if src.UserID.Value() != uuid.Nil {
		output.UserID = lo.ToPtr(src.UserID.Value())
	}

	if src.SourceID.Value() != uuid.Nil {
		output.SourceID = lo.ToPtr(src.SourceID.Value())
	}

	if src.RequestID.Value() != uuid.Nil {
		output.RequestID = lo.ToPtr(src.RequestID.Value())
	}

	if src.Order == models.OrderScore {
		output.Order = &dao.ImproveSuggestionSearchQueryOrder{Score: true}
	}

	return output
}
