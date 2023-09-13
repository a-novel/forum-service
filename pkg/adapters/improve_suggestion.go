package adapters

import (
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
)

func ImproveSuggestionToModel(src *dao.ImproveSuggestionModel) *models.ImproveSuggestion {
	if src == nil {
		return nil
	}

	return &models.ImproveSuggestion{
		ID:        src.ID,
		CreatedAt: src.CreatedAt,
		UpdatedAt: src.UpdatedAt,
		SourceID:  src.SourceID,
		UserID:    src.UserID,
		Validated: src.Validated,
		UpVotes:   src.UpVotes,
		DownVotes: src.DownVotes,
		RequestID: src.RequestID,
		Title:     src.Title,
		Content:   src.Content,
	}
}
