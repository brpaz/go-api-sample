package healthcheck

import (
	"github.com/brpaz/go-api-sample/internal/buildinfo"
	"net/http"

	"github.com/brpaz/go-healthcheck"
	"github.com/brpaz/go-healthcheck/checks"
	"github.com/labstack/echo/v4"
)

// HealthHandler Handler for the health check endpoint
func HealthHandler(c echo.Context) error {

	health := healthcheck.New(buildinfo.AppName, buildinfo.AppDescription, buildinfo.BuildCommit, "")
	health.AddCheckProvider(checks.NewSysInfoChecker())

	healthResult := health.Get()

	c.Response().Header().Set(echo.HeaderContentType, "application/health+json")

	statusCode := http.StatusOK
	if healthResult.Status == healthcheck.Fail {
		statusCode = http.StatusServiceUnavailable
	}

	return c.JSON(statusCode, healthResult)
}
