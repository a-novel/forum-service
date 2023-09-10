package models

import "github.com/google/uuid"

type CreateImproveRequestForm struct {
	Title    string    `json:"title" form:"title"`
	Content  string    `json:"content" form:"content"`
	SourceID uuid.UUID `json:"sourceID" form:"sourceID"`
}

type ImproveSuggestionForm struct {
	RequestID uuid.UUID `json:"requestID" form:"requestID"`
	Title     string    `json:"title" form:"title"`
	Content   string    `json:"content" form:"content"`
}

type ValidateImproveSuggestionForm struct {
	ID        uuid.UUID `json:"id" form:"id"`
	Validated bool      `json:"validated" form:"validated"`
}

type UpdateImproveRequestVotesForm struct {
	ID        uuid.UUID `json:"id" form:"id"`
	UserID    uuid.UUID `json:"userID" form:"userID"`
	UpVotes   int       `json:"upVotes" form:"upVotes"`
	DownVotes int       `json:"downVotes" form:"downVotes"`
}

type UpdateImproveSuggestionVotesForm struct {
	ID        uuid.UUID `json:"id" form:"id"`
	UserID    uuid.UUID `json:"userID" form:"userID"`
	UpVotes   int       `json:"upVotes" form:"upVotes"`
	DownVotes int       `json:"downVotes" form:"downVotes"`
}
