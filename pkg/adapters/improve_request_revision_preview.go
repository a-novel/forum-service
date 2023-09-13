package adapters

import (
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
)

func ImproveRequestRevisionPreviewToModel(src *dao.ImproveRequestRevisionPreview) *models.ImproveRequestRevisionPreview {
	if src == nil {
		return nil
	}

	return &models.ImproveRequestRevisionPreview{
		ID:                       src.ID,
		CreatedAt:                src.CreatedAt,
		SuggestionsCount:         src.SuggestionsCount,
		AcceptedSuggestionsCount: src.AcceptedSuggestionsCount,
	}
}
