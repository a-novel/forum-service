package models

import (
	"github.com/google/uuid"
	"time"
)

type ImproveSuggestion struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`

	// SourceID is the ID of the first revision of the related improvement request. It cannot be changed.
	SourceID uuid.UUID `json:"sourceID"`
	// UserID is the ID of the user who created the suggestion.
	UserID uuid.UUID `json:"userID"`
	// Validated is true if the suggestion has been validated by the improvement request creator.
	Validated bool `json:"validated"`

	// UpVotes is the number of up votes the suggestion has received. This value is indirectly updated from the
	// votes table.
	UpVotes int `json:"upVotes"`
	// DownVotes is the number of down votes the suggestion has received. This value is indirectly updated from the
	// votes table.
	DownVotes int `json:"downVotes"`

	// RequestID is the ID of the improvement request revision the suggestion is tied to. It must point to a revision
	// of the improvement request with the Model.SourceID.
	RequestID uuid.UUID `json:"requestID"`
	// Title an improved version of the source Title. It should match it if no modifications are intended.
	Title string `json:"title"`
	// Content contains the updated content of the source request.
	Content string `json:"content"`
}
