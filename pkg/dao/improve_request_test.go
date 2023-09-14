package dao_test

import (
	"context"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/migrations"
	"github.com/a-novel/forum-service/pkg/dao"
	goframework "github.com/a-novel/go-framework"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"io/fs"
	"testing"
	"time"
)

func TestImproveRequestRepository_GetRevision(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
			UpVotes:   160,
			DownVotes: 80,
		},

		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with robots",
			Content:  "my content with mechanics",
		},
		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(2), baseTime.Add(time.Hour), &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with spaceships",
			Content:  "my content with thrusters",
		},
	}

	data := []struct {
		name string

		id uuid.UUID

		expect    *dao.ImproveRequestRevisionModel
		expectErr error
	}{
		{
			name: "Success",
			id:   goframework.NumberUUID(1),
			expect: &dao.ImproveRequestRevisionModel{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
				SourceID: goframework.NumberUUID(10),
				UserID:   goframework.NumberUUID(100),
				Title:    "my title with robots",
				Content:  "my content with mechanics",
			},
		},
		{
			name:      "Error/NotFound",
			id:        goframework.NumberUUID(3),
			expectErr: bunovel.ErrNotFound,
		},
	}

	err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
		repository := dao.NewImproveRequestRepository(tx)

		for _, d := range data {
			t.Run(d.name, func(st *testing.T) {
				res, err := repository.GetRevision(ctx, d.id)
				require.ErrorIs(t, err, d.expectErr)
				require.Equal(t, d.expect, res)
			})
		}
	})
	require.NoError(t, err)
}

func TestImproveRequestRepository_Get(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
			UpVotes:   160,
			DownVotes: 80,
		},

		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with robots",
			Content:  "my content with mechanics",
		},
		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(2), baseTime.Add(time.Hour), &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with spaceships",
			Content:  "my content with thrusters",
		},

		// Suggestions for goframework.NumberUUID(1).
		&dao.ImproveSuggestionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   uuid.Nil,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},

		// Suggestions for goframework.NumberUUID(2).
		&dao.ImproveSuggestionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(4), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   uuid.Nil,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(5), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
	}

	data := []struct {
		name string

		id uuid.UUID

		expect    *dao.ImproveRequestPreview
		expectErr error
	}{
		{
			name: "Success",
			id:   goframework.NumberUUID(10),
			expect: &dao.ImproveRequestPreview{
				Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
				UserID:                   goframework.NumberUUID(100),
				Title:                    "my title with spaceships",
				Content:                  "my content with thrusters",
				UpVotes:                  160,
				DownVotes:                80,
				RevisionCount:            2,
				SuggestionsCount:         5,
				AcceptedSuggestionsCount: 3,
			},
		},
		{
			name:      "Error/NotFound",
			id:        goframework.NumberUUID(20),
			expectErr: bunovel.ErrNotFound,
		},
	}

	err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
		repository := dao.NewImproveRequestRepository(tx)

		for _, d := range data {
			t.Run(d.name, func(st *testing.T) {
				res, err := repository.Get(ctx, d.id)
				require.ErrorIs(t, err, d.expectErr)
				require.Equal(t, d.expect, res)
			})
		}
	})
	require.NoError(t, err)
}

func TestImproveRequestRepository_ListRevisions(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
			UpVotes:   160,
			DownVotes: 80,
		},

		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with robots",
			Content:  "my content with mechanics",
		},
		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(2), baseTime.Add(time.Hour), &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with spaceships",
			Content:  "my content with thrusters",
		},

		// Suggestions for goframework.NumberUUID(1).
		&dao.ImproveSuggestionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   uuid.Nil,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},

		// Suggestions for goframework.NumberUUID(2).
		&dao.ImproveSuggestionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(4), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   uuid.Nil,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(5), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
	}

	data := []struct {
		name string

		id uuid.UUID

		expect    []*dao.ImproveRequestRevisionPreview
		expectErr error
	}{
		{
			name: "Success",
			id:   goframework.NumberUUID(10),
			expect: []*dao.ImproveRequestRevisionPreview{
				{
					Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(2), baseTime.Add(time.Hour), &updateTime),
					SuggestionsCount:         2,
					AcceptedSuggestionsCount: 1,
				},
				{
					Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
					SuggestionsCount:         3,
					AcceptedSuggestionsCount: 2,
				},
			},
		},
		{
			name:      "Error/NotFound",
			id:        goframework.NumberUUID(3),
			expectErr: bunovel.ErrNotFound,
		},
	}

	err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
		repository := dao.NewImproveRequestRepository(tx)

		for _, d := range data {
			t.Run(d.name, func(st *testing.T) {
				res, err := repository.ListRevisions(ctx, d.id)
				require.ErrorIs(t, err, d.expectErr)
				require.EqualValues(t, d.expect, res)
			})
		}
	})
	require.NoError(t, err)
}

func TestImproveRequestRepository_UpdateVotes(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
			UpVotes:   160,
			DownVotes: 80,
		},
		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with robots",
			Content:  "my content with mechanics",
		},
	}

	data := []struct {
		name string

		id        uuid.UUID
		upVotes   int
		downVotes int

		expectErr error
	}{
		{
			name:      "Success",
			id:        goframework.NumberUUID(10),
			upVotes:   256,
			downVotes: 512,
		},
		{
			name:      "Error/NotFound",
			id:        goframework.NumberUUID(3),
			expectErr: bunovel.ErrNotFound,
			upVotes:   256,
			downVotes: 512,
		},
	}

	for _, d := range data {
		err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
			repository := dao.NewImproveRequestRepository(tx)

			t.Run(d.name, func(st *testing.T) {
				err := repository.UpdateVotes(ctx, d.id, d.upVotes, d.downVotes)
				require.ErrorIs(t, err, d.expectErr)
			})
		})
		require.NoError(t, err)
	}
}

func TestImproveRequestRepository_Create(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
			UpVotes:   160,
			DownVotes: 80,
		},
		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with robots",
			Content:  "my content with mechanics",
		},
	}

	data := []struct {
		name string

		userID   uuid.UUID
		title    string
		content  string
		sourceID uuid.UUID
		id       uuid.UUID
		now      time.Time

		expect    *dao.ImproveRequestPreview
		expectErr error
	}{
		{
			name:     "Success",
			userID:   goframework.NumberUUID(200),
			title:    "my title",
			content:  "my content",
			sourceID: goframework.NumberUUID(20),
			id:       goframework.NumberUUID(2),
			now:      baseTime,
			expect: &dao.ImproveRequestPreview{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(20), baseTime, nil),
				UserID:   goframework.NumberUUID(200),
				Title:    "my title",
				Content:  "my content",
			},
		},
		{
			name:     "Success/Revision",
			userID:   goframework.NumberUUID(200),
			title:    "my title",
			content:  "my content",
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(2),
			now:      baseTime,
			expect: &dao.ImproveRequestPreview{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, nil),
				UserID:   goframework.NumberUUID(200),
				Title:    "my title",
				Content:  "my content",
			},
		},
	}

	for _, d := range data {
		err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
			repository := dao.NewImproveRequestRepository(tx)
			t.Run(d.name, func(st *testing.T) {
				res, err := repository.Create(ctx, d.userID, d.title, d.content, d.sourceID, d.id, d.now)
				require.Equal(t, d.expect, res)
				require.ErrorIs(t, err, d.expectErr)
			})
		})
		require.NoError(t, err)
	}
}

func TestImproveRequestRepository_Delete(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
			UpVotes:   160,
			DownVotes: 80,
		},
		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with robots",
			Content:  "my content with mechanics",
		},
	}

	data := []struct {
		name string

		id uuid.UUID

		expectErr error
	}{
		{
			name: "Success",
			id:   goframework.NumberUUID(10),
		},
		{
			name: "Success/NotFound",
			id:   goframework.NumberUUID(20),
		},
	}

	for _, d := range data {
		err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
			repository := dao.NewImproveRequestRepository(tx)
			t.Run(d.name, func(st *testing.T) {
				err := repository.Delete(ctx, d.id)
				require.ErrorIs(t, err, d.expectErr)
			})
		})
		require.NoError(t, err)
	}
}

func TestImproveRequestRepository_DeleteRevision(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
			UpVotes:   160,
			DownVotes: 80,
		},
		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with robots",
			Content:  "my content with mechanics",
		},
	}

	data := []struct {
		name string

		id uuid.UUID

		expectErr error
	}{
		{
			name: "Success",
			id:   goframework.NumberUUID(1),
		},
		{
			name: "Success/NotFound",
			id:   goframework.NumberUUID(2),
		},
	}

	for _, d := range data {
		err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
			repository := dao.NewImproveRequestRepository(tx)
			t.Run(d.name, func(st *testing.T) {
				err := repository.DeleteRevision(ctx, d.id)
				require.ErrorIs(t, err, d.expectErr)
			})
		})
		require.NoError(t, err)
	}
}

func TestImproveRequestRepository_Search(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
			UpVotes:   160,
			DownVotes: 80,
		},

		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with robots",
			Content:  "my content with mechanics",
		},
		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(2), baseTime.Add(time.Hour), &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with spaceships",
			Content:  "my content with thrusters",
		},

		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(20), baseTime.Add(2*time.Hour), &updateTime),
			UpVotes:   128,
			DownVotes: 64,
		},

		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(3), baseTime.Add(2*time.Hour), &updateTime),
			SourceID: goframework.NumberUUID(20),
			UserID:   goframework.NumberUUID(200),
			Title:    "my title with thrusters",
			Content:  "my content with spaceships",
		},

		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(30), baseTime.Add(3*time.Hour), &updateTime),
			UpVotes:   128,
			DownVotes: 64,
		},

		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(4), baseTime.Add(3*time.Hour), &updateTime),
			SourceID: goframework.NumberUUID(30),
			UserID:   goframework.NumberUUID(300),
			Title:    "my title with super thrusters",
			Content:  "my content with super spaceships",
		},

		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(40), baseTime.Add(4*time.Hour), &updateTime),
			UpVotes:   128,
			DownVotes: 64,
		},

		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(5), baseTime.Add(4*time.Hour), &updateTime),
			SourceID: goframework.NumberUUID(40),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with tomatoes",
			Content:  "my content with super chips",
		},

		// Suggestions for goframework.NumberUUID(1).
		&dao.ImproveSuggestionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   uuid.Nil,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},

		// Suggestions for goframework.NumberUUID(2).
		&dao.ImproveSuggestionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(4), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   uuid.Nil,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(5), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
	}

	data := []struct {
		name string

		query  dao.ImproveRequestSearchQuery
		limit  int
		offset int

		expect      []*dao.ImproveRequestPreview
		expectCount int
		expectErr   error
	}{
		{
			name: "Success",
			expect: []*dao.ImproveRequestPreview{
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(40), baseTime.Add(4*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(100),
					Title:         "my title with tomatoes",
					Content:       "my content with super chips",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(30), baseTime.Add(3*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(300),
					Title:         "my title with super thrusters",
					Content:       "my content with super spaceships",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(20), baseTime.Add(2*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(200),
					Title:         "my title with thrusters",
					Content:       "my content with spaceships",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
				{
					Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
					UserID:                   goframework.NumberUUID(100),
					Title:                    "my title with spaceships",
					Content:                  "my content with thrusters",
					UpVotes:                  160,
					DownVotes:                80,
					RevisionCount:            2,
					SuggestionsCount:         5,
					AcceptedSuggestionsCount: 3,
				},
			},
			expectCount: 4,
		},
		{
			name: "Success/WithQuery",
			query: dao.ImproveRequestSearchQuery{
				Query: "spaceships",
			},
			expect: []*dao.ImproveRequestPreview{
				// Most relevant, presence in title has more value.
				{
					Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
					UserID:                   goframework.NumberUUID(100),
					Title:                    "my title with spaceships",
					Content:                  "my content with thrusters",
					UpVotes:                  160,
					DownVotes:                80,
					RevisionCount:            2,
					SuggestionsCount:         5,
					AcceptedSuggestionsCount: 3,
				},
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(30), baseTime.Add(3*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(300),
					Title:         "my title with super thrusters",
					Content:       "my content with super spaceships",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(20), baseTime.Add(2*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(200),
					Title:         "my title with thrusters",
					Content:       "my content with spaceships",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
			},
			expectCount: 3,
		},
		{
			name: "Success/WithUserID",
			query: dao.ImproveRequestSearchQuery{
				UserID: lo.ToPtr(goframework.NumberUUID(100)),
			},
			expect: []*dao.ImproveRequestPreview{
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(40), baseTime.Add(4*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(100),
					Title:         "my title with tomatoes",
					Content:       "my content with super chips",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
				{
					Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
					UserID:                   goframework.NumberUUID(100),
					Title:                    "my title with spaceships",
					Content:                  "my content with thrusters",
					UpVotes:                  160,
					DownVotes:                80,
					RevisionCount:            2,
					SuggestionsCount:         5,
					AcceptedSuggestionsCount: 3,
				},
			},
			expectCount: 2,
		},
		{
			name: "Success/WithOrderByScore",
			query: dao.ImproveRequestSearchQuery{
				Order: &dao.ImproveRequestSearchQueryOrder{
					Score: true,
				},
			},
			expect: []*dao.ImproveRequestPreview{
				{
					Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
					UserID:                   goframework.NumberUUID(100),
					Title:                    "my title with spaceships",
					Content:                  "my content with thrusters",
					UpVotes:                  160,
					DownVotes:                80,
					RevisionCount:            2,
					SuggestionsCount:         5,
					AcceptedSuggestionsCount: 3,
				},
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(40), baseTime.Add(4*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(100),
					Title:         "my title with tomatoes",
					Content:       "my content with super chips",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(30), baseTime.Add(3*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(300),
					Title:         "my title with super thrusters",
					Content:       "my content with super spaceships",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(20), baseTime.Add(2*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(200),
					Title:         "my title with thrusters",
					Content:       "my content with spaceships",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
			},
			expectCount: 4,
		},
		{
			name:  "Success/Limit",
			limit: 2,
			expect: []*dao.ImproveRequestPreview{
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(40), baseTime.Add(4*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(100),
					Title:         "my title with tomatoes",
					Content:       "my content with super chips",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(30), baseTime.Add(3*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(300),
					Title:         "my title with super thrusters",
					Content:       "my content with super spaceships",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
			},
			expectCount: 4,
		},
		{
			name:   "Success/Offset",
			offset: 1,
			limit:  2,
			expect: []*dao.ImproveRequestPreview{
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(30), baseTime.Add(3*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(300),
					Title:         "my title with super thrusters",
					Content:       "my content with super spaceships",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(20), baseTime.Add(2*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(200),
					Title:         "my title with thrusters",
					Content:       "my content with spaceships",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
			},
			expectCount: 4,
		},
		{
			name:        "Success/OffsetTooHigh",
			offset:      10,
			expect:      []*dao.ImproveRequestPreview{},
			expectCount: 4,
		},
		{
			name: "Success/NoResult",
			query: dao.ImproveRequestSearchQuery{
				Query: "foo bar qux",
			},
			expect: []*dao.ImproveRequestPreview{},
		},
	}

	err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
		repository := dao.NewImproveRequestRepository(tx)

		for _, d := range data {
			t.Run(d.name, func(st *testing.T) {
				res, count, err := repository.Search(ctx, d.query, d.limit, d.offset)
				require.ErrorIs(t, err, d.expectErr)
				require.Equal(t, d.expect, res)
				require.Equal(t, d.expectCount, count)
			})
		}
	})
	require.NoError(t, err)
}

func TestImproveRequestRepository_List(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
			UpVotes:   160,
			DownVotes: 80,
		},

		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with robots",
			Content:  "my content with mechanics",
		},
		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(2), baseTime.Add(time.Hour), &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with spaceships",
			Content:  "my content with thrusters",
		},

		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(20), baseTime.Add(2*time.Hour), &updateTime),
			UpVotes:   128,
			DownVotes: 64,
		},

		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(3), baseTime.Add(2*time.Hour), &updateTime),
			SourceID: goframework.NumberUUID(20),
			UserID:   goframework.NumberUUID(200),
			Title:    "my title with thrusters",
			Content:  "my content with spaceships",
		},

		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(30), baseTime.Add(3*time.Hour), &updateTime),
			UpVotes:   128,
			DownVotes: 64,
		},

		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(4), baseTime.Add(3*time.Hour), &updateTime),
			SourceID: goframework.NumberUUID(30),
			UserID:   goframework.NumberUUID(300),
			Title:    "my title with super thrusters",
			Content:  "my content with super spaceships",
		},

		&dao.ImproveRequestModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(40), baseTime.Add(4*time.Hour), &updateTime),
			UpVotes:   128,
			DownVotes: 64,
		},

		&dao.ImproveRequestRevisionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(5), baseTime.Add(4*time.Hour), &updateTime),
			SourceID: goframework.NumberUUID(40),
			UserID:   goframework.NumberUUID(100),
			Title:    "my title with tomatoes",
			Content:  "my content with super chips",
		},

		// Suggestions for goframework.NumberUUID(1).
		&dao.ImproveSuggestionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   uuid.Nil,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},

		// Suggestions for goframework.NumberUUID(2).
		&dao.ImproveSuggestionModel{
			Metadata: bunovel.NewMetadata(goframework.NumberUUID(4), baseTime, &updateTime),
			SourceID: goframework.NumberUUID(10),
			UserID:   uuid.Nil,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(5), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    uuid.Nil,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "suggestion title",
				Content:   "suggestion content",
			},
		},
	}

	data := []struct {
		name string

		ids []uuid.UUID

		expect    []*dao.ImproveRequestPreview
		expectErr error
	}{
		{
			name: "Success",
			ids:  []uuid.UUID{goframework.NumberUUID(10), goframework.NumberUUID(20), goframework.NumberUUID(60)},
			expect: []*dao.ImproveRequestPreview{
				{
					Metadata:                 bunovel.NewMetadata(goframework.NumberUUID(10), baseTime, &updateTime),
					UserID:                   goframework.NumberUUID(100),
					Title:                    "my title with spaceships",
					Content:                  "my content with thrusters",
					UpVotes:                  160,
					DownVotes:                80,
					RevisionCount:            2,
					SuggestionsCount:         5,
					AcceptedSuggestionsCount: 3,
				},
				{
					Metadata:      bunovel.NewMetadata(goframework.NumberUUID(20), baseTime.Add(2*time.Hour), &updateTime),
					UserID:        goframework.NumberUUID(200),
					Title:         "my title with thrusters",
					Content:       "my content with spaceships",
					RevisionCount: 1,
					UpVotes:       128,
					DownVotes:     64,
				},
			},
		},
		{
			name:   "Success/NoResults",
			ids:    []uuid.UUID{},
			expect: []*dao.ImproveRequestPreview{},
		},
	}

	err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
		repository := dao.NewImproveRequestRepository(tx)

		for _, d := range data {
			t.Run(d.name, func(st *testing.T) {
				res, err := repository.List(ctx, d.ids)
				require.ErrorIs(t, err, d.expectErr)
				require.Equal(t, d.expect, res)
			})
		}
	})
	require.NoError(t, err)
}
