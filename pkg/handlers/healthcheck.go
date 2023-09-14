package handlers

import (
	auth "github.com/a-novel/auth-service/framework"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"net/http"
)

type HealthCheckHandler interface {
	Handle(c *gin.Context)
}

func NewHealthCheckHandler(db *bun.DB, authClient auth.Client) HealthCheckHandler {
	return &healthCheckHandlerImpl{db: db, authClient: authClient}
}

type healthCheckHandlerImpl struct {
	db         *bun.DB
	authClient auth.Client
}

type HealthCheckResponse struct {
	Database struct {
		Available bool   `json:"available"`
		Error     string `json:"error,omitempty"`
	} `json:"database"`
	Clients struct {
		Auth struct {
			Available bool   `json:"available"`
			Error     string `json:"error,omitempty"`
		} `json:"auth"`
	} `json:"clients"`
}

func (h *healthCheckHandlerImpl) Handle(c *gin.Context) {
	res := new(HealthCheckResponse)

	dbErr := h.db.PingContext(c)
	res.Database.Available = dbErr == nil
	if dbErr != nil {
		res.Database.Error = dbErr.Error()
	}

	// Internal api has no auth client.
	if h.authClient != nil {
		authClientErr := h.authClient.Ping()
		res.Clients.Auth.Available = authClientErr == nil
		if authClientErr != nil {
			res.Clients.Auth.Error = authClientErr.Error()
		}
	}

	c.JSON(http.StatusOK, res)
}
