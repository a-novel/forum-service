package config

import (
	auth "github.com/a-novel/auth-service/framework"
	"github.com/rs/zerolog"
	"net/url"
)

func GetAuthClient(logger zerolog.Logger) auth.Client {
	authURL, err := new(url.URL).Parse(API.External.AuthAPI)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not parse auth API URL")
	}

	return auth.NewClient(authURL)
}
