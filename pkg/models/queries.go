package models

import "github.com/a-novel/go-apis"

const (
	OrderScore = "score"
)

type SearchImproveRequestsQuery struct {
	UserID apis.StringUUID `json:"userID" form:"userID"`
	Query  string          `json:"query" form:"query"`
	Order  string          `json:"order" form:"order"`
	Limit  int             `json:"limit" form:"limit"`
	Offset int             `json:"offset" form:"offset"`
}

type SearchImproveSuggestionsQuery struct {
	UserID    apis.StringUUID `json:"userID" form:"userID"`
	SourceID  apis.StringUUID `json:"sourceID" form:"sourceID"`
	RequestID apis.StringUUID `json:"requestID" form:"requestID"`
	Validated *bool           `json:"validated,omitempty" form:"validated,omitempty"`
	Order     string          `json:"order" form:"order"`
	Limit     int             `json:"limit" form:"limit"`
	Offset    int             `json:"offset" form:"offset"`
}

type DeleteImproveRequestQuery struct {
	ID apis.StringUUID `json:"id" form:"id"`
}

type DeleteImproveRequestRevisionQuery struct {
	ID apis.StringUUID `json:"id" form:"id"`
}

type DeleteImproveSuggestionQuery struct {
	ID apis.StringUUID `json:"id" form:"id"`
}

type GetImproveRequestQuery struct {
	ID apis.StringUUID `json:"id" form:"id"`
}

type GetImproveRequestRevisionQuery struct {
	ID apis.StringUUID `json:"id" form:"id"`
}

type ListImproveRequestRevisionsQuery struct {
	ID apis.StringUUID `json:"id" form:"id"`
}

type GetImproveSuggestionQuery struct {
	ID apis.StringUUID `json:"id" form:"id"`
}

type ListImproveRequestsQuery struct {
	IDs apis.StringUUIDs `json:"id" form:"ids"`
}

type ListImproveSuggestionQuery struct {
	IDs apis.StringUUIDs `json:"id" form:"ids"`
}
