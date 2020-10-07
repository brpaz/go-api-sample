package healthcheck

import (
	"net/http"

	"github.com/brpaz/go-healthcheck"
	"github.com/brpaz/go-healthcheck/checks"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	AppName string
	AppDescription string
	BuildCommit string
}

func NewHandler(appName string, appDescription string, buildCommit string) Handler {
	return Handler{
		AppName:        appName,
		AppDescription: appDescription,
		BuildCommit:    buildCommit,
	}
}

func (h *Handler) Handle(c echo.Context) error {
	health := healthcheck.New(h.AppName, h.AppDescription, h.BuildCommit, "")
	health.AddCheckProvider(checks.NewSysInfoChecker())

	healthResult := health.Get()

	c.Response().Header().Set(echo.HeaderContentType, "application/health+json")

	statusCode := http.StatusOK
	if healthResult.Status == healthcheck.Fail {
		statusCode = http.StatusServiceUnavailable
	}

	return c.JSON(statusCode, healthResult)
}
