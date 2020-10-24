package healthcheck

import (
	"gorm.io/gorm"
	"net/http"

	"github.com/brpaz/go-healthcheck"
	"github.com/brpaz/go-healthcheck/checks"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	AppName        string
	AppDescription string
	BuildCommit    string
	db             *gorm.DB
}

func NewHandler(appName string, appDescription string, buildCommit string, db *gorm.DB) *Handler {
	return &Handler{
		AppName:        appName,
		AppDescription: appDescription,
		BuildCommit:    buildCommit,
		db:             db,
	}
}

func (h *Handler) Handle(c echo.Context) error {

	dbConn, _ := h.db.DB()
	health := healthcheck.New(h.AppName, h.AppDescription, h.BuildCommit, "")
	health.AddCheckProvider(checks.NewSysInfoChecker())
	health.AddCheckProvider(checks.NewDBChecker("database", dbConn))
	healthResult := health.Get()

	c.Response().Header().Set(echo.HeaderContentType, "application/health+json")

	statusCode := http.StatusOK
	if healthResult.Status == healthcheck.Fail {
		statusCode = http.StatusServiceUnavailable
	}

	return c.JSON(statusCode, healthResult)
}
