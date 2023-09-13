package adapters

import (
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
)

func ImproveSuggestionFormToDAO(src *models.ImproveSuggestionForm) *dao.ImproveSuggestionModelCore {
	if src == nil {
		return nil
	}

	return &dao.ImproveSuggestionModelCore{
		RequestID: src.RequestID,
		Title:     src.Title,
		Content:   src.Content,
	}
}
