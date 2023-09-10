package dao

import (
	"context"
	"github.com/a-novel/go-framework/errors"
	"github.com/a-novel/go-framework/postgresql"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"strings"
	"time"
)

type ImproveSuggestionRepository interface {
	// Get returns the improvement suggestion with the given ID.
	Get(ctx context.Context, id uuid.UUID) (*ImproveSuggestionModel, error)
	// Create creates a new improvement suggestion for a given improvement request revision.
	Create(ctx context.Context, data *ImproveSuggestionModelCore, userID, sourceID, id uuid.UUID, now time.Time) (*ImproveSuggestionModel, error)
	// Update updates an existing improvement suggestion.
	Update(ctx context.Context, data *ImproveSuggestionModelCore, id uuid.UUID, now time.Time) (*ImproveSuggestionModel, error)
	// Delete deletes an existing improvement suggestion.
	Delete(ctx context.Context, id uuid.UUID) error

	// Validate validates an existing improvement suggestion.
	Validate(ctx context.Context, validated bool, id uuid.UUID) (*ImproveSuggestionModel, error)
	// UpdateVotes updates the number of up and down votes of a suggestion.
	UpdateVotes(ctx context.Context, id uuid.UUID, upVotes, downVotes int) error

	// Search returns a list of improvement suggestions, matching the provided query. Results must be paginated using
	// the limit and offset parameters.
	// It also returns the total number of available results, to help with pagination.
	Search(ctx context.Context, query ImproveSuggestionSearchQuery, limit, offset int) ([]*ImproveSuggestionModel, int, error)
	List(ctx context.Context, ids []uuid.UUID) ([]*ImproveSuggestionModel, error)
}

type ImproveSuggestionModel struct {
	bun.BaseModel `bun:"table:improve_suggestions"`
	postgresql.Metadata

	// SourceID is the ID of the first revision of the related improvement request. It cannot be changed.
	SourceID uuid.UUID `json:"sourceID" bun:"source_id,type:uuid"`
	// UserID is the ID of the user who created the suggestion.
	UserID uuid.UUID `json:"userID" bun:"user_id,type:uuid"`
	// Validated is true if the suggestion has been validated by the improvement request creator.
	Validated bool `json:"validated" bun:"validated"`

	// UpVotes is the number of up votes the suggestion has received. This value is indirectly updated from the
	// votes table.
	UpVotes int `json:"upVotes" bun:"up_votes"`
	// DownVotes is the number of down votes the suggestion has received. This value is indirectly updated from the
	// votes table.
	DownVotes int `json:"downVotes" bun:"down_votes"`

	ImproveSuggestionModelCore
}

type ImproveSuggestionModelCore struct {
	// RequestID is the ID of the improvement request revision the suggestion is tied to. It must point to a revision
	// of the improvement request with the Model.SourceID.
	RequestID uuid.UUID `json:"requestID" bun:"request_id,type:uuid"`
	// Title an improved version of the source Title. It should match it if no modifications are intended.
	Title string `json:"title" bun:"title"`
	// Content contains the updated content of the source request.
	Content string `json:"content" bun:"content"`
}

type ImproveSuggestionSearchQueryOrder struct {
	Score bool `json:"score"`
}

type ImproveSuggestionSearchQuery struct {
	// UserID is an optional parameter, to only target suggestions that were created by a specific author.
	UserID *uuid.UUID `json:"userID"`
	// SourceID is an optional parameter, to only target suggestions that were created for a specific improvement
	// request.
	SourceID *uuid.UUID `json:"sourceID"`
	// RequestID is an optional parameter, to only target suggestions that were created for a specific improvement
	// request revision.
	RequestID *uuid.UUID `json:"requestID"`
	// Validated is an optional parameter, to only target suggestions that have been validated by the improvement
	// request creator.
	Validated *bool `json:"validated"`
	// Order specifies custom ordering for the search results.
	Order *ImproveSuggestionSearchQueryOrder `json:"order"`
}

type improveSuggestionRepositoryImpl struct {
	db bun.IDB
}

func NewImproveSuggestionRepository(db bun.IDB) ImproveSuggestionRepository {
	return &improveSuggestionRepositoryImpl{
		db: db,
	}
}

func (repository *improveSuggestionRepositoryImpl) Get(ctx context.Context, id uuid.UUID) (*ImproveSuggestionModel, error) {
	suggestion := &ImproveSuggestionModel{Metadata: postgresql.Metadata{ID: id}}
	if err := repository.db.NewSelect().Model(suggestion).WherePK().Scan(ctx); err != nil {
		return nil, errors.HandlePGError(err)
	}

	return suggestion, nil
}

func (repository *improveSuggestionRepositoryImpl) Create(ctx context.Context, data *ImproveSuggestionModelCore, userID, sourceID, id uuid.UUID, now time.Time) (*ImproveSuggestionModel, error) {
	suggestion := &ImproveSuggestionModel{
		Metadata: postgresql.Metadata{
			ID:        id,
			CreatedAt: now,
		},
		SourceID:                   sourceID,
		UserID:                     userID,
		ImproveSuggestionModelCore: *data,
	}

	err := repository.db.NewInsert().Model(suggestion).Returning("*").Scan(ctx)
	if err != nil {
		return nil, errors.HandlePGError(err)
	}

	return suggestion, nil
}

func (repository *improveSuggestionRepositoryImpl) Update(ctx context.Context, data *ImproveSuggestionModelCore, id uuid.UUID, now time.Time) (*ImproveSuggestionModel, error) {
	suggestion := &ImproveSuggestionModel{
		Metadata: postgresql.Metadata{
			ID:        id,
			UpdatedAt: &now,
		},
		ImproveSuggestionModelCore: *data,
	}

	err := repository.db.NewUpdate().Model(suggestion).Column("updated_at", "request_id", "title", "content").WherePK().Returning("*").Scan(ctx)
	if err != nil {
		return nil, errors.HandlePGError(err)
	}

	return suggestion, nil
}

func (repository *improveSuggestionRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	suggestion := &ImproveSuggestionModel{Metadata: postgresql.Metadata{ID: id}}

	_, err := repository.db.NewDelete().Model(suggestion).WherePK().Exec(ctx)
	if err != nil {
		return errors.HandlePGError(err)
	}

	return nil
}

func (repository *improveSuggestionRepositoryImpl) Validate(ctx context.Context, validated bool, id uuid.UUID) (*ImproveSuggestionModel, error) {
	suggestion := &ImproveSuggestionModel{
		Metadata:  postgresql.Metadata{ID: id},
		Validated: validated,
	}

	err := repository.db.NewUpdate().Model(suggestion).Column("validated").WherePK().Returning("*").Scan(ctx)
	if err != nil {
		return nil, errors.HandlePGError(err)
	}

	return suggestion, nil
}

func (repository *improveSuggestionRepositoryImpl) UpdateVotes(ctx context.Context, id uuid.UUID, upVotes, downVotes int) error {
	suggestion := &ImproveSuggestionModel{
		Metadata:  postgresql.Metadata{ID: id},
		UpVotes:   upVotes,
		DownVotes: downVotes,
	}

	rows, err := repository.db.NewUpdate().Model(suggestion).Column("up_votes", "down_votes").WherePK().Exec(ctx)
	if err != nil {
		return errors.HandlePGError(err)
	}

	if err := errors.ForceRowsUpdate(rows); err != nil {
		return err
	}

	return nil
}

func (repository *improveSuggestionRepositoryImpl) Search(ctx context.Context, query ImproveSuggestionSearchQuery, limit, offset int) ([]*ImproveSuggestionModel, int, error) {
	suggestions := make([]*ImproveSuggestionModel, 0)

	queryBuilder := repository.db.NewSelect().Model(&suggestions).Limit(limit).Offset(offset)

	if query.UserID != nil {
		queryBuilder.Where("user_id = ?", *query.UserID)
	}

	if query.SourceID != nil {
		queryBuilder.Where("source_id = ?", *query.SourceID)
	}

	if query.RequestID != nil {
		queryBuilder.Where("request_id = ?", *query.RequestID)
	}

	if query.Validated != nil {
		queryBuilder.Where("validated = ?", *query.Validated)
	}

	var orderBy []string

	if query.Order != nil {
		if query.Order.Score {
			orderBy = append(orderBy, "up_votes - down_votes DESC")
		}
	}

	orderBy = append(orderBy, "updated_at DESC")

	queryBuilder = queryBuilder.OrderExpr(strings.Join(orderBy, ", "))

	count, err := queryBuilder.ScanAndCount(ctx)
	if err != nil {
		return nil, 0, errors.HandlePGError(err)
	}

	return suggestions, count, nil
}

func (repository *improveSuggestionRepositoryImpl) List(ctx context.Context, ids []uuid.UUID) ([]*ImproveSuggestionModel, error) {
	suggestions := make([]*ImproveSuggestionModel, 0)

	err := repository.db.NewSelect().Model(&suggestions).Where("id IN (?)", bun.In(ids)).Scan(ctx)
	if err != nil {
		return nil, errors.HandlePGError(err)
	}

	return suggestions, nil
}
