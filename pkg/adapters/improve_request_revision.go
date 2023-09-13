package adapters

import (
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
)

func ImproveRequestRevisionToModel(src *dao.ImproveRequestRevisionModel) *models.ImproveRequestRevision {
	if src == nil {
		return nil
	}

	return &models.ImproveRequestRevision{
		ID:        src.ID,
		CreatedAt: src.CreatedAt,
		SourceID:  src.SourceID,
		UserID:    src.UserID,
		Title:     src.Title,
		Content:   src.Content,
	}
}
