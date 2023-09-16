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
	authClient := config.GetAuthClient(logger)
	authorizationsClient := config.GetAuthorizationsClient(logger)

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

	createImproveRequestService := services.NewCreateImproveRequestService(improveRequestsDAO, authClient, authorizationsClient)
	createImproveSuggestionService := services.NewCreateImproveSuggestionService(improveSuggestionDAO, improveRequestsDAO, authClient, authorizationsClient)
	deleteImproveRequestService := services.NewDeleteImproveRequestService(improveRequestsDAO, authClient)
	deleteImproveRequestRevisionService := services.NewDeleteImproveRequestRevisionService(improveRequestsDAO, authClient)
	deleteImproveSuggestionService := services.NewDeleteImproveSuggestionService(improveSuggestionDAO, authClient)
	getImproveRequestService := services.NewGetImproveRequestService(improveRequestsDAO)
	getImproveRequestRevisionService := services.NewGetImproveRequestRevisionService(improveRequestsDAO)
	listImproveRequestRevisionsService := services.NewListImproveRequestRevisionsService(improveRequestsDAO)
	getImproveSuggestionService := services.NewGetImproveSuggestionService(improveSuggestionDAO)
	listImproveRequestsService := services.NewListImproveRequestsService(improveRequestsDAO)
	listImproveSuggestionsService := services.NewListImproveSuggestionsService(improveSuggestionDAO)
	searchImproveRequestsService := services.NewSearchImproveRequestsService(improveRequestsDAO)
	searchImproveSuggestionsService := services.NewSearchImproveSuggestionsService(improveSuggestionDAO)
	updateImproveSuggestionService := services.NewUpdateImproveSuggestionService(improveSuggestionDAO, improveRequestsDAO, authClient, authorizationsClient)
	validateImproveSuggestionService := services.NewValidateImproveSuggestionService(improveSuggestionDAO, improveRequestsDAO, authClient)

	createImproveRequestHandler := handlers.NewCreateImproveRequestHandler(createImproveRequestService)
	createImproveSuggestionHandler := handlers.NewCreateImproveSuggestionHandler(createImproveSuggestionService)
	deleteImproveRequestHandler := handlers.NewDeleteImproveRequestHandler(deleteImproveRequestService)
	deleteImproveRequestRevisionHandler := handlers.NewDeleteImproveRequestRevisionHandler(deleteImproveRequestRevisionService)
	deleteImproveSuggestionHandler := handlers.NewDeleteImproveSuggestionHandler(deleteImproveSuggestionService)
	getImproveRequestHandler := handlers.NewGetImproveRequestHandler(getImproveRequestService)
	getImproveRequestRevisionHandler := handlers.NewGetImproveRequestRevisionHandler(getImproveRequestRevisionService)
	listImproveRequestRevisionsHandler := handlers.NewListImproveRequestRevisionsHandler(listImproveRequestRevisionsService)
	getImproveSuggestionHandler := handlers.NewGetImproveSuggestionHandler(getImproveSuggestionService)
	listImproveRequestsHandler := handlers.NewListImproveRequestsHandler(listImproveRequestsService)
	listImproveSuggestionsHandler := handlers.NewListImproveSuggestionsHandler(listImproveSuggestionsService)
	searchImproveRequestsHandler := handlers.NewSearchImproveRequestsHandler(searchImproveRequestsService)
	searchImproveSuggestionsHandler := handlers.NewSearchImproveSuggestionsHandler(searchImproveSuggestionsService)
	updateImproveSuggestionHandler := handlers.NewUpdateImproveSuggestionHandler(updateImproveSuggestionService)
	validateImproveSuggestionHandler := handlers.NewValidateImproveSuggestionHandler(validateImproveSuggestionService)

	router := apis.GetRouter(apis.RouterConfig{
		Logger:    logger,
		ProjectID: config.Deploy.ProjectID,
		CORS:      apis.GetCORS(config.App.Frontend.URLs),
		Prod:      config.ENV == config.ProdENV,
		Health: map[string]apis.HealthChecker{
			"postgres": func() error {
				return postgres.PingContext(ctx)
			},
			"auth-client": func() error {
				return authClient.Ping(ctx)
			},
			"authorizations-client": func() error {
				return authorizationsClient.Ping(ctx)
			},
		},
	})

	router.PUT("/improve-request", createImproveRequestHandler.Handle)
	router.PUT("/improve-suggestion", createImproveSuggestionHandler.Handle)
	router.DELETE("/improve-request", deleteImproveRequestHandler.Handle)
	router.DELETE("/improve-suggestion", deleteImproveSuggestionHandler.Handle)
	router.GET("/improve-request", getImproveRequestHandler.Handle)
	router.GET("/improve-request/revision", getImproveRequestRevisionHandler.Handle)
	router.DELETE("/improve-request/revision", deleteImproveRequestRevisionHandler.Handle)
	router.GET("/improve-request/revisions", listImproveRequestRevisionsHandler.Handle)
	router.GET("/improve-suggestion", getImproveSuggestionHandler.Handle)
	router.GET("/improve-requests", listImproveRequestsHandler.Handle)
	router.GET("/improve-suggestions", listImproveSuggestionsHandler.Handle)
	router.GET("/improve-requests/search", searchImproveRequestsHandler.Handle)
	router.GET("/improve-suggestions/search", searchImproveSuggestionsHandler.Handle)
	router.PATCH("/improve-suggestion", updateImproveSuggestionHandler.Handle)
	router.POST("/improve-suggestion/validate", validateImproveSuggestionHandler.Handle)

	if err := router.Run(fmt.Sprintf(":%d", config.API.Port)); err != nil {
		logger.Fatal().Err(err).Msg("a fatal error occurred while running the API, and the server had to shut down")
	}
}
