package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthResponse The response struct
type HealthResponse struct {
	Status string `json:"status"`
}

// Health Healthcheck handler
func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{
		Status: "OK",
	})
}
