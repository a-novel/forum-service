package adapters

import (
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

func ImproveRequestSearchQueryToDAO(src models.SearchImproveRequestsQuery) dao.ImproveRequestSearchQuery {
	output := dao.ImproveRequestSearchQuery{
		Query: src.Query,
	}

	if src.UserID.Value() != uuid.Nil {
		output.UserID = lo.ToPtr(src.UserID.Value())
	}

	if src.Order == models.OrderScore {
		output.Order = &dao.ImproveRequestSearchQueryOrder{Score: true}
	}

	return output
}
