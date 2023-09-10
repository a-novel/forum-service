package main

import (
	"fmt"
	"github.com/a-novel/forum-service/config"
	"github.com/a-novel/forum-service/pkg/dao"
	"github.com/a-novel/forum-service/pkg/handlers"
	"github.com/a-novel/forum-service/pkg/services"
)

func main() {
	logger := config.GetLogger()
	authClient := config.GetAuthClient()

	postgres, closer := config.GetPostgres(logger)
	defer closer()

	improveRequestsDAO := dao.NewImproveRequestRepository(postgres)
	improveSuggestionDAO := dao.NewImproveSuggestionRepository(postgres)

	createImproveRequestService := services.NewCreateImproveRequestService(improveRequestsDAO, authClient)
	createImproveSuggestionService := services.NewCreateImproveSuggestionService(improveSuggestionDAO, improveRequestsDAO, authClient)
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
	updateImproveSuggestionService := services.NewUpdateImproveSuggestionService(improveSuggestionDAO, improveRequestsDAO, authClient)
	validateImproveSuggestionService := services.NewValidateImproveSuggestionService(improveSuggestionDAO, improveRequestsDAO, authClient)

	pingHandler := handlers.NewPingHandler()
	healthCheckHandler := handlers.NewHealthCheckHandler(postgres)
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

	router := config.GetRouter(logger)

	router.GET("/ping", pingHandler.Handle)
	router.GET("/healthcheck", healthCheckHandler.Handle)
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
