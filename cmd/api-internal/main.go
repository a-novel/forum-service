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

	postgres, closer := config.GetPostgres(logger)
	defer closer()

	improveRequestsDAO := dao.NewImproveRequestRepository(postgres)
	improveSuggestionDAO := dao.NewImproveSuggestionRepository(postgres)

	voteImproveRequestService := services.NewVoteImproveRequestService(improveRequestsDAO)
	voteImproveSuggestionService := services.NewVoteImproveSuggestionService(improveSuggestionDAO)
	getImproveRequestService := services.NewGetImproveRequestService(improveRequestsDAO)
	getImproveSuggestionService := services.NewGetImproveSuggestionService(improveSuggestionDAO)

	pingHandler := handlers.NewPingHandler()
	healthCheckHandler := handlers.NewHealthCheckHandler(postgres, nil)
	voteImproveRequestHandler := handlers.NewVoteImproveRequestHandler(voteImproveRequestService)
	voteImproveSuggestionHandler := handlers.NewVoteImproveSuggestionHandler(voteImproveSuggestionService)
	getImproveRequestHandler := handlers.NewGetImproveRequestHandler(getImproveRequestService)
	getImproveSuggestionHandler := handlers.NewGetImproveSuggestionHandler(getImproveSuggestionService)

	router := config.GetRouter(logger)

	router.GET("/ping", pingHandler.Handle)
	router.GET("/healthcheck", healthCheckHandler.Handle)
	router.POST("/improve-request/vote", voteImproveRequestHandler.Handle)
	router.POST("/improve-suggestion/vote", voteImproveSuggestionHandler.Handle)
	router.GET("/improve-request", getImproveRequestHandler.Handle)
	router.GET("/improve-suggestion", getImproveSuggestionHandler.Handle)

	if err := router.Run(fmt.Sprintf(":%d", config.API.PortInternal)); err != nil {
		logger.Fatal().Err(err).Msg("a fatal error occurred while running the API, and the server had to shut down")
	}
}
