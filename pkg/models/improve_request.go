package models

import (
	"github.com/google/uuid"
	"time"
)

type ImproveRequestRevision struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`

	SourceID uuid.UUID `json:"sourceID"`

	// UserID is the ID of the user who created the request, or edited the revision.
	UserID uuid.UUID `json:"userID"`
	// Title is a quick summary of the Content, and the goal it tries to achieve.
	Title string `json:"title"`
	// Content is a novel scene that the user wants to improve.
	Content string `json:"content"`
}

type ImproveRequestRevisionPreview struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`

	// SuggestionsCount returns the total number of suggestions, associated with the request revision.
	SuggestionsCount int `json:"suggestionsCount"`
	// AcceptedSuggestionsCount returns the number of suggestions that have been accepted by the user, on the current
	// revision.
	AcceptedSuggestionsCount int `json:"acceptedSuggestionsCount"`
}

type ImproveRequestPreview struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`

	// UserID is the ID of the user who created the request, or edited the revision.
	UserID uuid.UUID `json:"userID"`
	// Title is a quick summary of the Content, and the goal it tries to achieve.
	Title string `json:"title"`
	// Content is a novel scene that the user wants to improve.
	Content string `json:"content"`

	// UpVotes is the number of up votes the request has received. This value is indirectly updated from the
	// votes table.
	UpVotes int `json:"upVotes"`
	// DownVotes is the number of down votes the request has received. This value is indirectly updated from the
	// votes table.
	DownVotes int `json:"downVotes"`

	// SuggestionsCount returns the total number of suggestions, associated with the request revision.
	SuggestionsCount int `json:"suggestionsCount"`
	// AcceptedSuggestionsCount returns the number of suggestions that have been accepted by the user, on the current
	// revision.
	AcceptedSuggestionsCount int `json:"acceptedSuggestionsCount"`
	// RevisionCount is the number of revisions the request has.
	RevisionCount int `json:"revisionsCount"`
}
