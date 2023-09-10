package services

import (
	goerrors "errors"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/models"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"regexp"
)

var (
	// Just prevents line breaks in title.
	titleRegexp = regexp.MustCompile(`^[^\n\r]+$`)
)

var (
	ErrNotTheCreator = goerrors.New("only the source post creator is allowed to perform this action")
	ErrTheCreator    = goerrors.New("the source post creator is not allowed to perform this action")
	ErrSwitchSource  = goerrors.New("the new improve request id is on a different source than the original one")

	ErrInvalidToken       = goerrors.New("(data) invalid tokenRaw")
	ErrInvalidTitle       = goerrors.New("(data) invalid title")
	ErrInvalidContent     = goerrors.New("(data) invalid content")
	ErrInvalidSearchLimit = goerrors.New("(data) invalid search limit")

	ErrIntrospectToken = goerrors.New("(dep) failed to introspect tokenRaw")

	ErrListImproveRequestRevisions  = goerrors.New("(dao) failed to list improve request revisions")
	ErrGetImproveRequestRevision    = goerrors.New("(dao) failed to get improve request revision")
	ErrCreateImproveRequest         = goerrors.New("(dao) failed to create improve request")
	ErrDeleteImproveRequest         = goerrors.New("(dao) failed to delete improve request")
	ErrListImproveRequests          = goerrors.New("(dao) failed to list improve requests")
	ErrSearchImproveRequests        = goerrors.New("(dao) failed to search improve requests")
	ErrUpdateImproveRequestRevision = goerrors.New("(dao) failed to update improve request revision")
	ErrGetImproveSuggestion         = goerrors.New("(dao) failed to get improve suggestion")
	ErrCreateImproveSuggestion      = goerrors.New("(dao) failed to create improve suggestions")
	ErrUpdateImproveSuggestion      = goerrors.New("(dao) failed to update improve suggestions")
	ErrDeleteImproveSuggestion      = goerrors.New("(dao) failed to delete improve suggestions")
	ErrSearchImproveSuggestions     = goerrors.New("(dao) failed to search improve suggestions")
	ErrListImproveSuggestions       = goerrors.New("(dao) failed to list improve suggestions")
	ErrValidateImproveSuggestion    = goerrors.New("(dao) failed to validate improve suggestions")
	ErrGetImproveRequest            = goerrors.New("(dao) failed to get improve request")
	ErrDeleteImproveRequestRevision = goerrors.New("(dao) failed to delete improve request revision")
)

const (
	MinTitleLength   = 4
	MaxTitleLength   = 128
	MinContentLength = 4
	MaxContentLength = 4096

	MaxSearchLimit = 100
)

func ParseImproveRequestRevision(src *dao.ImproveRequestRevisionModel) *models.ImproveRequestRevision {
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

func ParseImproveRequestRevisionPreview(src *dao.ImproveRequestRevisionPreview) *models.ImproveRequestRevisionPreview {
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

func ParseImproveRequestPreview(src *dao.ImproveRequestPreview) *models.ImproveRequestPreview {
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

func ParseImproveRequestSearchQuery(src models.SearchImproveRequestsQuery) dao.ImproveRequestSearchQuery {
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

func ParseImproveSuggestion(src *dao.ImproveSuggestionModel) *models.ImproveSuggestion {
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

func ParseImproveSuggestionForm(src *models.ImproveSuggestionForm) *dao.ImproveSuggestionModelCore {
	if src == nil {
		return nil
	}

	return &dao.ImproveSuggestionModelCore{
		RequestID: src.RequestID,
		Title:     src.Title,
		Content:   src.Content,
	}
}

func ParseImproveSuggestionSearchQuery(src models.SearchImproveSuggestionsQuery) dao.ImproveSuggestionSearchQuery {
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
