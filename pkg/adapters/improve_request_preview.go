package adapters

import (
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
)

func ImproveRequestPreviewToModel(src *dao.ImproveRequestPreview) *models.ImproveRequestPreview {
	if src == nil {
		return nil
	}

	return &models.ImproveRequestPreview{
		ID:                       src.ID,
		CreatedAt:                src.CreatedAt,
		UserID:                   src.UserID,
		Title:                    src.Title,
		Content:                  src.Content,
		UpVotes:                  src.UpVotes,
		DownVotes:                src.DownVotes,
		RevisionCount:            src.RevisionCount,
		SuggestionsCount:         src.SuggestionsCount,
		AcceptedSuggestionsCount: src.AcceptedSuggestionsCount,
	}
}
