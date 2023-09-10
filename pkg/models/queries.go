package models

import (
	"github.com/a-novel/go-framework/types"
)

const (
	OrderScore = "score"
)

type SearchImproveRequestsQuery struct {
	UserID types.StringUUID `json:"userID" form:"userID"`
	Query  string           `json:"query" form:"query"`
	Order  string           `json:"order" form:"order"`
	Limit  int              `json:"limit" form:"limit"`
	Offset int              `json:"offset" form:"offset"`
}

type SearchImproveSuggestionsQuery struct {
	UserID    types.StringUUID `json:"userID" form:"userID"`
	SourceID  types.StringUUID `json:"sourceID" form:"sourceID"`
	RequestID types.StringUUID `json:"requestID" form:"requestID"`
	Validated *bool            `json:"validated,omitempty" form:"validated,omitempty"`
	Order     string           `json:"order" form:"order"`
	Limit     int              `json:"limit" form:"limit"`
	Offset    int              `json:"offset" form:"offset"`
}

type DeleteImproveRequestQuery struct {
	ID types.StringUUID `json:"id" form:"id"`
}

type DeleteImproveRequestRevisionQuery struct {
	ID types.StringUUID `json:"id" form:"id"`
}

type DeleteImproveSuggestionQuery struct {
	ID types.StringUUID `json:"id" form:"id"`
}

type GetImproveRequestQuery struct {
	ID types.StringUUID `json:"id" form:"id"`
}

type GetImproveRequestRevisionQuery struct {
	ID types.StringUUID `json:"id" form:"id"`
}

type ListImproveRequestRevisionsQuery struct {
	ID types.StringUUID `json:"id" form:"id"`
}

type GetImproveSuggestionQuery struct {
	ID types.StringUUID `json:"id" form:"id"`
}

type ListImproveRequestsQuery struct {
	IDs types.StringUUIDs `json:"id" form:"ids"`
}

type ListImproveSuggestionQuery struct {
	IDs types.StringUUIDs `json:"id" form:"ids"`
}
