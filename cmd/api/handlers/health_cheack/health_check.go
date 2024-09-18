package health_check

import (
	"context"
	"fdms/src/infrastructure/keycloak"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HealthCheckStatus struct {
	DbStatus       string `json:"db_status"`
	KeycloakStatus string `json:"keycloak_status"`
}

type HealthCheckService struct {
	db              *pgxpool.Pool
	keycloakService *keycloak.KeycloakAuthenticationService
}

func NewHealthCheckService(db *pgxpool.Pool, keycloakService *keycloak.KeycloakAuthenticationService) *HealthCheckService {
	return &HealthCheckService{
		db:              db,
		keycloakService: keycloakService,
	}
}

func (h *HealthCheckService) Check() HealthCheckStatus {

	conn, err := h.db.Acquire(context.Background())

	if err != nil {
		return HealthCheckStatus{
			DbStatus: "down",
		}
	}

	defer conn.Release()

	err = conn.Ping(context.Background())

	if err != nil {
		return HealthCheckStatus{
			DbStatus: "down",
		}
	}

	err = h.keycloakService.Ping(context.Background())

	if err != nil {
		return HealthCheckStatus{
			KeycloakStatus: "down",
		}
	}

	return HealthCheckStatus{
		DbStatus:       "up",
		KeycloakStatus: "up",
	}
}

func (h *HealthCheckService) HealthCheckHandler(c *gin.Context) {
	status := h.Check()
	c.JSON(http.StatusOK, status)
}
