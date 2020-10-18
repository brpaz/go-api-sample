package db

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brpaz/go-api-sample/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

const MockDriverName = "mock"

// GetConnection Returns a database connection
func GetConnection(cfg config.Config, logger *zap.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Database,
		cfg.DB.Port)

	logger.Info(dsn)
	l := zapgorm2.New(logger)

	if cfg.Env == config.EnvDev {
		l.LogMode(gormLogger.Info)
	}

	var dialector gorm.Dialector
	if cfg.DB.Driver == MockDriverName {
		mockConn, _, _ := sqlmock.New() // mock sql.DB
		dialector = postgres.New(postgres.Config{
			Conn: mockConn,
		})
	} else {
		dialector = postgres.Open(dsn)
	}

	l.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	return gorm.Open(dialector, &gorm.Config{
		Logger: l,
	})
}
