package dao

import (
	"context"
	"fmt"
	"github.com/a-novel/bunovel"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"strings"
	"time"
)

type ImproveRequestRepository interface {
	GetRevision(ctx context.Context, id uuid.UUID) (*ImproveRequestRevisionModel, error)
	Get(ctx context.Context, id uuid.UUID) (*ImproveRequestPreview, error)
	ListRevisions(ctx context.Context, id uuid.UUID) ([]*ImproveRequestRevisionPreview, error)
	UpdateVotes(ctx context.Context, id uuid.UUID, upVotes, downVotes int) error
	Create(ctx context.Context, userID uuid.UUID, title, content string, sourceID, id uuid.UUID, now time.Time) (*ImproveRequestPreview, error)
	DeleteRevision(ctx context.Context, id uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
	Search(ctx context.Context, query ImproveRequestSearchQuery, limit, offset int) ([]*ImproveRequestPreview, int, error)
	List(ctx context.Context, ids []uuid.UUID) ([]*ImproveRequestPreview, error)
}

type ImproveRequestModel struct {
	bun.BaseModel `bun:"table:improve_requests"`
	bunovel.Metadata

	// UpVotes is the number of up votes the request has received. This value is indirectly updated from the
	// votes table.
	UpVotes int `bun:"up_votes"`
	// DownVotes is the number of down votes the request has received. This value is indirectly updated from the
	// votes table.
	DownVotes int `bun:"down_votes"`
}

type ImproveRequestRevisionModel struct {
	bun.BaseModel `bun:"table:improve_requests_revisions"`
	bunovel.Metadata

	// SourceID points to the first revision. It is equal to the ID if no other revision exist.
	// An improvement request is never updated. Instead, new revisions are created every time.
	SourceID uuid.UUID `bun:"source_id,type:uuid"`

	// UserID is the ID of the user who created the request, or edited the revision.
	UserID uuid.UUID `bun:"user_id,type:uuid"`
	// Title is a quick summary of the Content, and the goal it tries to achieve.
	Title string `bun:"title"`
	// Content is a novel scene that the user wants to improve.
	Content string `bun:"content"`
}

type ImproveRequestRevisionPreview struct {
	bun.BaseModel `bun:"table:improve_requests_revisions_list"`
	bunovel.Metadata

	SuggestionsCount         int `bun:"suggestions_count"`
	AcceptedSuggestionsCount int `bun:"accepted_suggestions_count"`
}

type ImproveRequestPreview struct {
	bun.BaseModel `bun:"table:improve_requests_previews"`
	bunovel.Metadata

	// UserID is the ID of the user who created the request, or edited the revision.
	UserID uuid.UUID `bun:"user_id,type:uuid"`
	// Title is a quick summary of the Content, and the goal it tries to achieve.
	Title string `bun:"title"`
	// Content is a novel scene that the user wants to improve.
	Content string `bun:"content"`

	// UpVotes is the number of up votes the request and all its revisions has received. This value is indirectly
	// updated from the votes table.
	UpVotes int `bun:"up_votes"`
	// DownVotes is the number of down votes the request and all its revisions has received. This value is indirectly
	// updated from the votes table.
	DownVotes int `bun:"down_votes"`

	// RevisionCount is the number of revisions the request has.
	RevisionCount int `bun:"revisions_count"`
	// SuggestionsCount returns the total number of suggestions, associated with the request and all its revisions.
	SuggestionsCount int `bun:"suggestions_count"`
	// AcceptedSuggestionsCount returns the total number of accepted suggestions, associated with the request and all
	// its revisions.
	AcceptedSuggestionsCount int `bun:"accepted_suggestions_count"`
}

type ImproveRequestSearchQueryOrder struct {
	Score bool
}

// ImproveRequestSearchQuery allows to filter improve requests.
type ImproveRequestSearchQuery struct {
	// UserID is an optional parameter, to only target requests that were created/revised by a specific author.
	UserID *uuid.UUID
	// Query is an optional parameter, to filter requests based on their title or content.
	Query string
	// Order specifies custom ordering for the search results.
	Order *ImproveRequestSearchQueryOrder
}

type improveRequestRepositoryImpl struct {
	db bun.IDB
}

func NewImproveRequestRepository(db bun.IDB) ImproveRequestRepository {
	return &improveRequestRepositoryImpl{db: db}
}

func (repository *improveRequestRepositoryImpl) GetRevision(ctx context.Context, id uuid.UUID) (*ImproveRequestRevisionModel, error) {
	model := &ImproveRequestRevisionModel{
		Metadata: bunovel.Metadata{ID: id},
	}

	if err := repository.db.NewSelect().Model(model).WherePK().Scan(ctx); err != nil {
		return nil, bunovel.HandlePGError(err)
	}

	return model, nil
}

func (repository *improveRequestRepositoryImpl) Get(ctx context.Context, id uuid.UUID) (*ImproveRequestPreview, error) {
	model := &ImproveRequestPreview{
		Metadata: bunovel.Metadata{ID: id},
	}

	if err := repository.db.NewSelect().Model(model).WherePK().Scan(ctx); err != nil {
		return nil, bunovel.HandlePGError(err)
	}

	return model, nil
}

func (repository *improveRequestRepositoryImpl) ListRevisions(ctx context.Context, id uuid.UUID) ([]*ImproveRequestRevisionPreview, error) {
	models := make([]*ImproveRequestRevisionPreview, 0)

	if err := repository.db.NewSelect().Model(&models).Where("source_id = ?", id).Order("created_at DESC").Scan(ctx); err != nil {
		return nil, bunovel.HandlePGError(err)
	}

	if len(models) == 0 {
		return nil, bunovel.ErrNotFound
	}

	return models, nil
}

func (repository *improveRequestRepositoryImpl) UpdateVotes(ctx context.Context, id uuid.UUID, upVotes, downVotes int) error {
	model := &ImproveRequestModel{Metadata: bunovel.Metadata{ID: id}, UpVotes: upVotes, DownVotes: downVotes}

	rows, err := repository.db.NewUpdate().
		Model(model).
		Column("up_votes", "down_votes").
		WherePK().
		Exec(ctx)
	if err != nil {
		return bunovel.HandlePGError(err)
	}

	if err := bunovel.ForceRowsUpdate(rows); err != nil {
		return err
	}

	return nil
}

func (repository *improveRequestRepositoryImpl) Create(ctx context.Context, userID uuid.UUID, title, content string, sourceID, id uuid.UUID, now time.Time) (*ImproveRequestPreview, error) {
	output := new(ImproveRequestPreview)

	if err := repository.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		model := &ImproveRequestModel{
			Metadata: bunovel.NewMetadata(sourceID, now, nil),
		}

		exists, err := repository.db.NewSelect().Model(model).WherePK().Exists(ctx)
		if err != nil {
			return fmt.Errorf("failed to check if improve request exists: %w", err)
		}

		if !exists {
			if err := repository.db.NewInsert().Model(model).Scan(ctx); err != nil {
				return fmt.Errorf("failed to create improve request: %w", err)
			}
		}

		revisionModel := &ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(id, now, nil),
			SourceID: sourceID,
			UserID:   userID,
			Title:    title,
			Content:  content,
		}

		if err := repository.db.NewInsert().Model(revisionModel).Scan(ctx); err != nil {
			return fmt.Errorf("failed to create improve request revision: %w", err)
		}

		output.UserID = userID
		output.Title = title
		output.Content = content
		output.Metadata = bunovel.Metadata{ID: sourceID, CreatedAt: now}

		return nil
	}); err != nil {
		return nil, bunovel.HandlePGError(err)
	}

	return output, nil
}

func (repository *improveRequestRepositoryImpl) DeleteRevision(ctx context.Context, id uuid.UUID) error {
	model := &ImproveRequestRevisionModel{Metadata: bunovel.Metadata{ID: id}}
	if _, err := repository.db.NewDelete().Model(model).WherePK().Exec(ctx); err != nil {
		return bunovel.HandlePGError(err)
	}

	return nil
}

func (repository *improveRequestRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	if err := repository.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		model := &ImproveRequestModel{Metadata: bunovel.Metadata{ID: id}}
		if _, err := repository.db.NewDelete().Model(model).WherePK().Exec(ctx); err != nil {
			return fmt.Errorf("failed to delete improve request: %w", err)
		}

		revisionModel := new(ImproveRequestRevisionModel)
		if _, err := repository.db.NewDelete().Model(revisionModel).Where("source_id = ?", id).Exec(ctx); err != nil {
			return fmt.Errorf("failed to delete improve request revisions: %w", err)
		}

		return nil
	}); err != nil {
		return bunovel.HandlePGError(err)
	}

	return nil
}

func (repository *improveRequestRepositoryImpl) Search(ctx context.Context, query ImproveRequestSearchQuery, limit, offset int) ([]*ImproveRequestPreview, int, error) {
	model := make([]*ImproveRequestPreview, 0)

	queryBuilder := repository.db.NewSelect().Model(&model).Limit(limit).Offset(offset)

	if query.UserID != nil {
		queryBuilder.Where("user_id = ?", query.UserID)
	}

	var orderBy []string

	if query.Query != "" {
		queryFullText := repository.db.NewSelect().
			ColumnExpr("to_tsquery('french', string_agg(lexeme || ':*', ' & ' order by positions)) AS query").
			TableExpr("unnest(to_tsvector('french', unaccent(?)))", query.Query)

		queryBuilder = queryBuilder.
			TableExpr("(?) AS search", queryFullText).
			Where("text_searchable_index_col @@ search.query")

		orderBy = append(orderBy, "ts_rank_cd(text_searchable_index_col, search.query) DESC")
	}

	if query.Order != nil {
		if query.Order.Score {
			orderBy = append(orderBy, "up_votes - down_votes DESC")
		}
	}

	orderBy = append(orderBy, "created_at DESC")

	queryBuilder = queryBuilder.OrderExpr(strings.Join(orderBy, ", "))

	count, err := queryBuilder.ScanAndCount(ctx)
	if err != nil {
		return nil, 0, bunovel.HandlePGError(err)
	}

	return model, count, nil
}

func (repository *improveRequestRepositoryImpl) List(ctx context.Context, ids []uuid.UUID) ([]*ImproveRequestPreview, error) {
	model := make([]*ImproveRequestPreview, 0)

	if err := repository.db.NewSelect().Model(&model).Where("id IN (?)", bun.In(ids)).Scan(ctx); err != nil {
		return nil, bunovel.HandlePGError(err)
	}

	return model, nil
}
