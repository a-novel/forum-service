package services

import (
	goerrors "errors"
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
