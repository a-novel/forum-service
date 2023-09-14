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

func TestImproveSuggestionRepository_Get(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   128,
			DownVotes: 64,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
	}

	data := []struct {
		name string

		id uuid.UUID

		expect    *dao.ImproveSuggestionModel
		expectErr error
	}{
		{
			name: "Success",
			id:   goframework.NumberUUID(1),
			expect: &dao.ImproveSuggestionModel{
				Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
				SourceID:  goframework.NumberUUID(10),
				UserID:    goframework.NumberUUID(100),
				UpVotes:   128,
				DownVotes: 64,
				Validated: true,
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: goframework.NumberUUID(1),
					Title:     "title",
					Content:   "content",
				},
			},
		},
		{
			name:      "Error/NotFound",
			id:        goframework.NumberUUID(2),
			expectErr: bunovel.ErrNotFound,
		},
	}

	err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
		repository := dao.NewImproveSuggestionRepository(tx)

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

func TestImproveSuggestionRepository_Create(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   128,
			DownVotes: 64,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
	}

	data := []struct {
		name string

		data     *dao.ImproveSuggestionModelCore
		userID   uuid.UUID
		sourceID uuid.UUID
		id       uuid.UUID
		now      time.Time

		expect    *dao.ImproveSuggestionModel
		expectErr error
	}{
		{
			name: "Success",
			data: &dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "my title",
				Content:   "my content",
			},
			userID:   goframework.NumberUUID(200),
			sourceID: goframework.NumberUUID(20),
			id:       goframework.NumberUUID(2),
			now:      baseTime,
			expect: &dao.ImproveSuggestionModel{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, nil),
				SourceID: goframework.NumberUUID(20),
				UserID:   goframework.NumberUUID(200),
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: goframework.NumberUUID(2),
					Title:     "my title",
					Content:   "my content",
				},
			},
		},
		{
			name: "Success/OnRevisionWithOtherSuggestions",
			data: &dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "my title",
				Content:   "my content",
			},
			userID:   goframework.NumberUUID(200),
			sourceID: goframework.NumberUUID(10),
			id:       goframework.NumberUUID(2),
			now:      baseTime,
			expect: &dao.ImproveSuggestionModel{
				Metadata: bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, nil),
				SourceID: goframework.NumberUUID(10),
				UserID:   goframework.NumberUUID(200),
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: goframework.NumberUUID(1),
					Title:     "my title",
					Content:   "my content",
				},
			},
		},
	}

	for _, d := range data {
		err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
			repository := dao.NewImproveSuggestionRepository(tx)
			t.Run(d.name, func(st *testing.T) {
				res, err := repository.Create(ctx, d.data, d.userID, d.sourceID, d.id, d.now)
				require.Equal(t, d.expect, res)
				require.ErrorIs(t, err, d.expectErr)
			})
		})
		require.NoError(t, err)
	}
}

func TestImproveSuggestionRepository_Update(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &baseTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   128,
			DownVotes: 64,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
	}

	data := []struct {
		name string

		data *dao.ImproveSuggestionModelCore
		id   uuid.UUID
		now  time.Time

		expect    *dao.ImproveSuggestionModel
		expectErr error
	}{
		{
			name: "Success",
			data: &dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "new title",
				Content:   "new content",
			},
			id:  goframework.NumberUUID(1),
			now: updateTime,
			expect: &dao.ImproveSuggestionModel{
				Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &updateTime),
				SourceID:  goframework.NumberUUID(10),
				UserID:    goframework.NumberUUID(100),
				UpVotes:   128,
				DownVotes: 64,
				Validated: true,
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: goframework.NumberUUID(2),
					Title:     "new title",
					Content:   "new content",
				},
			},
		},
		{
			name: "Error/NotFound",
			data: &dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "new title",
				Content:   "new content",
			},
			id:        goframework.NumberUUID(2),
			now:       updateTime,
			expectErr: bunovel.ErrNotFound,
		},
	}

	for _, d := range data {
		err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
			repository := dao.NewImproveSuggestionRepository(tx)
			t.Run(d.name, func(st *testing.T) {
				res, err := repository.Update(ctx, d.data, d.id, d.now)
				require.Equal(t, d.expect, res)
				require.ErrorIs(t, err, d.expectErr)
			})
		})
		require.NoError(t, err)
	}
}

func TestImproveSuggestionRepository_Delete(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &baseTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   128,
			DownVotes: 64,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
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
			repository := dao.NewImproveSuggestionRepository(tx)
			t.Run(d.name, func(st *testing.T) {
				err := repository.Delete(ctx, d.id)
				require.ErrorIs(t, err, d.expectErr)
			})
		})
		require.NoError(t, err)
	}
}

func TestImproveSuggestionRepository_Validate(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &baseTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   128,
			DownVotes: 64,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, &baseTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   32,
			DownVotes: 16,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
	}

	data := []struct {
		name string

		validated bool
		id        uuid.UUID

		expect    *dao.ImproveSuggestionModel
		expectErr error
	}{
		{
			name: "Success/UnValidate",
			id:   goframework.NumberUUID(1),
			expect: &dao.ImproveSuggestionModel{
				Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &baseTime),
				SourceID:  goframework.NumberUUID(10),
				UserID:    goframework.NumberUUID(100),
				UpVotes:   128,
				DownVotes: 64,
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: goframework.NumberUUID(1),
					Title:     "title",
					Content:   "content",
				},
			},
		},
		{
			name:      "Success/Validate",
			validated: true,
			id:        goframework.NumberUUID(2),
			expect: &dao.ImproveSuggestionModel{
				Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, &baseTime),
				SourceID:  goframework.NumberUUID(10),
				UserID:    goframework.NumberUUID(100),
				UpVotes:   32,
				DownVotes: 16,
				Validated: true,
				ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
					RequestID: goframework.NumberUUID(1),
					Title:     "title",
					Content:   "content",
				},
			},
		},
		{
			name:      "Error/NotFound",
			id:        goframework.NumberUUID(3),
			expectErr: bunovel.ErrNotFound,
		},
	}

	for _, d := range data {
		err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
			repository := dao.NewImproveSuggestionRepository(tx)
			t.Run(d.name, func(st *testing.T) {
				res, err := repository.Validate(ctx, d.validated, d.id)
				require.Equal(t, d.expect, res)
				require.ErrorIs(t, err, d.expectErr)
			})
		})
		require.NoError(t, err)
	}
}

func TestImproveSuggestionRepository_UpdateVotes(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, &baseTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   128,
			DownVotes: 64,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
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
			id:        goframework.NumberUUID(1),
			upVotes:   256,
			downVotes: 128,
		},
		{
			name:      "Error/NotFound",
			id:        goframework.NumberUUID(2),
			upVotes:   256,
			downVotes: 128,
			expectErr: bunovel.ErrNotFound,
		},
	}

	for _, d := range data {
		err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
			repository := dao.NewImproveSuggestionRepository(tx)
			t.Run(d.name, func(st *testing.T) {
				err := repository.UpdateVotes(ctx, d.id, d.upVotes, d.downVotes)
				require.ErrorIs(t, err, d.expectErr)
			})
		})
		require.NoError(t, err)
	}
}

func TestImproveSuggestionRepository_Search(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, lo.ToPtr(baseTime.Add(3*time.Hour))),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(200),
			UpVotes:   16,
			DownVotes: 8,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
			SourceID:  goframework.NumberUUID(20),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   32,
			DownVotes: 16,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, lo.ToPtr(baseTime.Add(time.Hour))),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   64,
			DownVotes: 32,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "title",
				Content:   "content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(4), baseTime, &baseTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   128,
			DownVotes: 64,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
	}

	data := []struct {
		name string

		query  dao.ImproveSuggestionSearchQuery
		limit  int
		offset int

		expect      []*dao.ImproveSuggestionModel
		expectCount int
		expectErr   error
	}{
		{
			name: "Success",
			expect: []*dao.ImproveSuggestionModel{
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, lo.ToPtr(baseTime.Add(3*time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
					SourceID:  goframework.NumberUUID(20),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, lo.ToPtr(baseTime.Add(time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   64,
					DownVotes: 32,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(2),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(4), baseTime, &baseTime),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   128,
					DownVotes: 64,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
			},
			expectCount: 4,
		},
		{
			name: "Success/FilterUserID",
			query: dao.ImproveSuggestionSearchQuery{
				UserID: lo.ToPtr(goframework.NumberUUID(100)),
			},
			expect: []*dao.ImproveSuggestionModel{
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
					SourceID:  goframework.NumberUUID(20),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, lo.ToPtr(baseTime.Add(time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   64,
					DownVotes: 32,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(2),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(4), baseTime, &baseTime),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   128,
					DownVotes: 64,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
			},
			expectCount: 3,
		},
		{
			name: "Success/FilterSourceID",
			query: dao.ImproveSuggestionSearchQuery{
				SourceID: lo.ToPtr(goframework.NumberUUID(10)),
			},
			expect: []*dao.ImproveSuggestionModel{
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, lo.ToPtr(baseTime.Add(3*time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, lo.ToPtr(baseTime.Add(time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   64,
					DownVotes: 32,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(2),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(4), baseTime, &baseTime),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   128,
					DownVotes: 64,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
			},
			expectCount: 3,
		},
		{
			name: "Success/FilterRequestID",
			query: dao.ImproveSuggestionSearchQuery{
				RequestID: lo.ToPtr(goframework.NumberUUID(1)),
			},
			expect: []*dao.ImproveSuggestionModel{
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, lo.ToPtr(baseTime.Add(3*time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
					SourceID:  goframework.NumberUUID(20),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(4), baseTime, &baseTime),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   128,
					DownVotes: 64,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
			},
			expectCount: 3,
		},
		{
			name: "Success/FilterValidated",
			query: dao.ImproveSuggestionSearchQuery{
				Validated: lo.ToPtr(true),
			},
			expect: []*dao.ImproveSuggestionModel{
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, lo.ToPtr(baseTime.Add(3*time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
					SourceID:  goframework.NumberUUID(20),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, lo.ToPtr(baseTime.Add(time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   64,
					DownVotes: 32,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(2),
						Title:     "title",
						Content:   "content",
					},
				},
			},
			expectCount: 3,
		},
		{
			name: "Success/OrderByScore",
			query: dao.ImproveSuggestionSearchQuery{
				Order: &dao.ImproveSuggestionSearchQueryOrder{
					Score: true,
				},
			},
			expect: []*dao.ImproveSuggestionModel{
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(4), baseTime, &baseTime),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   128,
					DownVotes: 64,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, lo.ToPtr(baseTime.Add(time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   64,
					DownVotes: 32,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(2),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
					SourceID:  goframework.NumberUUID(20),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, lo.ToPtr(baseTime.Add(3*time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
			},
			expectCount: 4,
		},
		{
			name:  "Success/Limit",
			limit: 2,
			expect: []*dao.ImproveSuggestionModel{
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, lo.ToPtr(baseTime.Add(3*time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
					SourceID:  goframework.NumberUUID(20),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
			},
			expectCount: 4,
		},
		{
			name:   "Success/Offset",
			offset: 1,
			limit:  2,
			expect: []*dao.ImproveSuggestionModel{
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
					SourceID:  goframework.NumberUUID(20),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, lo.ToPtr(baseTime.Add(time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   64,
					DownVotes: 32,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(2),
						Title:     "title",
						Content:   "content",
					},
				},
			},
			expectCount: 4,
		},
		{
			name:        "Success/OffsetTooLarge",
			offset:      5,
			expect:      []*dao.ImproveSuggestionModel{},
			expectCount: 4,
		},
		{
			name: "Success/NoResults",
			query: dao.ImproveSuggestionSearchQuery{
				RequestID: lo.ToPtr(goframework.NumberUUID(3)),
			},
			expect: []*dao.ImproveSuggestionModel{},
		},
	}

	err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
		repository := dao.NewImproveSuggestionRepository(tx)

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

func TestImproveSuggestionRepository_List(t *testing.T) {
	db, sqlDB := bunovel.GetTestPostgres(t, []fs.FS{migrations.Migrations})
	defer db.Close()
	defer sqlDB.Close()

	fixtures := []interface{}{
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, lo.ToPtr(baseTime.Add(3*time.Hour))),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(200),
			UpVotes:   16,
			DownVotes: 8,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
			SourceID:  goframework.NumberUUID(20),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   32,
			DownVotes: 16,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(3), baseTime, lo.ToPtr(baseTime.Add(time.Hour))),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   64,
			DownVotes: 32,
			Validated: true,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(2),
				Title:     "title",
				Content:   "content",
			},
		},
		&dao.ImproveSuggestionModel{
			Metadata:  bunovel.NewMetadata(goframework.NumberUUID(4), baseTime, &baseTime),
			SourceID:  goframework.NumberUUID(10),
			UserID:    goframework.NumberUUID(100),
			UpVotes:   128,
			DownVotes: 64,
			ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
				RequestID: goframework.NumberUUID(1),
				Title:     "title",
				Content:   "content",
			},
		},
	}

	data := []struct {
		name string

		ids []uuid.UUID

		expect    []*dao.ImproveSuggestionModel
		expectErr error
	}{
		{
			name: "Success",
			ids:  []uuid.UUID{goframework.NumberUUID(1), goframework.NumberUUID(2), goframework.NumberUUID(6)},
			expect: []*dao.ImproveSuggestionModel{
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(1), baseTime, lo.ToPtr(baseTime.Add(3*time.Hour))),
					SourceID:  goframework.NumberUUID(10),
					UserID:    goframework.NumberUUID(200),
					UpVotes:   16,
					DownVotes: 8,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
				{
					Metadata:  bunovel.NewMetadata(goframework.NumberUUID(2), baseTime, lo.ToPtr(baseTime.Add(2*time.Hour))),
					SourceID:  goframework.NumberUUID(20),
					UserID:    goframework.NumberUUID(100),
					UpVotes:   32,
					DownVotes: 16,
					Validated: true,
					ImproveSuggestionModelCore: dao.ImproveSuggestionModelCore{
						RequestID: goframework.NumberUUID(1),
						Title:     "title",
						Content:   "content",
					},
				},
			},
		},
		{
			name:   "Success/NoResults",
			ids:    []uuid.UUID{},
			expect: []*dao.ImproveSuggestionModel{},
		},
	}

	err := bunovel.RunTransactionalTest(db, fixtures, func(ctx context.Context, tx bun.Tx) {
		repository := dao.NewImproveSuggestionRepository(tx)

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
