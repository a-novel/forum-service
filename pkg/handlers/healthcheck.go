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

func NewHealthCheckHandler(db *bun.DB) HealthCheckHandler {
	return &healthCheckHandlerImpl{db: db}
}

type healthCheckHandlerImpl struct {
	db         *bun.DB
	authClient auth.Client
}

func (h *healthCheckHandlerImpl) Handle(c *gin.Context) {
	dbErr := h.db.PingContext(c)
	dbStatus := gin.H{"available": dbErr == nil}
	if dbErr != nil {
		dbStatus["error"] = dbErr.Error()
	}

	authClientErr := h.authClient.Ping()
	authClientStatus := gin.H{"available": authClientErr == nil}
	if authClientErr != nil {
		authClientStatus["error"] = authClientErr.Error()
	}

	c.JSON(http.StatusOK, gin.H{
		"database": dbStatus,
		"clients": gin.H{
			"auth": authClientStatus,
		},
	})
}
