package main

import (
	"context"
	"fmt"
	"github.com/a-novel/bunovel"
	"github.com/a-novel/forum-service/config"
	"github.com/a-novel/forum-service/migrations"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/handlers"
	"github.com/a-novel/forum-service/pkg/services"
	"github.com/a-novel/go-apis"
	"io/fs"
)

func main() {
	ctx := context.Background()
	logger := config.GetLogger()

	postgres, sql, err := bunovel.NewClient(ctx, bunovel.Config{
		Driver:                &bunovel.PGDriver{DSN: config.Postgres.DSN, AppName: config.App.Name},
		Migrations:            &bunovel.MigrateConfig{Files: []fs.FS{migrations.Migrations}},
		DiscardUnknownColumns: true,
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("error connecting to postgres")
	}
	defer func() {
		_ = postgres.Close()
		_ = sql.Close()
	}()

	improveRequestsDAO := dao.NewImproveRequestRepository(postgres)
	improveSuggestionDAO := dao.NewImproveSuggestionRepository(postgres)

	voteImproveRequestService := services.NewVoteImproveRequestService(improveRequestsDAO)
	voteImproveSuggestionService := services.NewVoteImproveSuggestionService(improveSuggestionDAO)
	getImproveRequestService := services.NewGetImproveRequestService(improveRequestsDAO)
	getImproveSuggestionService := services.NewGetImproveSuggestionService(improveSuggestionDAO)

	voteImproveRequestHandler := handlers.NewVoteImproveRequestHandler(voteImproveRequestService)
	voteImproveSuggestionHandler := handlers.NewVoteImproveSuggestionHandler(voteImproveSuggestionService)
	getImproveRequestHandler := handlers.NewGetImproveRequestHandler(getImproveRequestService)
	getImproveSuggestionHandler := handlers.NewGetImproveSuggestionHandler(getImproveSuggestionService)

	router := apis.GetRouter(apis.RouterConfig{
		Logger:    logger,
		ProjectID: config.Deploy.ProjectID,
		CORS:      apis.GetCORS(config.App.Frontend.URLs),
		Prod:      config.ENV == config.ProdENV,
		Health: map[string]apis.HealthChecker{
			"postgres": func() error {
				return postgres.PingContext(ctx)
			},
		},
	})

	router.POST("/improve-request/vote", voteImproveRequestHandler.Handle)
	router.POST("/improve-suggestion/vote", voteImproveSuggestionHandler.Handle)
	router.GET("/improve-request", getImproveRequestHandler.Handle)
	router.GET("/improve-suggestion", getImproveSuggestionHandler.Handle)

	if err := router.Run(fmt.Sprintf(":%d", config.API.PortInternal)); err != nil {
		logger.Fatal().Err(err).Msg("a fatal error occurred while running the API, and the server had to shut down")
	}
}
