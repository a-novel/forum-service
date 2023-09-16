package config

import (
	apiclients "github.com/a-novel/go-api-clients"
	"github.com/rs/zerolog"
	"net/url"
)

func GetAuthClient(logger zerolog.Logger) apiclients.AuthClient {
	authURL, err := new(url.URL).Parse(API.External.AuthAPI)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not parse auth API URL")
	}

	return apiclients.NewAuthClient(authURL)
}

func GetAuthorizationsClient(logger zerolog.Logger) apiclients.AuthorizationsClient {
	authorizationsURL, err := new(url.URL).Parse(API.External.AuthorizationsAPI)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not parse auth API URL")
	}

	return apiclients.NewAuthorizationsClient(authorizationsURL)
}
