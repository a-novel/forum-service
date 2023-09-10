package config

import (
	auth "github.com/a-novel/auth-service/framework"
	"github.com/samber/lo"
)

func GetAuthClient() auth.Client {
	authURL := lo.Ternary(ENV == ProdENV, auth.ProdURL, auth.DevURL)
	return auth.NewClient(authURL)
}
